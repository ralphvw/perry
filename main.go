package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Step 1: Ask user for project name
	projectName := getUserInput("Enter project name: ")

	// Step 2: Download template repository zip
	fmt.Println("Setting up...")
	downloadURL := "https://github.com/ralphvw/go-template/archive/main.zip"
	templateZip := "template.zip"
	err := downloadFile(templateZip, downloadURL)
	if err != nil {
		fmt.Println("Error downloading template repository:", err)
		os.Exit(1)
	}
	defer os.Remove(templateZip)

	// Step 3: Extract template repository zip and rename root folder
	err = unzipAndRename(templateZip, projectName)
	if err != nil {
		fmt.Println("Error extracting and renaming template repository:", err)
		os.Exit(1)
	}

	// Step 4: Ask user for module name
	moduleName := getUserInput("Enter module name (e.g., github.com/yourname/project): ")

	// Step 5: Replace module name in all files
	replaceModuleNames(projectName, moduleName)

	// Step 6: Initialize Git repository
	initGitRepository(projectName)

	// Step 7: Run go mod tidy
	runGoModTidy(projectName)

	fmt.Println("Project initialized successfully!")
}

// getUserInput prompts the user and returns their input
func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// downloadFile downloads a file from the specified URL
func downloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// unzipAndRename extracts a zip file to the current directory and renames the root folder to the project name
func unzipAndRename(src, projectName string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// Extract files to the current directory
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Determine the destination path
		path := filepath.Join(projectName, strings.TrimPrefix(f.Name, "go-template-main"))

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), os.ModePerm)
			fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer fw.Close()

			_, err = io.Copy(fw, rc)
			if err != nil {
				return err
			}
		}
	}

	// Rename the root directory to the project name
	// err = os.Rename(filepath.Join(projectName, "go-template-main"), projectName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// replaceModuleNames replaces module names in all files
func replaceModuleNames(projectName, moduleName string) {
	fmt.Println("Replacing module names...")

	// Walk through all files in the project directory
	err := filepath.Walk(projectName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read file contents
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace module name in file contents
		newData := strings.ReplaceAll(string(data), "github.com/ralphvw/go-template", moduleName)

		// Write modified contents back to file
		err = os.WriteFile(path, []byte(newData), 0644)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error replacing module names:", err)
		os.Exit(1)
	}
}

// initGitRepository initializes a Git repository in the project directory
func initGitRepository(projectName string) {
	fmt.Println("Initializing Git repository...")

	cmd := exec.Command("git", "init", projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error initializing Git repository:", err)
		os.Exit(1)
	}
}

// runGoModTidy runs 'go mod tidy' command in the project directory
func runGoModTidy(projectName string) {
	fmt.Println("Running 'go mod tidy'...")

	// Change to project directory
	err := os.Chdir(projectName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	// Run 'go mod tidy' command
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running 'go mod tidy':", err)
		os.Exit(1)
	}
}
