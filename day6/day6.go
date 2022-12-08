package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	for i := 0; i < len(data[0]); i++ {
		tempMap := make(map[string]bool)

		for j := 0; j < 4; j++ {
			tempMap[string(data[0][i+j])] = true
		}

		if len(tempMap) == 4 {
			fmt.Println(i + 4)
			return
		}
	}
}

func part2(data []string) {
	fmt.Println("Part 2")
	for i := 0; i < len(data[0]); i++ {
		tempMap := make(map[string]bool)

		for j := 0; j < 14; j++ {
			tempMap[string(data[0][i+j])] = true
		}

		if len(tempMap) == 14 {
			fmt.Println(i + 14)
			return
		}
	}
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
Advent of Code[About][Events][Shop][Settings][Log Out]graphicnapkin 10*
   int y=2022;[Calendar][AoC++][Sponsors][Leaderboard][Stats]
Our sponsors help make Advent of Code possible:
Axis - All we want for Christmas is your application, pls! ----------------- <embedded, cloud, Machine learning, fullstack> our cameras require it all
--- Day 6: Tuning Trouble ---
The preparations are finally complete; you and the Elves leave camp on foot and begin to make your way toward the star fruit grove.

As you move through the dense undergrowth, one of the Elves gives you a handheld device. He says that it has many fancy features, but the most important one to set up right now is the communication system.

However, because he's heard you have significant experience dealing with signal-based systems, he convinced the other Elves that it would be okay to give you their one malfunctioning device - surely you'll have no problem fixing it.

As if inspired by comedic timing, the device emits a few colorful sparks.

To be able to communicate with the Elves, the device needs to lock on to their signal. The signal is a series of seemingly-random characters that the device receives one at a time.

To fix the communication system, you need to add a subroutine to the device that detects a start-of-packet marker in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a sequence of four characters that are all different.

The device will send your subroutine a datastream buffer (your puzzle input); your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.

For example, suppose you receive the following datastream buffer:

mjqjpqmgbljsphdztnvjfqwrcgsmlb
After the first three characters (mjq) have been received, there haven't been enough characters received yet to find the marker. The first time a marker could occur is after the fourth character is received, making the most recent four characters mjqj. Because j is repeated, this isn't a marker.

The first time a marker appears is after the seventh character arrives. Once it does, the last four characters received are jpqm, which are all different. In this case, your subroutine should report the value 7, because the first start-of-packet marker is complete after 7 characters have been processed.

Here are a few more examples:

bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 5
nppdvjthqldpwncqszvftbrmjlhg: first marker after character 6
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 10
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 11
How many characters need to be processed before the first start-of-packet marker is detected?

To begin, get your puzzle input.

Answer:


You can also [Share] this puzzle.
**/
