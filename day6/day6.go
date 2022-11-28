package main

import (
	"AOC/day6/input"
	"fmt"
	"strconv"
	"strings"
)

func main(){
  //map the number of fish at each day marker
  //[testData, realData]
  _, data:= input.Data()

  part1(data)
  part2(data)
}

func part1(data []string){
  fmt.Println("Part 1")
  calculateFishPopulaton(data,80)
}


func part2(data []string){
  fmt.Println("Part 2")
  calculateFishPopulaton(data,256)
}


func calculateFishPopulaton (data []string, days int) {
  totalFish := 0
  school := map[int]int{
    0:0,
    1:0,
    2:0,
    3:0,
    4:0,
    5:0,
    6:0,
    7:0,
    8:0,
  }

  for _, fishString := range strings.Split(data[0],","){
    fish,_ := strconv.Atoi(fishString)
    school[fish] = school[fish] + 1
  }

  for i := 0; i < days; i++ {
    newSchool := make(map[int]int)

    newSchool[0] = school[1]
    newSchool[1] = school[2]
    newSchool[2] = school[3]
    newSchool[3] = school[4]
    newSchool[4] = school[5]
    newSchool[5] = school[6]
    newSchool[6] = school[7] + school[0]
    newSchool[7] = school[8]
    newSchool[8] = school[0]

    school = newSchool
  }

  for _, fish := range school{
    totalFish += fish
  }
  fmt.Println("Total",totalFish)
}


