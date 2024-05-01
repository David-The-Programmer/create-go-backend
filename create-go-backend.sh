#! /bin/bash

read -p "Enter your project folder path: " folderPath
read -p "Enter your go module path: " modulePath

# TODO: Need to fix being unable to use "~" to denote home user directory
mkdir $folderPath
cp -r /tmp/create-go-backend/. $folderPath
rm -rf $folderPath/.git
rm $folderPath/create-go-backend.sh

cd $folderPath
go mod init $modulePath

rm -rf /tmp/create-go-backend
echo "Go backend project folder created"

# TODO: Make an automated test to make sure all content is copied properly
