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
	//part1(data)
	part2(data)
}

func part1(data []string) {
	headPos := []int{0, 0}
	tailPos := []int{0, 0}
	tailPositions := make(map[string]bool)
	tailPositions["0,0"] = true

	for index, row := range data {
		movement := strings.Split(row, " ")
		direction := movement[0]
		moves, _ := strconv.Atoi(movement[1])

		if index < 100 {
			fmt.Println("Headpos Start", headPos)
			fmt.Println("Tailpos Start", tailPos)
		}

		if direction == "U" {
			for i := 0; i < moves; i++ {
				headPos[1]++
				if abs(headPos[1]-tailPos[1]) > 1 {
					tailPos[1]++
					tailPos[0] = headPos[0]
					tailPositions[makeCord(tailPos)] = true
				}

			}
		}

		if direction == "D" {
			for i := 0; i < moves; i++ {
				headPos[1]--
				if abs(headPos[1]-tailPos[1]) > 1 {
					tailPos[1]--
					tailPos[0] = headPos[0]
					tailPositions[makeCord(tailPos)] = true
				}
			}
		}

		if direction == "R" {
			for i := 0; i < moves; i++ {
				headPos[0]++
				if abs(headPos[0]-tailPos[0]) > 1 {
					tailPos[0]++
					tailPos[1] = headPos[1]
					tailPositions[makeCord(tailPos)] = true
				}
			}
		}

		if direction == "L" {
			for i := 0; i < moves; i++ {
				headPos[0]--
				if abs(headPos[0]-tailPos[0]) > 1 {
					tailPos[0]--
					tailPos[1] = headPos[1]
					tailPositions[makeCord(tailPos)] = true
				}
			}
		}
	}

	fmt.Println(len(tailPositions))
}

func part2(data []string) {
	headPos := []int{0, 0}
	tailPositions := make([][]int, 9)

	for i := 0; i < 9; i++ {
		tailPositions[i] = []int{0, 0}
	}

	lastTailPosition := make(map[string]bool)
	lastTailPosition["0,0"] = true

	for _, row := range data {
		movement := strings.Split(row, " ")
		direction := movement[0]
		moves, _ := strconv.Atoi(movement[1])

		for i := 0; i < moves; i++ {
			changeHeadPosition(direction, headPos)
			changeTailPositions(headPos, tailPositions, lastTailPosition)
		}
	}

	fmt.Println(len(lastTailPosition))
}

func changeHeadPosition(direction string, headPos []int) {
	if direction == "U" {
		headPos[1]--
	}
	if direction == "D" {
		headPos[1]++
	}
	if direction == "R" {
		headPos[0]++
	}
	if direction == "L" {
		headPos[0]--
	}
}

func changeTailPositions(headPos []int, tailPositions [][]int, tailPositionsMap map[string]bool) {
	for i := 0; i < len(tailPositions); i++ {
		var comparePos []int
		if i == 0 {
			comparePos = headPos
		} else {
			comparePos = tailPositions[i-1]
		}
		// up and to the left
		if comparePos[1]-tailPositions[i][1] < -1 && comparePos[0]-tailPositions[i][0] < -1 {
			tailPositions[i][1]--
			tailPositions[i][0]--
		}
		// up and to the right
		if comparePos[1]-tailPositions[i][1] < -1 && comparePos[0]-tailPositions[i][0] > 1 {
			tailPositions[i][1]--
			tailPositions[i][0]++
		}

		// down and to the left
		if comparePos[1]-tailPositions[i][1] > 1 && comparePos[0]-tailPositions[i][0] < -1 {
			tailPositions[i][1]++
			tailPositions[i][0]--
		}
		// down and to the right
		if comparePos[1]-tailPositions[i][1] > 1 && comparePos[0]-tailPositions[i][0] > 1 {
			tailPositions[i][1]++
			tailPositions[i][0]++
		}
		// if head y is more than one ABOVE tail y
		if comparePos[1]-tailPositions[i][1] < -1 {
			tailPositions[i][1]--
			tailPositions[i][0] = comparePos[0]
		}
		// if head y is more than one BELOW tail y
		if comparePos[1]-tailPositions[i][1] > 1 {
			tailPositions[i][1]++
			tailPositions[i][0] = comparePos[0]
		}
		// if head x is more than one to the RIGHT of tail x
		if comparePos[0]-tailPositions[i][0] > 1 {
			tailPositions[i][0]++
			tailPositions[i][1] = comparePos[1]
		}
		// if head x is more than one to the LEFT of tail x
		if comparePos[0]-tailPositions[i][0] < -1 {
			tailPositions[i][0]--
			tailPositions[i][1] = comparePos[1]
		}

		if i == len(tailPositions)-1 {
			tailPositionsMap[makeCord(tailPositions[i])] = true
		}
	}
}

func makeCord(spot []int) string {
	return fmt.Sprintf("%v,%v", strconv.Itoa(spot[0]), strconv.Itoa(spot[1]))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

/**

**/
