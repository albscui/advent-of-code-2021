package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	unmarkedNumbers map[int]*[2]int
	// markedNumbers   map[int]*[2]int
	rowsMatched   [5]int // rowsMatched[i] represents how many matches in ith row
	colsMatched   [5]int // columnsMatched[i] represents how many matches in ith column
	winningNumber int
}

func (b Board) sumUnMarkedNumbers() (s int) {
	for k := range b.unmarkedNumbers {
		s += k
	}
	return
}

type Bingo struct {
	seq               []int
	boards            []Board
	firstWinningBoard *Board
	lastWinningBoard  *Board
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(path string) (bingo Bingo) {
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)

	// first line is the sequence of numbers drawn
	scanner.Scan()
	seqStrs := strings.Split(scanner.Text(), ",")
	for _, n := range seqStrs {
		x, err := strconv.Atoi(n)
		check(err)
		bingo.seq = append(bingo.seq, x)
	}

	// Deserialize the boards
	// the first scanner.Scan() should be an empty line
	for scanner.Scan() {
		board := Board{unmarkedNumbers: make(map[int]*[2]int)}
		for r := 0; r < 5; r++ {
			scanner.Scan()
			rowValues := strings.Fields(scanner.Text())
			for c, v := range rowValues {
				x, err := strconv.Atoi(v)
				check(err)
				board.unmarkedNumbers[x] = &[2]int{r, c}
			}
		}
		bingo.boards = append(bingo.boards, board)
	}

	return
}

func main() {
	bingo := parseInput("input")

	for _, n := range bingo.seq {
		for i := range bingo.boards {
			board := &bingo.boards[i]
			if board.winningNumber != 0 {
				continue
			}
			if coord, prs := board.unmarkedNumbers[n]; prs {
				r := coord[0]
				c := coord[1]
				delete(board.unmarkedNumbers, n)
				board.rowsMatched[r]++
				board.colsMatched[c]++
				if board.rowsMatched[r] == 5 || board.colsMatched[c] == 5 {
					board.winningNumber = n
					if bingo.firstWinningBoard == nil {
						bingo.firstWinningBoard = board
					}
					bingo.lastWinningBoard = board
				}
			}
		}
	}
	fmt.Println("Bingo!")
	fmt.Println(bingo.firstWinningBoard.winningNumber * bingo.firstWinningBoard.sumUnMarkedNumbers())
	fmt.Println(bingo.lastWinningBoard.winningNumber * bingo.lastWinningBoard.sumUnMarkedNumbers())
}
