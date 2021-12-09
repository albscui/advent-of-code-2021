package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type costFn func(int, int) int

func parseInput(path string) (result []int) {
	dat, _ := os.ReadFile(path)
	datSplitted := strings.Split(string(dat), ",")
	for _, n := range datSplitted {
		x, _ := strconv.Atoi(n)
		result = append(result, x)
	}
	return
}

func avg(a []int) float64 {
	return float64(sum(a)) / float64(len(a))
}

func abs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func partialSum(n int) int {
	return n * (n + 1) / 2
}

func partTwoCostFn(a, b int) int {
	return partialSum(abs(a, b))
}

func sum(a []int) (s int) {
	for i := range a {
		s += a[i]
	}
	return
}

func totalCost(a []int, pos int, cost costFn) (c int) {
	for i := range a {
		c += cost(a[i], pos)
	}
	return
}

func partOne(a []int) int {
	sort.Ints(a)
	return totalCost(a, a[len(a)/2], abs)
}

func partTwo(a []int) int {
	mean := avg(a)
	floor := int(math.Floor(mean))
	ceil := int(math.Ceil(mean))
	floorCost := totalCost(a, floor, partTwoCostFn)
	ceilCost := totalCost(a, ceil, partTwoCostFn)
	return int(math.Min(float64(floorCost), float64(ceilCost)))
}

func main() {
	input := parseInput("input")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
