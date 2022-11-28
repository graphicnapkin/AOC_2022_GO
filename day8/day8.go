package main

import (
	"AOC/day8/input"
	"fmt"
	"strings"
)

func main(){
  //[testData, realData]
  _,data:= input.Data()
  instructions := []Instruction{}

  for _, row := range data {
    instructions = append(instructions, parseInstructions(row))
  }

  part1(instructions)
  part2(instructions)
}

func part1(data []Instruction){
  fmt.Println("Part 1")
  screenNumberCount := make(map[int]int)
  output := 0


  for _, instruction := range data {
    for _, item := range instruction.output {
      itemLength := len(item)
      if itemLength == 2 {
        screenNumberCount[1] = screenNumberCount[1] + 1
      }
      if itemLength == 3 {
        screenNumberCount[7]++
      }
      if itemLength == 4 {
        screenNumberCount[4]++
      }
      if itemLength == 7 {
        screenNumberCount[8]++
      }
    }
  }
  fmt.Println(screenNumberCount)
  for k, num := range screenNumberCount {
    fmt.Printf("%v:%v\n",k,num)
    output += num
  }
  fmt.Println("Answer:",output)
}

func part2(data []Instruction){
  fmt.Println("Part 2")
}

func parseInstructions(row string) Instruction{
  instruction := Instruction{}
  splitInput := strings.Split(row," | ")

  instruction.signalPatterns = strings.Split(splitInput[0]," ")
  instruction.output = strings.Split(splitInput[1]," ")

  return instruction
}

type Instruction struct {
  signalPatterns []string
  output []string
}
