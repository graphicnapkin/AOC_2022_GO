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
	monkeyMap = make(map[int]monkey)
	modulus   = 1
)

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
	part2(data)
}

func part1(data []string) {
	fmt.Println("Part 1")
	monkeyMap = buildMonkeyMap(data)
	//part 1 was 20 rounds
	//rounds := 20
	rounds := 10000

	for _, monkey := range monkeyMap {
		modulus *= monkey.testVal.agaist
	}

	for r := 0; r < rounds; r++ {
		numMonkeys := len(monkeyMap)
		for i := 0; i < numMonkeys; i++ {
			m := monkeyMap[i]
			items := len(m.inventory)
			for j := 0; j < items; j++ {
				m.inspect()
				m.test()

				//fmt.Print("--------------------------\n\n\n")
			}
		}
		if r == 0 || r == 19 || r == 999 {
			fmt.Println("After round", r+1, ", the monkeys are holding items with these worry levels:")

			for k, v := range monkeyMap {
				fmt.Println("Monkey ", k, "inspected items", v.inspectedTotal, "times.")
			}
		}

	}

	topTwo := []int{0, 0}

	for _, v := range monkeyMap {
		if v.inspectedTotal > topTwo[0] {
			currentTop := topTwo[0]
			topTwo[0] = v.inspectedTotal
			if currentTop > topTwo[1] {
				topTwo[1] = currentTop
			}
		} else {
			if v.inspectedTotal > topTwo[1] {
				topTwo[1] = v.inspectedTotal
			}
		}
	}

	fmt.Println("Total Monkey Business:", topTwo[0]*topTwo[1])
}

func buildMonkeyMap(data []string) map[int]monkey {
	monkeyMap := map[int]monkey{}
	for i := 0; i < len(data); i++ {
		newMonkey := newMonkey(data[i : i+6])
		monkeyMap[newMonkey.name] = newMonkey
		i += 6
	}
	return monkeyMap
}

func newMonkey(data []string) monkey {
	var m monkey
	m.name, _ = strconv.Atoi(strings.Split(strings.Replace(data[0], ":", "", 1), " ")[1])
	inventory := strings.Split(data[1], " ")[4:]
	for _, itemString := range inventory {
		item := strings.Replace(itemString, ",", "", 1)
		intItem, _ := strconv.Atoi(item)
		m.inventory = append(m.inventory, intItem)
	}
	m.operation.sign = strings.Split(data[2], " ")[6]
	m.operation.value = strings.Split(data[2], " ")[7]

	m.testVal.agaist, _ = strconv.Atoi(strings.Split(data[3], " ")[5])

	m.testVal.pass, _ = strconv.Atoi(strings.Split(data[4], " ")[9])
	m.testVal.fail, _ = strconv.Atoi(strings.Split(data[5], " ")[9])
	//fmt.Printf("%+v\n", m)

	return m
}

func (m *monkey) inspect() {
	m.inspectedTotal++
	var modifier int
	item := m.inventory[0]
	//fmt.Println("Monkey", m.name, "inpsects an item with worry level of", item)
	if m.operation.value == "old" {
		modifier = item
	} else {
		num, _ := strconv.Atoi(m.operation.value)
		modifier = num
	}

	if m.operation.sign == "*" {
		item = (item * modifier)
		//fmt.Println("	Worry level is multiplied by", modifier, "to", item)
		//below is only valid for part 1
		//item /= 3
		//fmt.Println("	Monkey gets bored with item. Worry level is divided by 3 to", item)
		m.inventory[0] = item % modulus
	} else {
		item += modifier
		//fmt.Println("	Worry level increases by", modifier, "to", item)
		//below is only valid for part 1
		//item /= 3
		//fmt.Println("	Monkey gets bored with item. Worry level is divided by 3 to", item)
		m.inventory[0] = item % modulus
	}
	monkeyMap[m.name] = *m
}

func (m *monkey) test() {
	item := m.inventory[0]

	if len(m.inventory) > 0 {
		m.inventory = m.inventory[1:]
	} else {
		m.inventory = []int{}
	}
	var targetMonkey monkey
	var monkeyIndex int

	if item%m.testVal.agaist == 0 {
		monkeyIndex = m.testVal.pass
		//fmt.Println("	Worry level is divisible by", m.testVal.pass)
		targetMonkey = monkeyMap[monkeyIndex]
	} else {
		//fmt.Println("	Worry level is NOT divisible by", m.testVal.fail)
		monkeyIndex = m.testVal.fail
		targetMonkey = monkeyMap[monkeyIndex]
	}

	targetMonkey.inventory = append(targetMonkey.inventory, item)
	//fmt.Println("	Item with worry level", item, "is thrown to monkey", monkeyIndex)
	monkeyMap[monkeyIndex] = targetMonkey
	monkeyMap[m.name] = *m
}

func part2(data []string) {
	fmt.Println("Part 2")
}

type monkey struct {
	name      int
	inventory []int
	operation struct {
		sign  string
		value string
	}
	testVal struct {
		agaist int
		pass   int
		fail   int
	}
	inspectedTotal int
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
