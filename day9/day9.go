package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HeightMap [][]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(path string) (hm HeightMap) {
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "")
		row := make([]int, len(numbers))
		for i, n := range numbers {
			x, _ := strconv.Atoi(n)
			row[i] = x
		}
		hm = append(hm, row)
	}
	return
}

func neighbours(hm HeightMap, coord [2]int) (ret [][2]int) {
	r, c := coord[0], coord[1]
	rows, cols := len(hm), len(hm[0])
	nbs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, nb := range nbs {
		_r, _c := r+nb[0], c+nb[1]
		if 0 <= _r && _r < rows && 0 <= _c && _c < cols {
			ret = append(ret, [2]int{_r, _c})
		}
	}
	return
}

func lowpoints(hm HeightMap) (ret [][2]int) {
	for r := range hm {
		for c := range hm[0] {
			lowPoint := true
			for _, nb := range neighbours(hm, [2]int{r, c}) {
				if hm[r][c] >= hm[nb[0]][nb[1]] {
					// not a low point
					lowPoint = false
					break
				}
			}
			if lowPoint {
				ret = append(ret, [2]int{r, c})
			}
		}
	}
	return
}

func bfs(hm HeightMap, start [2]int) int {
	visited := make(map[[2]int]bool)
	visited[start] = true
	q := [][2]int{start}
	for len(q) > 0 {
		// TODO why doesn't coord, q := q[0], q[1:] work?
		coord := q[0]
		q = q[1:]
		for _, nb := range neighbours(hm, coord) {
			_r, _c := nb[0], nb[1]
			if hm[_r][_c] != 9 && !visited[nb] {
				visited[nb] = true
				q = append(q, nb)
			}
		}
	}
	return len(visited)
}

func partOne(hm HeightMap) (ret int) {
	for _, lp := range lowpoints(hm) {
		r, c := lp[0], lp[1]
		ret += hm[r][c] + 1
	}
	return
}

func partTwo(hm HeightMap) (ret int) {
	x, y, z := 0, 0, 0
	for _, lp := range lowpoints(hm) {
		size := bfs(hm, lp)
		if size > x {
			x, y, z = size, x, y
		} else if size > y {
			y, z = size, y
		} else if size > z {
			z = size
		}
	}
	return x * y * z
}

func main() {
	path := os.Args[1]
	input := parseInput(path)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
