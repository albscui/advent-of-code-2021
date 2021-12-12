package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Octopus struct {
	energy  int
	flashed bool
}

type OctopusMatrix [][]Octopus

func (m OctopusMatrix) Size() int {
	return len(m) * len(m[0])
}

func (m OctopusMatrix) neighbours(r, c int) (ret [][2]int) {
	ds := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range ds {
		_r := r + d[0]
		_c := c + d[1]
		if 0 <= _r && _r < len(m) && 0 <= _c && _c < len(m[0]) {
			ret = append(ret, [2]int{_r, _c})
		}
	}
	return
}

func (m OctopusMatrix) flash(r, c int) int {
	count := 1
	octopus := &m[r][c]
	octopus.flashed = true
	for _, nbCoord := range m.neighbours(r, c) {
		_r, _c := nbCoord[0], nbCoord[1]
		neighbour := &m[_r][_c]
		if !neighbour.flashed {
			neighbour.energy++
			if neighbour.energy > 9 {
				count += m.flash(_r, _c)
			}
		}
	}
	octopus.energy = 0
	return count
}

// Run simulates a single step across all octopi, and returns the number of flashes
func (m OctopusMatrix) Run() (count int) {
	// First, the energy level of each octopus increases by 1
	for r := range m {
		for c := range m[0] {
			m[r][c].energy++
		}
	}
	// Then, any octopus with an energy level greater than 9 flashes.
	for r := range m {
		for c := range m[0] {
			if m[r][c].energy > 9 {
				count += m.flash(r, c)
			}
		}
	}
	// reset flashed state
	for r := range m {
		for c := range m[0] {
			m[r][c].flashed = false
		}
	}
	return
}

func parseInput(path string) (m OctopusMatrix) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "")
		row := []Octopus{}
		for _, s := range splitted {
			energy, _ := strconv.Atoi(s)
			row = append(row, Octopus{energy: energy})
		}
		m = append(m, row)
	}
	return
}

func partOne(matrix OctopusMatrix) (flashes int) {
	for step := 0; step < 100; step++ {
		flashes += matrix.Run()
	}
	return
}

func partTwo(matrix OctopusMatrix) (step int) {
	for matrix.Run() != matrix.Size() {
		step++
	}
	step++
	return
}

func main() {
	path := os.Args[1]
	matrix := parseInput(path)
	fmt.Println(partOne(matrix))

	matrix = parseInput(path)
	fmt.Println(partTwo(matrix))
}
