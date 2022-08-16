package main

import (
	"AOC/day2/input"
	"fmt"
	"strconv"
	"strings"
)

func main(){
  //[testData, realData]
  _,data:= input.Data()
  //part1(data)
  part2(data)
}

func part1(data []string){
  pos := position{}
  for _, instruction := range data {
    applyInstruction(&pos,instruction)
  }
  fmt.Println(pos.depthCalculation())
}

func part2(data []string){
  pos := position{}
  for _, instruction := range data {
    applyRealInstruction(&pos, instruction)
  }
  fmt.Println(pos.depthCalculation())
}

type position struct {
  x int64
  y int64
  aim int64
}

func applyInstruction(p *position, input string) {
  instruction := strings.Split(input, " ")
  direction := instruction[0]
  amount, err :=  strconv.ParseInt(instruction[1],0,8)
  if err != nil {
    return 
  }
  if direction == "forward" { p.x += amount }
  if direction == "up" { p.y -= amount}
  if direction == "down" { p.y += amount }
}

func applyRealInstruction(p *position, input string) {
  instruction := strings.Split(input, " ")
  direction := instruction[0]
  amount, err :=  strconv.ParseInt(instruction[1],0,8)
  if err != nil {
    return 
  }
  if direction == "forward" { 
    p.x += amount 
    p.y += amount * p.aim
  }
  if direction == "up" { p.aim -= amount}
  if direction == "down" { p.aim += amount}

}


func (p *position) depthCalculation() int64 {
  return p.x * p.y
}

