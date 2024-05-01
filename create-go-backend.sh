#! /bin/bash

read -p "Enter your project folder path: " folderPath
folderPath="${folderPath/#\~/$HOME}"

read -p "Enter your go module path: " modulePath

mkdir $folderPath
cp -r /tmp/create-go-backend/. $folderPath
rm -rf $folderPath/.git
rm $folderPath/create-go-backend.sh

cd $folderPath
go mod init $modulePath

rm -rf /tmp/create-go-backend
echo "Go backend project folder created"

# TODO: Make an automated test to make sure all content is copied properly
# TODO: README.md (prereqs, how to run via docker compose)
