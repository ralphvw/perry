@echo off

REM Build Perry using Go
go build

REM Move the compiled binary to a directory in the system's PATH
move perry.exe C:\Windows\System32

echo Perry has been installed successfully.
