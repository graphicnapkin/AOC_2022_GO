package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//[testData, realData]
	data, _ := input()
	part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	sum := 0

	for _, row := range data {
		split := len(row) / 2
		firstHalf := row[:split]
		secondHalf := row[split:]

		for _, char := range firstHalf {
			if contains(secondHalf, char) {
				if int(char) > 96 {
					sum += (int(char) - 96)
				} else {
					sum += (int(char) - 38)
				}
				break
			}
		}
	}
	fmt.Println(sum)
}

func part2(data []string) {
	fmt.Println("Part 2")
	sum := 0

	for i := 0; i < len(data); i += 3 {
		first := data[i]
		second := data[i+1]
		third := data[i+2]

		for _, char := range first {
			if contains(second, char) {
				if contains(third, char) {
					if int(char) > 96 {
						sum += (int(char) - 96)
					} else {
						sum += (int(char) - 38)
					}
					break
				}
			}
		}
	}
	fmt.Println(sum)
}

func input() ([]string, []string) {
	test := openCSV("./input/testInput.csv")
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

func contains(s string, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
