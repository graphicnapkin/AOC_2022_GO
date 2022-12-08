#!/bin/bash
read -p 'What day is this for?
' dayNumber

read -p 'Copy test data to clipboard and hit enter
' > /dev/null
testData=$(eval "pbcopy")

read -p 'Copy real data to clipboard and hit enter
' > /dev/null 

realData=$(eval "pbcopy")

dirName="day"$dayNumber
mkdir $dirName 
cp -r template/* $dirName/

cd $dirName
go mod init > /dev/null

cd input
go install > /dev/null

cd ..
mv day.go $dirName".go"

sed -i "s/template/$dirName/" $dirName".go" 
sed -i "s/\"test_data\"/$testData/" ./input/input.go
sed -i "s/\"real_data\"/$realData/" ./input/input.go

nvim $dirName".go"
