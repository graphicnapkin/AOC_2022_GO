package main

import (
	"AOC/day5/input"
	"fmt"
	"strconv"
	"strings"
)

// this returns slighlty less than it should, not sure why
func main() {
	//[testData, realData]
	data, _:= input.Data()
  fmt.Println(len(data))

	part1(data)
	part2(data)
}

func part1(data []string) {
	grid := make(map[string]int)
	dangerousCords := 0

	for _, row := range data {
		start, end := getCoordinates(row)

		if start.x == end.x {
			currentY := start.y

			for currentY >= end.y {
				stringPos := fmt.Sprintf("%v,%v", start.x, currentY)
				grid[stringPos] = grid[stringPos] + 1
				currentY--
			}

			currentY = start.y

			for currentY <= end.y {
				stringPos := fmt.Sprintf("%v,%v", start.x, currentY)
				grid[stringPos] = grid[stringPos] + 1
				currentY++
			}

      continue
		}

		if start.y == end.y {
			currentX := start.x

			for currentX >= end.x {
				stringPos := fmt.Sprintf("%v,%v", currentX, start.y)
				grid[stringPos] = grid[stringPos] + 1
				currentX--
			}

			currentX = start.x

			for currentX <= end.x {
				stringPos := fmt.Sprintf("%v,%v", currentX, start.y)
				grid[stringPos] = grid[stringPos] + 1
				currentX++
			}

      continue
		}

    currentX := start.x
    currentY := start.y

    for currentX != end.x || currentY != end.y {
      stringPos := fmt.Sprintf("%v,%v", currentX, currentY)
      grid[stringPos] = grid[stringPos] + 1

      if currentX < end.x{
        currentX++
      } else {
        currentX--
      }

      if currentY < end.y{
        currentY++
      } else {
        currentY--
      }
    }
	}

	for _, v := range grid {
		if v > 1 {
			dangerousCords++
		}
	}
	fmt.Println("Number of dangerous coordinates:", dangerousCords)
}

func part2(data []string) {
	output := 0
	fmt.Println(output)
}

func getCoordinates(row string) (Coordinate, Coordinate) {
	items := strings.Split(row, " -> ")

	firstCordString := strings.Split(items[0], ",")
	firstX, _ := strconv.Atoi(firstCordString[0])
	firstY, _ := strconv.Atoi(firstCordString[1])

	secondCordString := strings.Split(items[1], ",")
	secondX, _ := strconv.Atoi(secondCordString[0])
	secondY, _ := strconv.Atoi(secondCordString[1])

	return Coordinate{
			x: firstX,
			y: firstY,
		}, Coordinate{
			x: secondX,
			y: secondY,
		}
}

type Coordinate struct {
	x int
	y int
}
