package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	//targetRow = 10
	targetRow = 2000000
	positions = make(map[string]bool)
	skipPos   = make(map[string]bool)
)

func main() {
	//[testData, realData]
	_, data := input()
	// data, _ := input()
	//part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")

	sensors := getSensors(data)
	for _, sensor := range sensors {
		checkPositions(sensor)
	}

	for k, _ := range skipPos {
		delete(positions, k)
	}

	fmt.Println("Answer", len(positions))
}
func part2(data []string) {
	fmt.Println("Part 2")
	boundries := [][]int{}
	sensors := getSensors(data)
	for i, sensor := range sensors {
		fmt.Println("Getting boundries for sensor:", i)
		boundries = append(boundries, getBoundry(sensor)...)
	}

	for _, boundry := range boundries {
		if boundry[0] < 0 || boundry[0] > 4000001 || boundry[1] < 0 || boundry[1] > 4000001 {
			continue
		}
		winningSpot := true
		for _, sensor := range sensors {
			if mDist(boundry[0], boundry[1], sensor.x, sensor.y) <= sensor.beaconDistance {
				winningSpot = false
			}
		}
		if winningSpot {
			fmt.Println("Answer:", boundry[0]*4000000+boundry[1])
			return
		}
	}
}

func getSensors(data []string) []sensor {
	sensors := []sensor{}
	for _, row := range data {
		row = strings.Replace(row, "Sensor at x=", "", 1)
		row = strings.Replace(row, " y=", "", 2)
		row = strings.Replace(row, " closest beacon is at x=", "", 1)
		items := strings.Split(row, ":")

		beaconPos := strings.Split(items[1], ",")
		beaconX, _ := strconv.Atoi(beaconPos[0])
		beaconY, _ := strconv.Atoi(beaconPos[1])

		if beaconY == targetRow {
			skipPos[fmt.Sprint(beaconX, beaconY)] = true
		}

		sensorPos := strings.Split(items[0], ",")
		x, _ := strconv.Atoi(sensorPos[0])
		y, _ := strconv.Atoi(sensorPos[1])

		distance := mDist(x, y, beaconX, beaconY)
		sensors = append(sensors, sensor{x, y, distance, beaconX, beaconY})
	}
	return sensors
}

func checkPositions(s sensor) {
	if targetRow > s.y+s.beaconDistance || targetRow < s.y-s.beaconDistance {
		return
	}

	diff := int(math.Abs(float64(s.y - targetRow)))
	startX := s.x - (s.beaconDistance - diff)
	endX := s.x + (s.beaconDistance - diff)

	for i := startX; i <= endX; i++ {
		pos := fmt.Sprint(i, targetRow)
		positions[pos] = true
	}
}

func getBoundry(s sensor) [][]int {
	boundry := [][]int{}

	top := []int{s.x, s.y - s.beaconDistance - 1}
	bottom := []int{s.x, s.y + s.beaconDistance + 1}
	left := []int{s.x - s.beaconDistance - 1, s.y}
	right := []int{s.x + s.beaconDistance + 1, s.y}

	current := []int{top[0], top[1]}
	for current[0] != right[0] {
		current[0]++
		current[1]++
		boundry = append(boundry, []int{current[0], current[1]})
	}

	current = []int{right[0], right[1]}
	for current[0] != bottom[0] {
		current[0]--
		current[1]++
		boundry = append(boundry, []int{current[0], current[1]})
	}

	current = []int{bottom[0], bottom[1]}
	for current[0] != left[0] {
		current[0]--
		current[1]--
		boundry = append(boundry, []int{current[0], current[1]})
	}

	current = []int{left[0], left[1]}
	for current[0] != top[0] {
		current[0]++
		current[1]--
		boundry = append(boundry, []int{current[0], current[1]})
	}
	return boundry
}

type sensor struct {
	x              int
	y              int
	beaconDistance int
	beaconX        int
	beaconY        int
}

func mDist(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
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
