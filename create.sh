#! /bin/bash

# Get the absolute path of the actual script directory from root directory
SCRIPT_DIR=$( dirname -- "$( readlink -f -- "$0"; )"; )

read -p "Enter your project folder path: " project_path
# Expands "~" into user $HOME directory
project_path="${project_path/#\~/$HOME}"

read -p "Enter your go module path: " module_path

# read -p "Enter the go version of your go module: " go_version

# Make the project folder and copy all starter files into the folder
mkdir $project_path
cp -r $SCRIPT_DIR/. $project_path
rm -rf $project_path/.git
rm $project_path/create.sh

# Create go.mod file
cd $project_path
go mod init $module_path
# go mod edit -go=$go_version

# TODO: Test prompt for go version
# TODO: Change docker file go version accordingly as well
# TODO: Ensure script stops when there are errors
# TODO: Warn user of directory override when project folder path already exists
# TODO: Re-prompt user if user does not want to override existing folder in project folder path

echo "Go backend project folder created"

