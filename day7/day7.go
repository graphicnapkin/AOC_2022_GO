package main

import (
	"AOC/day7/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
  //[testData, realData]
  _, data:= input.Data()

  part1(data)
  part2(data)
}

func part1(data []string){
  fmt.Println("Part 1")
  calculateFuelCost(data, "1")
}

func part2(data []string){
  fmt.Println("Part 2")
  calculateFuelCost(data,"2")
}

func calculateFuelCost(data []string, part string){
  crabPostitions := []int{}

  min := math.MaxInt
  max := 0

  for _, stringPos := range strings.Split(data[0],",") {
    pos, _ := strconv.Atoi(stringPos)
    if pos < min {
      min = pos
    }
    if pos > max {
      max = pos
    }
    crabPostitions = append(crabPostitions, pos)
  }

  bestFuelSpend := math.MaxInt

  for i := min; i < max; i++{
    posFuelSpend := 0
    for _, crabPos := range crabPostitions {
      if (part == "1") {
        posFuelSpend += Abs(crabPos -i)
      } else {
        posFuelSpend += realFuelCost(Abs(crabPos - i))
      }
    }
    if posFuelSpend < bestFuelSpend {
      bestFuelSpend = posFuelSpend
    }
  }

  fmt.Println("Answer:",bestFuelSpend)
}


func realFuelCost(num int) int {
  total := 0
  for i := 0; i <= num; i++ {
    total += i
  }
  return total
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

