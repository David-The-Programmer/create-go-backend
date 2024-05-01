#! /bin/bash

read -p "Enter your project folder path: " folderPath
read -p "Enter your go module path: " modulePath

mkdir $folderPath
cp -r /tmp/create-go-backend-temp/create-go-backend/* $folderPath
rm -r /tmp/create-go-backend-temp

cd $folderPath
go mod init $modulePath

echo "Go backend project folder created"
