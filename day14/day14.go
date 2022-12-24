package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//[testData, realData]
	data, _ := input()
	part1(data)
	part2()
}

func part1(data []string) {
	fmt.Println("Part 1")

	//setup grid
	gridSize := 15
	grid := make([][]string, 10)
	for i := range grid {
		row := []string{}
		for j := 0; j < gridSize; j++ {
			row = append(row, ".")
		}
		grid[i] = row
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	//parse input
	for _, row := range data {

		coords := []coord{}
		items := strings.Split(row, " -> ")
		for _, item := range items {
			cordPair := strings.Split(item, ",")
			x, _ := strconv.Atoi(cordPair[0])
			y, _ := strconv.Atoi(cordPair[1])
			coords = append(coords, coord{x - 490, y})
		}

		for i, cur := range coords {
			if i == len(coords)-1 {
				continue
			}
			next := coords[i+1]

			// if x's are the same, draw vertical
			if cur.x == next.x {
				targetY := cur.y
				for targetY <= next.y {
					grid[targetY][cur.x] = "#"
					targetY++
				}

				for targetY >= next.y {
					grid[targetY][cur.x] = "#"
					targetY--
				}
			}

			// if y's are the same, horiztonal vertical
			if cur.y == next.y {
				targetX := cur.x
				for targetX <= next.x {
					grid[cur.y][targetX] = "#"
					targetX++
				}

				for targetX >= next.x {
					grid[cur.y][targetX] = "#"
					targetX--
				}
			}
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}
}

type coord struct {
	x int
	y int
}

func part2() {
	fmt.Println("Part 2")
}

func input() ([]string, []string) {
	test := openFile("./input/testInput.txt")
	data := openFile("./input/input.txt")
	return test, data
}

func openFile(fileName string) []string {
	data := []string{}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		data = append(data, s)
		s, e = Readln(r)
	}

	return data
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
