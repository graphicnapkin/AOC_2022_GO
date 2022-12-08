package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
	part2(data)
}

func part1(data [][]string) {
	fmt.Println("Part 1")
	sum := 0

	for _, row := range data {
		firstString := strings.Split(row[0], "-")
		firstStart, _ := strconv.Atoi(firstString[0])
		firstEnd, _ := strconv.Atoi(firstString[1])

		secondString := strings.Split(row[1], "-")
		secondStart, _ := strconv.Atoi(secondString[0])
		secondEnd, _ := strconv.Atoi(secondString[1])

		if (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd) {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2(data [][]string) {
	fmt.Println("Part 2")
	sum := 0

	for _, row := range data {
		firstString := strings.Split(row[0], "-")
		firstStart, _ := strconv.Atoi(firstString[0])
		firstEnd, _ := strconv.Atoi(firstString[1])

		secondString := strings.Split(row[1], "-")
		secondStart, _ := strconv.Atoi(secondString[0])
		secondEnd, _ := strconv.Atoi(secondString[1])

		if (firstStart >= secondStart && firstStart <= secondEnd) || (secondStart >= firstStart && secondStart <= firstEnd) {
			sum++
		}
	}
	fmt.Println(sum)
}

func input() ([][]string, [][]string) {
	test := openCSV("./input/testInput.csv")
	data := openCSV("./input/input.csv")
	return test, data
}

func openCSV(fileName string) [][]string {
	data := [][]string{}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, rec)
	}

	return data
}
