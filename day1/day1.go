package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	//[testData, realData]
	_, data := input()

	//part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	currentMax := 0
	tempCalories := 0

	for i, item := range data {
		fmt.Println(i)
		if item == "" {
			if tempCalories > currentMax {
				currentMax = tempCalories
			}
			tempCalories = 0
			continue
		}
		currentItem, _ := strconv.Atoi(item)
		tempCalories += currentItem
	}

	fmt.Println(currentMax)
}

func part2(data []string) {
	fmt.Println("Part 2")
	largest := []int{0, 0, 0}
	tempCalories := 0

	for i, item := range data {
		lastItem := i == len(data)-1

		if lastItem || item != "" {
			currentItem, _ := strconv.Atoi(item)
			tempCalories += currentItem
		}

		if item == "" || lastItem {
			if tempCalories > largest[0] {
				largest[1] = largest[0]
				largest[0] = tempCalories
			} else if tempCalories > largest[1] {
				largest[2] = largest[1]
				largest[1] = tempCalories
			} else if tempCalories > largest[2] {
				largest[2] = tempCalories
			}
			tempCalories = 0
			continue
		}
	}

	fmt.Println(largest[0] + largest[1] + largest[2])
}

func input() ([]string, []string) {
	test := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	data := openCSV("./input/input.csv")
	return test, data
}

func openCSV(fileName string) []string {
	data := []string{}

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

		data = append(data, []string{rec[0]}...)
	}

	return data
}
