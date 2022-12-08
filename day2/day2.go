package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//[testData, realData]
	_, data := input()
	part1(data)
	part2(data)
}

//A, X = Rock, score = 1
//B, Y = Paper, score = 2
//C, Z = Scissors, score = 3

//A = X
//A < Y
//A > Z

//B = Y
//B < Z
//B > X

//C = Z
//C < X
//C > Y

//round score
//loss = 0
//draw = 3
//win = 6

func part1(data []string) {
	fmt.Println("Part 1")
	totalScore := 0
	outcome := make(map[string]map[string]int)
	outcome["A"] = make(map[string]int)
	outcome["A"]["X"] = 3 + 1
	outcome["A"]["Y"] = 6 + 2
	outcome["A"]["Z"] = 0 + 3

	outcome["B"] = make(map[string]int)
	outcome["B"]["X"] = 0 + 1
	outcome["B"]["Y"] = 3 + 2
	outcome["B"]["Z"] = 6 + 3

	outcome["C"] = make(map[string]int)
	outcome["C"]["X"] = 6 + 1
	outcome["C"]["Y"] = 0 + 2
	outcome["C"]["Z"] = 3 + 3

	for _, match := range data {
		moves := strings.Split(match, " ")

		opponent := moves[0]
		me := moves[1]
		roundScore := outcome[opponent][me]
		fmt.Println(roundScore)
		totalScore += roundScore
	}

	fmt.Println(totalScore)
}

func part2(data []string) {
	fmt.Println("Part 2")
	totalScore := 0
	for _, match := range data {
		moves := strings.Split(match, " ")
		opponent := moves[0]
		outcome := moves[1]
		//loss
		if outcome == "X" {
			if opponent == "A" {
				totalScore += 3
			}
			if opponent == "B" {
				totalScore += 1
			}
			if opponent == "C" {
				totalScore += 2
			}
		}
		//tie
		if outcome == "Y" {
			if opponent == "A" {
				totalScore += 3 + 1
			}
			if opponent == "B" {
				totalScore += 3 + 2
			}
			if opponent == "C" {
				totalScore += 3 + 3
			}
		}
		//win
		if outcome == "Z" {
			if opponent == "A" {
				totalScore += 6 + 2
			}
			if opponent == "B" {
				totalScore += 6 + 3
			}
			if opponent == "C" {
				totalScore += 6 + 1
			}
		}
	}
	fmt.Println(totalScore)
}

func input() ([]string, []string) {
	test := openCSV("./input/testInput.csv")
	data := openCSV("./input/input.csv")
	return test, data
}

func openCSV(fileName string) []string {
	data := []string{}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, []string{rec[0]}...)
	}

	return data
}

/**
Advent of Code[About][Events][Shop][Settings][Log Out]graphicnapkin 2*
  {:year 2022}[Calendar][AoC++][Sponsors][Leaderboard][Stats]
Our sponsors help make Advent of Code possible:
ZBRA - Um time que se preocupa não só com a qualidade do seu código, mas com sua qualidade de vida. Conheça a ZBRA. Antes dos algoritmos, PESSOAS!
--- Day 2: Rock Paper Scissors ---
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?

To begin, get your puzzle input.

Answer:


You can also [Share] this puzzle.
**/
