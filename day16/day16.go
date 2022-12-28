package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	tunnels          = make(map[string]tunnel)
	distanceToTunnel = make(map[string]int)
)

func main() {
	//[testData, realData]
	//data, _ := input()
	_, data := input()
	part1(data)
	part2()
}

func part1(data []string) {
	fmt.Println("Part 1")

	start := tunnel{}
	answer := 0
	path := "AA"

	start = makeTunnels(data, start)

	setDistances()

	for _, tun := range tunnels {
		if tun.rate == 0 {
			delete(tunnels, tun.name)
		}
	}

	path, answer = makeMoves(start, path, answer)
	fmt.Println(answer)
	fmt.Println(path)
}

func makeMoves(start tunnel, path string, answer int) (string, int) {
	queue := []Move{
		{start, tunnels, 30, 0, path},
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		newQueueItems, pressure, newPath := takeMove(cur)
		queue = append(queue, newQueueItems...)
		if pressure > answer {
			answer = pressure
			path = newPath
		}
	}
	return path, answer
}

func setDistances() {
	for _, tun := range tunnels {
		distQueue := []distQueue{{tun.name, tun, 0, make(map[string]bool)}}

		for len(distQueue) > 0 {
			cur := distQueue[0]
			distQueue = distQueue[1:]

			nextQueue := setDistance(cur.source, cur.tunnels, cur.distance, cur.visited)
			distQueue = append(distQueue, nextQueue...)
		}
	}
}

func makeTunnels(data []string, start tunnel) tunnel {
	for _, row := range data {
		row = strings.Replace(row, "Valve ", "", 1)
		row = strings.Replace(row, " has flow rate=", ",", 1)
		row = strings.Replace(row, "; tunnels lead to valves ", ",", 1)
		row = strings.Replace(row, "; tunnel leads to valve ", ",", 1)
		row = strings.ReplaceAll(row, " ", "")

		items := strings.Split(row, ",")
		rate, _ := strconv.Atoi(items[1])
		tun := tunnel{
			name:      items[0],
			rate:      rate,
			neighbors: items[2:],
		}
		if tun.name == "AA" {
			start = tun
			tun.activated = true
		}
		tunnels[tun.name] = tun
	}
	return start
}

func takeMove(move Move) ([]Move, int, string) {
	queue := []Move{}
	for _, tun := range move.tunnels {
		keyName := ""
		if move.tun.name < tun.name {
			keyName = move.tun.name + ":" + tun.name
		} else {
			keyName = tun.name + ":" + move.tun.name
		}
		distance := distanceToTunnel[keyName]

		if move.mins > distance+1 {
			tunnelsCopy := make(map[string]tunnel)
			for k, v := range move.tunnels {
				if k != tun.name {
					tunnelsCopy[k] = v
				}
			}
			queue = append(queue, Move{
				tun,
				tunnelsCopy,
				//move there and activate
				move.mins - distance - 1,
				move.pressure + tun.rate*(move.mins-distance-1),
				move.path + "->" + tun.name,
			})

		}
	}
	return queue, move.pressure, move.path
}

func setDistance(sourceTunnel string, tun tunnel, currentDistance int, visited map[string]bool) []distQueue {
	queue := []distQueue{}
	if visited[tun.name] {
		return queue
	}

	visited[tun.name] = true

	if tun.name != sourceTunnel {
		keyName := ""
		if sourceTunnel < tun.name {
			keyName = sourceTunnel + ":" + tun.name
		} else {
			keyName = tun.name + ":" + sourceTunnel

		}
		distanceToTunnel[keyName] = currentDistance
	}

	for _, neighbor := range tun.neighbors {
		queue = append(queue, distQueue{sourceTunnel, tunnels[neighbor], currentDistance + 1, visited})
	}
	return queue
}

type (
	tunnel struct {
		name      string
		activated bool
		rate      int
		neighbors []string
	}

	Move struct {
		tun      tunnel
		tunnels  map[string]tunnel
		mins     int
		pressure int
		path     string
	}

	distQueue struct {
		source   string
		tunnels  tunnel
		distance int
		visited  map[string]bool
	}
)

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
