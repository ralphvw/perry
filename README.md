# Perry

Perry is a command-line interface (CLI) tool that simplifies the process of initializing a new Go web server with auth based on a template repository.

## Features

- Interactive CLI for project initialization.
- Customizable module name and project structure.
- Initializes Git repository and runs `go mod tidy` automatically.
- Uses net/http

## Installation

### From Source

1. Clone the Perry repository:

   ```bash
   git clone https://github.com/yourusername/perry.git
   ```

2. Navigate to the cloned directory:

   ```bash
   cd perry
   ```

3. Build Perry using Go:

   ```bash
   go build
   ```

4. Move the compiled binary to a directory in your system's PATH:

   ```bash
   sudo mv perry /usr/local/bin
   ```

### Using Installation Script

1. Clone the Perry repository:

   ```bash
   git clone https://github.com/ralphvw/perry.git
   ```

2. Navigate to the cloned directory:

   ```bash
   cd perry
   ```

3. Run the installation script appropriate for your system:

   - For Unix-based systems:

     ```bash
     ./install.sh
     ```

   - For Windows:

     ```batch
     install.bat
     ```

## Usage

To initialize a new Go project using Perry:

```bash
perry init
```
