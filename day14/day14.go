package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"time"
)

var minX = 10000
var maxX = 0
var maxY = 0

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
	part2()
}

func part1(data []string) {
	fmt.Println("Part 1")

	gridSize := 1000
	emptyRow := []string{}
	for i := 0; i < gridSize; i++ {
		emptyRow = append(emptyRow, ".")
	}

	//setup grid
	grid := make(Grid, gridSize)
	for i := range grid {
		row := []string{}
		grid[i] = append(row, emptyRow...)
	}

	setInitialState(grid, data)

	start := coord{500, 0}

	i := 0

	for {
		done := dropSand(grid, start, i)
		if done {
			break
		}
		i++
	}
	fmt.Println("Answer", i)
}

func setInitialState(grid [][]string, data []string) {
	for _, row := range data {
		coords := []coord{}
		items := strings.Split(row, " -> ")
		for _, item := range items {
			cordPair := strings.Split(item, ",")
			x, _ := strconv.Atoi(cordPair[0])
			y, _ := strconv.Atoi(cordPair[1])
			coords = append(coords, coord{x, y})
			if x < minX {
				minX = x
			}

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}
		}

		for i, cur := range coords {
			if i == len(coords)-1 {
				continue
			}
			next := coords[i+1]
			grid[cur.y][cur.x] = "#"
			grid[next.y][next.x] = "#"

			// if x's are the same, draw vertical
			if cur.x == next.x {
				targetY := cur.y
				for targetY < next.y {
					grid[targetY][cur.x] = "#"
					targetY++
				}

				for targetY > next.y {
					grid[targetY][cur.x] = "#"
					targetY--
				}
			}

			// if y's are the same, horiztonal vertical
			if cur.y == next.y {
				targetX := cur.x
				for targetX < next.x {
					grid[cur.y][targetX] = "#"
					targetX++
				}

				for targetX > next.x {
					grid[cur.y][targetX] = "#"
					targetX--
				}
			}
		}
	}
	lastRow := []string{}
	for i := 0; i < 1000; i++ {
		lastRow = append(lastRow, "#")
	}
	grid[maxY+2] = lastRow
}

type Grid [][]string

func (g *Grid) printGrid() {
	for j := 0; j < 20; j++ {
		fmt.Println("")
	}
	for i, row := range *g {
		if i <= maxY+2 {
			fmt.Println(row[minX-5 : maxX+10])
		}
	}
}

// returns boolean for done
func dropSand(grid Grid, cur coord, moves int) bool {
	// part 1
	//if cur.y + 1 > maxY || cur.x-1 < minX || cur.x > maxX{
	down := coord{cur.x, cur.y + 1}
	if grid.getSpot(down) == "." {
		return dropSand(grid, down, moves)
	}

	downLeft := coord{cur.x - 1, cur.y + 1}
	if grid.getSpot(downLeft) == "." {
		return dropSand(grid, downLeft, moves)
	}

	downRight := coord{cur.x + 1, cur.y + 1}
	if grid.getSpot(downRight) == "." {
		return dropSand(grid, downRight, moves)
	}

	if grid.getSpot(cur) == "o" {
		return true
	}

	grid.setSpot(cur, "o")
	return false
}

func (g Grid) setSpot(spot coord, val string) {
	g[spot.y][spot.x] = val
}

func (g Grid) getSpot(spot coord) string {
	if spot.x > 999 || spot.y > 999 {
		fmt.Println(spot)
	}
	return g[spot.y][spot.x]
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
