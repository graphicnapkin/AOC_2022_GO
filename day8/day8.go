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
	part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	visbileTrees := 0

	for y, row := range data {
		for x, tree := range row {
			//count all edge trees
			if x == 0 || x == len(row)-1 || y == 0 || y == len(data)-1 {
				visbileTrees++
				continue
			}

			spot, _ := strconv.Atoi(string(tree))
			up := checkDirection(data, x, y, spot, "up")
			down := checkDirection(data, x, y, spot, "down")
			left := checkDirection(data, x, y, spot, "left")
			right := checkDirection(data, x, y, spot, "right")

			if up || down || left || right {
				visbileTrees++
			}
		}
	}
	fmt.Println(visbileTrees)
}

func checkDirection(data []string, x int, y int, height int, dir string) bool {
	if dir == "up" {
		if y == 0 {
			return true
		}
		nextSpot, _ := strconv.Atoi(string(data[y-1][x]))
		if nextSpot >= height {
			return false
		}
		return checkDirection(data, x, y-1, height, "up")
	}

	if dir == "down" {
		if y == len(data)-1 {
			return true
		}
		nextSpot, _ := strconv.Atoi(string(data[y+1][x]))
		if nextSpot >= height {
			return false
		}
		return checkDirection(data, x, y+1, height, "down")
	}

	if dir == "left" {
		if x == 0 {
			return true
		}
		nextSpot, _ := strconv.Atoi(string(data[y][x-1]))
		if nextSpot >= height {
			return false
		}
		return checkDirection(data, x-1, y, height, "left")
	}

	if dir == "right" {
		if x == len(data)-1 {
			return true
		}
		nextSpot, _ := strconv.Atoi(string(data[y][x+1]))
		if nextSpot >= height {
			return false
		}
		return checkDirection(data, x+1, y, height, "right")
	}
	return false
}

func part2(data []string) {
	fmt.Println("Part 2")
	bestScore := 0

	for y, row := range data {
		for x, tree := range row {
			if x == 0 || x == len(row)-1 || y == 0 || y == len(data)-1 {
				continue
			}

			spot, _ := strconv.Atoi(string(tree))
			upScore := getDirectionScore(data, x, y, spot, 0, "up")
			downScore := getDirectionScore(data, x, y, spot, 0, "down")
			leftScore := getDirectionScore(data, x, y, spot, 0, "left")
			rightScore := getDirectionScore(data, x, y, spot, 0, "right")

			score := upScore * downScore * leftScore * rightScore
			if score > bestScore {
				bestScore = score
			}
		}
	}
	fmt.Println(bestScore)
}

func getDirectionScore(data []string, x int, y int, height int, score int, dir string) int {
	if dir == "up" {
		if y == 0 {
			return score
		}
		nextSpot, _ := strconv.Atoi(string(data[y-1][x]))
		if nextSpot >= height {
			return score + 1
		}
		return getDirectionScore(data, x, y-1, height, score+1, "up")
	}

	if dir == "down" {
		if y == len(data)-1 {
			return score
		}
		nextSpot, _ := strconv.Atoi(string(data[y+1][x]))
		if nextSpot >= height {
			return score + 1
		}
		return getDirectionScore(data, x, y+1, height, score+1, "down")
	}

	if dir == "left" {
		if x == 0 {
			return score
		}
		nextSpot, _ := strconv.Atoi(string(data[y][x-1]))
		if nextSpot >= height {
			return score + 1
		}
		return getDirectionScore(data, x-1, y, height, score+1, "left")
	}

	if dir == "right" {
		if x == len(data)-1 {
			return score
		}
		nextSpot, _ := strconv.Atoi(string(data[y][x+1]))
		if nextSpot >= height {
			return score + 1
		}
		return getDirectionScore(data, x+1, y, height, score+1, "right")
	}
	return 1
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
