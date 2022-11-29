package main

import (
	"AOC/day8/input"
	"fmt"
	"strings"
)

func main(){
  //[testData, realData]
  data,_:= input.Data()
  instructions := []Instruction{}

  for _, row := range data {
    instructions = append(instructions, parseInstructions(row))
  }

  part1(instructions)
  part2(instructions)
}


func parseInstructions(row string) Instruction{
  instruction := Instruction{}
  splitInput := strings.Split(row," | ")

  instruction.signalPatterns = strings.Split(splitInput[0]," ")
  instruction.output = strings.Split(splitInput[1]," ")

  return instruction
}

func part2(data []Instruction){
  fmt.Println("Part 2")
  for _, instruction := range data {
    currentCell := DisplayCell{}
    currentCell.topOptions = make(map[string]bool)
    currentCell.bottomOptions = make(map[string]bool)
    currentCell.middleOptions = make(map[string]bool)
    currentCell.topRightOptions = make(map[string]bool)
    currentCell.topLeftOptions = make(map[string]bool)
    currentCell.bottomRightOptions = make(map[string]bool)
    currentCell.bottomLeftOptions = make(map[string]bool)

    len6RunesToCheck := []rune{}
    // set options for unknown characters in "1"
    for _, characters := range instruction.signalPatterns {
      if len(characters) == 2 {
        for _, character := range characters {
          currentCell.topRightOptions[string(character)] = true
          currentCell.bottomRightOptions[string(character)] = true
        }
      }
    }

    // set options for unknown characters in "7"
    for _, characters := range instruction.signalPatterns {
      if len(characters) == 3 {
        for _, character := range characters {
          if !currentCell.topRightOptions[string(character)]{
            currentCell.top = string(character)
            currentCell.topOptions[string(character)] = true
          }
        }
      }
    }

    // set options for unknown characters in "4"
    for _, characters := range instruction.signalPatterns {
      if len(characters) == 4 {
        for _, character := range characters {
          if !currentCell.topRightOptions[string(character)]{
            currentCell.topLeftOptions[string(character)] = true
            currentCell.middleOptions[string(character)] = true
          }
        }
      }
    }

    // check if 6 length is the "9" character and assign the bottom value
    for _, characters := range instruction.signalPatterns {
      if len(characters) == 6 {
        isNine := true
        len6RunesToCheck = append(len6RunesToCheck, rune(currentCell.top[0]))

        //add both topright and top left options (all of the "4" characters)
        for k,_ := range currentCell.topRightOptions {
          len6RunesToCheck = append(len6RunesToCheck, rune(k[0]))
        }
        for k,_ := range currentCell.topLeftOptions {
          len6RunesToCheck = append(len6RunesToCheck, rune(k[0]))
        }

        for _, char := range len6RunesToCheck {
          if !contains[rune]([]rune(characters), char){
            isNine = false
          }
        }

        if isNine == true {
          for _, character := range characters {
            if !contains[rune](len6RunesToCheck,character){
              currentCell.bottom = string(character)
              len6RunesToCheck = append(len6RunesToCheck, character)
              break
            }
          }
        }
      }
    }

    for _, characters := range instruction.signalPatterns {
      // if we come across len 6 at this point and bottom is defined we know this is the "6" character
      // TODO PICKUP HERE: This could be either "0" or "6". I assumed it was 6 but it could be 0 too.
      if len(characters) == 6 && currentCell.bottom != ""{
        for _, character := range characters {
          if !contains[rune](len6RunesToCheck,character){
            currentCell.bottomLeft = string(character)
          }
        }
      }
    }
    fmt.Printf("%+v\n",currentCell)

  }
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

func contains[T comparable](s []T, str T) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

type Instruction struct {
  signalPatterns []string
  output []string
}

type DisplayCell struct {
  top string
  bottom string
  middle string
  topRight string
  topLeft string
  bottomRight string
  bottomLeft string
  topOptions map[string]bool
  bottomOptions map[string]bool
  middleOptions map[string]bool
  topRightOptions map[string]bool
  topLeftOptions map[string]bool
  bottomRightOptions map[string]bool
  bottomLeftOptions map[string]bool
}
