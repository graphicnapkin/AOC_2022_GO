#!/bin/bash
read -p 'What day is this for?
' dayNumber

read -p 'Copy test data to clipboard and hit enter
'
testData=$(eval "powershell.exe Get-Clipboard | sed 's/\r//' | sed 's/.*/\"&\"/' | sed -z 's/\n/,/g;s/,$/\n/'")

read -p 'Copy real data to clipboard and hit enter
'
realData=$(eval "powershell.exe Get-Clipboard | sed 's/\r//' | sed 's/.*/\"&\"/' | sed -z 's/\n/,/g;s/,$/\n/'")

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
