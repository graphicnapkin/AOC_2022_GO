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
	//_, data := input()
	part1(data)
	part2()
}

func part1(data []string) {
	fmt.Println("Part 1")
	tunnels := make(map[string]tunnel)
	distanceToTunnel := make(map[string]int)

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
		tunnels[tun.name] = tun
	}

	for _, v := range tunnels {
		fmt.Printf("%+v\n", v)
	}
}

type tunnel struct {
	name      string
	rate      int
	neighbors []string
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
