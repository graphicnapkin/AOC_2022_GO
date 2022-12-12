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

var (
	cycle    int
	xReg     int
	readings []int
	screen   []string
	width    = 40
	height   = 6
)

func part1(data []string) {
	fmt.Println("Part 1")
	cycle = 1
	xReg = 1

	for _, row := range data {
		if row == "noop" {
			processCycle()
		} else {
			val, _ := strconv.Atoi(strings.Split(row, " ")[1])
			processCycle()
			processCycle()
			xReg += val
		}
	}
	total := 0
	for _, reading := range readings {
		total += reading
	}
	fmt.Println(total)
}

func processCycle() {
	cycle++
	if cycle == 20 || (cycle-20)%width == 0 {
		readings = append(readings, cycle*xReg)
	}
}

func part2(data []string) {
	fmt.Println("Part 2")
	screen = make([]string, width*height)
	for i := 0; i < width*height; i++ {
		screen[i] = "."
	}

	cycle = 0
	xReg = 1

	for _, row := range data {
		if row == "noop" {
			drawScreen(screen)
		} else {
			val, _ := strconv.Atoi(strings.Split(row, " ")[1])
			drawScreen(screen)
			drawScreen(screen)
			xReg += val
		}
	}
	for y := 0; y < height; y++ {
		fmt.Print(screen[width*y : width*(y+1)-1])
		fmt.Println(" ", width*y, "	", width*(y+1)-1)
	}
}

func drawScreen(screen []string) {
	pixelPos := getXPosition(cycle, 0)

	if pixelPos == xReg-1 || pixelPos == xReg || pixelPos == xReg+1 {
		screen[cycle] = "#"
	}
	cycle++
}

func getXPosition(pixel int, y int) int {
	if pixel < (y+1)*width {
		return pixel - (y * width)
	}
	return getXPosition(pixel, y+1)
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
