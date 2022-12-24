package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	correctOrder := 0
	fmt.Println("Part 1")
	row := 1
	for i := 0; i < len(data); i++ {
		left, _ := getData(data[i])
		right, _ := getData(data[i+1])

		_, correct := compareItems(left, right)
		if correct {
			correctOrder += row
		}

		i += 2
		row++
	}

	fmt.Println("Answer:", correctOrder)
}

// returns if we should continue, and if it's correct
func compareItems(left interface{}, right interface{}) (bool, bool) {
	leftType := getType(left)
	rightType := getType(right)

	if leftType == "float64" && rightType == "float64" {
		if left.(float64) < right.(float64) {
			return false, true
		}
		if left.(float64) > right.(float64) {
			return false, false
		}
		return true, false
	}

	if leftType == "[]interface {}" && rightType == "[]interface {}" {
		leftList := left.([]interface{})
		rightList := right.([]interface{})

		if len(leftList) == 0 && len(rightList) == 0 {
			return true, false
		}

		if len(leftList) == 0 && len(rightList) > 0 {
			return false, true
		}

		if len(rightList) == 0 {
			return false, false
		}

		for i := 0; i < len(leftList); i++ {
			if i > len(rightList)-1 {
				return false, false
			}

			cont, correct := compareItems(leftList[i], rightList[i])
			if !cont {
				return false, correct
			}
			if i == len(leftList)-1 && len(rightList) > len(leftList) {
				return false, true
			}
		}
		// equal slices
		return true, false
	}

	if leftType == "float64" {
		leftList := []interface{}{left}
		return compareItems(leftList, right)
	}

	// right is float, and left is list
	rightList := []interface{}{right}
	return compareItems(left, rightList)
}

func getData(row string) (interface{}, error) {
	var items interface{}
	err := json.Unmarshal([]byte(row), &items)
	if err != nil {

		return items, err
	}
	return items, nil
}

func getType(obj interface{}) string {
	return fmt.Sprintf("%T", obj)
}

func part2(data []string) {
	fmt.Println("Part 2")
	rows := []interface{}{}

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			continue
		}
		row, _ := getData(data[i])
		rows = append(rows, row)
	}
	firstDivider, _ := getData("[[2]]")
	secondDivider, _ := getData("[[6]]")
	rows = append(rows, firstDivider)
	rows = append(rows, secondDivider)

	sort := true

	for sort {
		sort = false
		for i := 0; i < len(rows); i++ {
			if i == len(rows)-1 {
				continue
			}
			_, compare := compareItems(rows[i], rows[i+1])
			if !compare {
				first := rows[i]
				rows[i] = rows[i+1]
				rows[i+1] = first
				sort = true
			}
		}
	}

	firstLocation := 0
	secondLocation := 0

	for i, row := range rows {
		rowString := fmt.Sprint(row)
		if rowString == "[[2]]" {
			firstLocation = i + 1
		}
		if rowString == "[[6]]" {
			secondLocation = i + 1
		}
	}
	fmt.Println("Answer", firstLocation*secondLocation)
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
