#!/bin/bash

# Build Perry using Go
go build

# Move the compiled binary to a directory in the system's PATH
sudo mv perry /usr/local/bin

echo "Perry has been installed successfully."
