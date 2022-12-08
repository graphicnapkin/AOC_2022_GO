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
	data, _ := input()
	part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	initialStackState := [][]string{}
	instructions := [][]int{}

	pastInitialState := false

	for _, row := range data {
		if row[1] == '1' {
			pastInitialState = true
			continue
		}

		if !pastInitialState {
			crates := []string{}
			// each instruction is 4 characters long, with the letter in the second index
			for i := 0; i < len(row); i += 4 {
				if string(row[i+1]) != "" {
					crates = append(crates, string(row[i+1]))
				}
			}
			initialStackState = append(initialStackState, crates)
		} else {
			stringRow := strings.Split(row, " ")
			parsed := []int{}
			for _, item := range stringRow {
				if item != "move" && item != "from" && item != "to" {
					number, _ := strconv.Atoi(item)
					parsed = append(parsed, number)
				}
			}
			instructions = append(instructions, parsed)
		}
	}

	stacks := make([]Stack, len(initialStackState[len(initialStackState)-1]))

	for _, row := range initialStackState {
		for i, column := range row {
			if column != " " {
				stacks[i].Push(column)
			}
		}
	}

	for _, instruction := range instructions {
		numberToMove := instruction[0]
		fromStack := instruction[1]
		toStack := instruction[2]

		for i := 0; i < numberToMove; i++ {
			crateToMove, _ := stacks[fromStack-1].Pop()
			stacks[toStack-1].AddCrate(crateToMove)
		}
	}

	output := ""
	for _, stack := range stacks {
		output += stack[0]
	}
	fmt.Println(output)
}

func part2(data []string) {
	fmt.Println("Part 2")
	initialStackState := [][]string{}
	instructions := [][]int{}

	pastInitialState := false

	for _, row := range data {
		if row[1] == '1' {
			pastInitialState = true
			continue
		}

		if !pastInitialState {
			crates := []string{}
			for i := 0; i < len(row); i += 4 {
				if string(row[i+1]) != "" {
					crates = append(crates, string(row[i+1]))
				}
			}
			initialStackState = append(initialStackState, crates)
		} else {
			stringRow := strings.Split(row, " ")
			parsed := []int{}
			for _, item := range stringRow {
				if item != "move" && item != "from" && item != "to" {
					number, _ := strconv.Atoi(item)
					parsed = append(parsed, number)
				}
			}
			instructions = append(instructions, parsed)
		}
	}

	stacks := make([]Stack, len(initialStackState[len(initialStackState)-1]))

	for _, row := range initialStackState {
		for i, column := range row {
			if column != " " {
				stacks[i].Push(column)
			}
		}
	}

	for _, instruction := range instructions {
		numberToMove := instruction[0]
		fromStack := instruction[1]
		toStack := instruction[2]

		cratesToMove := []string{}

		for i := 0; i < numberToMove; i++ {
			crateToMove, _ := stacks[fromStack-1].Pop()
			cratesToMove = append(cratesToMove, crateToMove)
		}

		stacks[toStack-1].AddCrates(cratesToMove)
	}

	output := ""
	for _, stack := range stacks {
		output += stack[0]
	}

	fmt.Println(output)
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

		data = append(data, rec...)
	}

	return data
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// prepend
func (s *Stack) AddCrate(str string) {
	*s = append([]string{str}, *s...)
}

func (s *Stack) AddCrates(strs []string) {
	*s = append(strs, *s...)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		element := (*s)[0]
		*s = (*s)[1:]
		return element, true
	}
}
