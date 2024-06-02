#! /bin/bash

# Get the absolute path of the actual script directory from root directory
SCRIPT_DIR=$( dirname -- "$( readlink -f -- "$0"; )"; )

read -p "Enter your project folder path: " project_path
# Expands "~" into user $HOME directory
project_path="${project_path/#\~/$HOME}"

read -p "Enter your go module path: " module_path

read -p "Enter the go version of your go module: " GO_VERSION

# Download template from repo
wget -O - https://api.github.com/repos/David-The-Programmer/create-go-backend/contents/template | jq '.[]' | jq -r '.download_url'| xargs -r -I{} wget -P $project_path "{}"

# Create the .env file with the entered go version
echo "GO_VERSION=$GO_VERSION" > $project_path/.env

# Create go.mod file
cd $project_path
go mod init $module_path
go mod edit -go=$GO_VERSION

echo "Go backend project folder created"

