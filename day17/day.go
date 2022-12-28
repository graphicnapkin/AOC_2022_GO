package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	grid = make(map[string]string)
	//shape offset values, starting from bottom left
	shapes = [][][]int{
		//flat
		{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		//cross
		{
			{1, 0},
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 2},
		},
		//backwards L
		{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
		},
		//vertical
		{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
		//square
		{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
	}
	drawTop      = 10
	shapeTop     = 0
	currentShape = [][]int{}
)

func main() {
	//[testData, realData]
	data, _ := input()
	//_, data := input()
	part1(data)
	part2()
}

func part1(data []string) {
	fmt.Println("Part 1")
	fallenRocks := 0
	windDirection := 0
	for fallenRocks < 2023 {
		for i := 0; i < 5; i++ {
			addPiece(i)
			printGrid()
			done := false
			for !done {
				windMove(windDirection, data[0])
				printGrid()
				windDirection++
				done = moveDown()
				printGrid()
			}
		}
	}
}

func windMove(index int, data string) {
	posMap := make(map[string]bool)
	newPos := [][]int{}
	direction := string(data[index%(len(data)-1)])
	fmt.Println(direction)

	//map all postions
	for _, xy := range currentShape {
		pos := fmt.Sprint(xy[0], ",", xy[1])
		posMap[pos] = true
	}

	//prevent out of bounds
	for _, xy := range currentShape {
		pos := ""
		if direction == "<" {
			if xy[0]-1 < 1 {
				return
			}
			pos = fmt.Sprint(xy[0]-1, ",", xy[1])
		} else {
			if xy[0]+1 > 7 {
				return
			}
			pos = fmt.Sprint(xy[0]+1, ",", xy[1])
		}
		if !posMap[pos] && grid[pos] == "#" {
			return
		}
	}

	//move position
	for _, xy := range currentShape {
		cur := fmt.Sprint(xy[0], ",", xy[1])
		new := ""
		old := ""
		if direction == "<" {
			new = fmt.Sprint(xy[0]-1, ",", xy[1])
			old = fmt.Sprint(xy[0]+1, ",", xy[1])
			newPos = append(newPos, []int{xy[0] - 1, xy[1]})
		} else {
			new = fmt.Sprint(xy[0]+1, ",", xy[1])
			old = fmt.Sprint(xy[0]-1, ",", xy[1])
			newPos = append(newPos, []int{xy[0] + 1, xy[1]})
		}
		grid[new] = "#"
		grid[cur] = grid[old]
	}
	currentShape = newPos
}

func moveDown() bool {
	posMap := make(map[string]bool)
	newPos := [][]int{}
	top := 0

	//map all postions
	for _, xy := range currentShape {
		if xy[1]-1 <= 0 {
			return true
		}
		pos := fmt.Sprint(xy[0], ",", xy[1])
		posMap[pos] = true
	}
	//check spots below, exludig positions in current shape
	for _, xy := range currentShape {
		pos := fmt.Sprint(xy[0], ",", xy[1]-1)
		if !posMap[pos] && grid[pos] == "#" {
			return true
		}
	}

	//move each xy down, and whatever was above it down
	for _, xy := range currentShape {
		cur := fmt.Sprint(xy[0], ",", xy[1])
		above := fmt.Sprint(xy[0], ",", xy[1]+1)
		below := fmt.Sprint(xy[0], ",", xy[1]-1)
		newPos = append(newPos, []int{xy[0], xy[1] - 1})
		if xy[1]-1 > top {
			top = xy[1] - 1
		}

		grid[below] = grid[cur]
		grid[cur] = grid[above]
	}
	currentShape = newPos
	shapeTop = top
	fmt.Println("Shape top", shapeTop)
	//printGrid()
	return false
}

func addPiece(idx int) {
	newTop := 0
	shapeCoords := [][]int{}
	for _, xy := range shapes[idx] {
		x := xy[0] + 2
		y := xy[1] + shapeTop + 4
		shapeCoords = append(shapeCoords, []int{x, y})
		pos := fmt.Sprint(x, ",", y)
		grid[pos] = "#"
		if y > newTop {
			newTop = y
		}
	}
	if newTop+1 > drawTop {
		drawTop = newTop + 6
	}
	currentShape = shapeCoords
	fmt.Println("Shape top", shapeTop)
	//printGrid()
}

func printGrid() {
	display := []string{}
	for y := 1; y <= drawTop; y++ {
		row := ""
		for x := 1; x <= 7; x++ {
			pos := fmt.Sprint(x, ",", y)
			if val, ok := grid[pos]; !ok || val == "" {
				row += " "
			} else {
				row += val
			}
		}
		display = append(display, row)
	}
	for i := len(display) - 1; i > -1; i-- {
		index := fmt.Sprint(i + 1)
		for len(index) < 3 {
			index = " " + index
		}
		fmt.Println(index, display[i], "|")
	}
	fmt.Println("    1234567")
	time.Sleep(500 * time.Millisecond)
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
