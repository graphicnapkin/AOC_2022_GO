package main

import (
	"AOC/day1/input"
	"fmt"
)

func main(){
  _,data:= input.Data()
  //part1(data)
  part2(data)
}

func part1(data []int){
  output := 0
  previous := struct{
      val int
      set bool
    }{
      val: 0, 
      set: false,
    }
  for _, current := range data {
    //fmt.Println(previous)
    if previous.set && current > previous.val {
      output ++
    }
    previous.val = current
    if !previous.set {
      previous.set = true
    }
  }
  fmt.Println(output)
}

func part2(data []int){
  output := 0
  previousWindow := struct{
    val int
    set bool
  }{
    val: 0,
    set: false,
  }

  for i, current := range data {
    temp := current 

    if i + 1 <= len(data) - 1 {
      temp += data[i + 1]
    }
    
    if i + 2 <= len(data) - 1{
      temp += data[i + 2]
    }
    if previousWindow.set && previousWindow.val < temp {
      output ++
    }
    previousWindow.val = temp 
    if !previousWindow.set {previousWindow.set = true}
  }
  fmt.Println(output)
}