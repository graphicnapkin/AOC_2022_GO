package main

import (
	"AOC/day4/input"
	"fmt"
	"strconv"
	"strings"
)

func main(){
  //[testData, realData]
  data,_:= input.Data()
  part1(data)
  part2(data)
}

func part1(data []string){
  input := parseInput(data)
  for k, v := range input.boards[0].grid{
    fmt.Println(k,v)
  }
}

func part2(data []string){
}
//solution idea
//count how many times a X position has been called
//as soon as the same X position has been called however wide the board is, that board won
//same for Y position, count how many it's been called
//as soon as the same Y position has been called for as many rows of input there is in that board
//the board has won... this is a half sleepy idea and probably totally wrong.
//also thinking through the map, the KEY should be bingoNumbers, the VAlUE should be it's X and Y coordinates
//and a new MAP should be created with each unique X position as a key, and how many times it's been hit is the value
//and as you hit an X position, if it's been enough hits it wins, same with y
//again sleepy but I think this is actually the right solution

func parseInput(data []string) Input{
  var input Input
  //the inputs first row will always be the numbers pulled
  input.bingoNumbers = strings.Split(data[0], ",")

  tempStringBoard := []string{}
  var newBoard Board
  newBoard.grid = make(Grid)
  //the inputs second row will always be blank so start at the third row (index 2)
  for i := 2; i < len(data); i++ {
    if len(data[i]) == 0 {
      for y := 0; y < len(tempStringBoard); y++{
        //split by spaces
        //itterate though each position
        //save the data in a map with number in that position as the key, x,y pos as value
        //add this number to total basescore value
        }
      }
      input.boards = append(input.boards, newBoard)
      newBoard.grid = make(Grid)
    } else {
      tempStringBoard = append(tempStringBoard, data[i])
    }
  }
  return input
}

type Input struct{
  bingoNumbers []string
  boards []Board
}

type Grid map[string]struct{
  x int
  y int
} 

type Board struct {
  grid Grid
  baseScore int
}