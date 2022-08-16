#!/bin/bash
read -p 'Data type (int, string, bool): ' dataType
read -p 'Test data (comma separated): ' testData
read -p 'Real data (comma separated): ' realData
dayNumber=$1
dirName="day"$dayNumber
mkdir $dirName 
cp -r template/* $dirName/
cd $dirName
go mod init
cd input
go install
cd ..
mv day.go $dirName".go"
sed -i "s/template/$dirName/" $dirName".go" 
sed -i "s:int):$dataType):g" $dirName".go" 
sed -i "s/00/$testData/" ./input/input.go
sed -i "s/01/$realData/" ./input/input.go
sed -i "s:int:$dataType:g" ./input/input.go