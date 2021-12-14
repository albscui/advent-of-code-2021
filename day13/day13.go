package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UP = iota
	LEFT
)

type Coord struct {
	row, col int
}

// Paper represents the 2D paper of dots, as well as the fold boundaries
type Paper struct {
	rows, cols int
	dots       map[Coord]bool
}

type FoldInstruction struct {
	direction int
	value     int
}

func (p Paper) CountDots() (count int) {
	for _, val := range p.dots {
		if val {
			count++
		}
	}
	return
}

// String implements the Stringer interface so that we can print out the message in the paper
func (p Paper) String() string {
	matrix := make([][]string, p.rows)
	for i := range matrix {
		matrix[i] = make([]string, p.cols)
	}
	for r := 0; r < p.rows; r++ {
		for c := 0; c < p.cols; c++ {
			coord := Coord{r, c}
			if p.dots[coord] {
				matrix[r][c] = "#"
			} else {
				matrix[r][c] = "."
			}
		}
	}
	rows := []string{}
	for _, row := range matrix {
		rows = append(rows, strings.Join(row, ""))
	}
	return strings.Join(rows, "\n")
}

func (paper *Paper) Fold(fold FoldInstruction) {
	if fold.direction == UP {
		for dot := range paper.dots {
			if dot.row > fold.value {
				_r := dot.row - 2*(dot.row-fold.value)
				paper.dots[Coord{_r, dot.col}] = true
				paper.dots[dot] = false
			}
		}
		// update the horizontal bound
		paper.rows = fold.value
	} else if fold.direction == LEFT {
		for dot := range paper.dots {
			if dot.col > fold.value {
				_c := dot.col - 2*(dot.col-fold.value)
				paper.dots[Coord{dot.row, _c}] = true
				paper.dots[dot] = false
			}
		}
		// update the vertical boundary
		paper.cols = fold.value
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(path string) (paper Paper, instructions []FoldInstruction) {
	paper.dots = make(map[Coord]bool)
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// stop scanner for coords
			// start scanning for instructions
			break
		}
		coord := strings.Split(scanner.Text(), ",")
		c, _ := strconv.Atoi(coord[0])
		r, _ := strconv.Atoi(coord[1])
		dot := Coord{r, c}
		paper.dots[dot] = true
		// keep expanding the boundaries
		if c+1 > paper.cols {
			paper.cols = c + 1
		}
		if r+1 > paper.rows {
			paper.rows = r + 1
		}
	}

	for scanner.Scan() {
		s := strings.Split(strings.TrimPrefix(scanner.Text(), "fold along "), "=")
		direction := UP
		if s[0] == "y" {
			direction = UP
		} else {
			direction = LEFT
		}
		value, _ := strconv.Atoi(s[1])
		instructions = append(instructions, FoldInstruction{direction, value})
	}
	return
}

func main() {
	paper, folds := parseInput(os.Args[1])

	// Part One
	fold := folds[0]
	paper.Fold(fold)
	fmt.Println(paper.CountDots())

	// Part Two
	folds = folds[1:]
	for _, fold := range folds {
		paper.Fold(fold)
	}
	fmt.Println(paper)
}
