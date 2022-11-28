package main

import (
	"AOC/day4/input"
	"fmt"
	"strconv"
	"strings"
  "math"
)

func main() {
	//[testData, realData]
	data, _ := input.Data()
	part1(data)
	part2(data)
}

func part2(data []string) {
	bingoNumbers, boards := parseInput(data)
  winningTurn := 0
  var winningBoard Board

	// for i, board := range boards{
	for boardNum, board := range boards {
		for turn, bingoNumber := range bingoNumbers {
			for y := 0; y < len(board.grid); y++ {
        if board.won {
          continue
        }
				for x := 0; x < len(board.grid[0]); x++ {
					if board.won {
						continue
					}
					// check if number is hit then do the below
					currentNumber := board.grid[y][x]
					if currentNumber == bingoNumber {
						posValue, _ := strconv.Atoi(currentNumber)
						board.baseScore -= posValue

						if board.yHits == nil {
							board.yHits = map[int]int{y: 1}
						} else {
							board.yHits[y]++
						}

						if board.xHits == nil {
							board.xHits = map[int]int{x: 1}
						} else {
							board.xHits[x]++
						}

						if board.yHits[y] == 5 || board.xHits[x] == 5{
              board.won = true
              board.score = board.baseScore * posValue
							board.winningTurn = turn - 1
							if board.winningTurn > winningTurn{
								winningBoard = board
                winningTurn = board.winningTurn
							}
						}
            boards[boardNum] = board
					}
				}
			}
		}
	}
  fmt.Println("Part 2 Final Score:")
	fmt.Println(winningBoard.score)
}

func part1(data []string) {
	//xHits := make(map[int]map[int]int)

	bingoNumbers, boards := parseInput(data)
  winningTurn := math.MaxInt
  var winningBoard Board

	// for i, board := range boards{
	for boardNum, board := range boards {
		for turn, bingoNumber := range bingoNumbers {
			for y := 0; y < len(board.grid); y++ {
        if board.won {
          continue
        }
				for x := 0; x < len(board.grid[0]); x++ {
					if board.won {
						continue
					}
					// check if number is hit then do the below
					currentNumber := board.grid[y][x]
					if currentNumber == bingoNumber {
						posValue, _ := strconv.Atoi(currentNumber)
						board.baseScore -= posValue

						if board.yHits == nil {
							board.yHits = map[int]int{y: 1}
						} else {
							board.yHits[y]++
						}

						if board.xHits == nil {
							board.xHits = map[int]int{x: 1}
						} else {
							board.xHits[x]++
						}

						if board.yHits[y] == 5 || board.xHits[x] == 5{
              board.won = true
              board.score = board.baseScore * posValue
							board.winningTurn = turn - 1
							if board.winningTurn < winningTurn{
								winningBoard = board
                winningTurn = board.winningTurn
							}
						}
            boards[boardNum] = board
					}
				}
			}
		}
	}
  fmt.Println("Part 1 Final Score:")
	fmt.Println(winningBoard.score)
}


func sliceContains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func parseInput(data []string) ([]string, []Board) {
	boards := []Board{}
	var tempBoard [][]string
	var tempBoardScore int
	bingoNumbers := strings.Split(data[0], ",")
	stringBoards := append(data[2:], "")

	for _, line := range stringBoards {
		if line != "" {
			trimed := strings.TrimSpace(line)
			noDoubleSpaces := strings.ReplaceAll(trimed, "  ", ",")
			noSpaces := strings.ReplaceAll(noDoubleSpaces, " ", ",")
			row := strings.Split(noSpaces, ",")
			tempBoard = append(tempBoard, row)

			for _, item := range row {
				num, _ := strconv.Atoi(item)
				tempBoardScore += num
			}
		} else {
			boards = append(boards, Board{
				grid:      tempBoard,
				baseScore: tempBoardScore,
			})
			tempBoard = [][]string{}
			tempBoardScore = 0
		}
	}
	return bingoNumbers, boards
}

type Board struct {
	grid      [][]string
	baseScore int
  score int
	yHits     map[int]int
	xHits     map[int]int
	won bool
	winningTurn int
}

type WinningBoard struct {
	turns int
  score int
  board Board
}
