package main

import (
	"AOC/day3/input"
	"fmt"
	"strconv"
)

func main(){
  //[testData, realData]
  _,data:= input.Data()
  part1(data)
  part2(data)
}

func part1(data []string){
  //create a report map, that will store the counts for each bit per position
  report := make(map[int]map[int]int)

  //initialize each position to a default value
  for i := 0; i < len(data[0]); i++ {
    report[i] = map[int]int{
     0:0,
     1:0,
    }
  }

  gammaString := ""
  epsilonString := ""

  //for each row and each position count either 0 or 1
  for _, row := range data {
    for i, bit := range row {
      //note to self: rune is equivalent of 'char' and is an individual character in a string
      if bit == '0' {
        report[i][0]++
      } else {
        report[i][1]++
      }
    }
  }

  //compare each position
  for i := 0; i < len(data[0]); i++ {
    if report[i][0] > report[i][1]{
      gammaString += "0"
      epsilonString += "1"
    } else {
      gammaString += "1"
      epsilonString += "0"
    }
  }

  //convert binary strings to decimal
  gamma, err := strconv.ParseInt(gammaString,2,64)
  if err != nil {
    fmt.Println("error")
    return
  }
  epsilon, err := strconv.ParseInt(epsilonString,2,64)
  if err != nil {
    fmt.Println("error")
    return
  }
  //output per spec 
  fmt.Println(gamma * epsilon)
}

func part2(data []string){
  oList := data
  co2List := data

  //loop through the oList once per character in each entry
    //if oList len == 1 break
    //loop through each rating in oList
    //count how many 0's and 1's there are
    //keep only whichever char is most represented
    //if there is a tie, keep the 1's

    //for each rating
        //count how many 0's and 1's are in 
  for i := 0; i < len(oList[0]); i++ {
    if len(oList) == 1 {break}
    zeros := 0
    ones := 0
    zerosList := []string{}
    onesList := []string{}

    for _, rating := range oList {
      if(rating[i] == '0'){
        zeros ++
        zerosList = append(zerosList, rating)
      } else {
        ones ++
        onesList = append(onesList, rating)
      }
    }

    if zeros > ones {
      oList = zerosList
    } else {
      oList = onesList
    }
  }
  
  for i := 0; i < len(co2List[0]); i++ {
    if len(co2List) == 1 {break}
    zeros := 0
    ones := 0
    zerosList := []string{}
    onesList := []string{}

    for _, rating := range co2List {
      if(rating[i] == '0'){
        zeros ++
        zerosList = append(zerosList, rating)
      } else {
        ones ++
        onesList = append(onesList, rating)
      }
    }

    if ones < zeros {
      co2List = onesList 
    } else {
      co2List = zerosList
    }
  }
  
  oxygenRating, err := strconv.ParseInt(oList[0],2,64)
  if err != nil { 
    fmt.Println("error")
    return 
  }
  co2Rating, err := strconv.ParseInt(co2List[0],2,64)
  if err != nil { 
    fmt.Println("error")
    return 
  }
  fmt.Println(oxygenRating * co2Rating)
}
