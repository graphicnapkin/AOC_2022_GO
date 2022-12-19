package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	grid = [][]node{}
)

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	startX := 0
	startY := 0

	grid = make([][]node, len(data))
	for y := 0; y < len(data); y++ {
		grid[y] = make([]node, len(data[0]))
		for x := 0; x < len(data[0]); x++ {
			location := data[y][x]
			node := node{
				height: int(location),
			}
			if location == 'S' {
				node.start = true
				node.height = int('a')
				startX = x
				startY = y
			}
			if location == 'E' {
				node.end = true
				node.height = int('z')
			}
			node.name = fmt.Sprint(x, ":", y)

			grid[y][x] = node

		}
	}

	for y, row := range grid {
		for x, node := range row {
			node.neighbors = getNeigbors(x, y, grid)
			grid[y][x] = node
		}
	}

	fmt.Println(BFS(startX, startY))
	part2(grid)
}

func BFS(x int, y int) int {
	visited := make(map[string]bool)
	start := grid[y][x]
	queue := []queueItem{{start, 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.node.end {
			return item.distance
		}

		for _, neighbor := range item.node.neighbors {
			node := grid[neighbor[1]][neighbor[0]]
			if !visited[node.name] {
				visited[node.name] = true
				queue = append(queue, queueItem{node, item.distance + 1})
			}
		}
	}
	return 0
}

func getNeigbors(x int, y int, grid [][]node) [][]int {
	neighbors := [][]int{}
	current := grid[y][x]
	// down
	if y-1 >= 0 {
		neighbor := grid[y-1][x]
		if neighbor.height-current.height <= 1 {
			neighbors = append(neighbors, []int{x, y - 1})
		}
	}

	//up
	if y+1 < len(grid) {
		neighbor := grid[y+1][x]
		if neighbor.height-current.height <= 1 {
			neighbors = append(neighbors, []int{x, y + 1})
		}
	}

	//left
	if x-1 >= 0 {
		neighbor := grid[y][x-1]
		if neighbor.height-current.height <= 1 {
			neighbors = append(neighbors, []int{x - 1, y})
		}
	}

	//right
	if x+1 < len(grid[0]) {
		neighbor := grid[y][x+1]
		if neighbor.height-current.height <= 1 {
			neighbors = append(neighbors, []int{x + 1, y})
		}
	}

	return neighbors
}

func part2(grid [][]node) {
	fmt.Println("Part 2")
	lowPoints := [][]int{}
	lowest := 100000000

	for y, row := range grid {
		for x, node := range row {
			if node.height == 97 {
				lowPoints = append(lowPoints, []int{x, y})
			}
		}
	}

	for _, point := range lowPoints {
		result := BFS(point[0], point[1])
		if result != 0 && result < lowest {
			lowest = result
		}
	}

	fmt.Println(lowest)
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

type node struct {
	name      string
	start     bool
	end       bool
	height    int
	neighbors [][]int
}
type queueItem struct {
	node
	distance int
}
