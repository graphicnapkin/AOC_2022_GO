#!/bin/bash
read -p 'What day is this for?
' dayNumber

dirName="day"$dayNumber
mkdir $dirName 
cp -r day/* $dirName/

cd $dirName
mv day.go $dirName".go"