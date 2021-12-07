package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HORIZONTAL IntervalType = iota
	VERTICAL
	DIAGONAL
)

type IntervalType int

type Point struct {
	x, y int
}

type Interval struct {
	start, end   Point
	intervalType IntervalType
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(path string) (intervals []Interval) {
	f, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), " -> ")
		p1 := strings.Split(lineSplitted[0], ",")
		p2 := strings.Split(lineSplitted[1], ",")
		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])

		var intervalType IntervalType
		if x1 == x2 {
			intervalType = VERTICAL
		} else if y1 == y2 {
			intervalType = HORIZONTAL
		} else {
			intervalType = DIAGONAL
		}
		intervals = append(intervals, Interval{Point{x1, y1}, Point{x2, y2}, intervalType})
	}
	return
}

func partOnePointsBetween(interval Interval) (points []Point) {
	// only care about horizontal and vertical intervals
	if interval.start.x == interval.end.x {
		x, y1, y2 := interval.start.x, interval.start.y, interval.end.y
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			points = append(points, Point{x, y})
		}
	} else if interval.start.y == interval.end.y {
		x1, x2, y := interval.start.x, interval.end.x, interval.start.y
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			points = append(points, Point{x, y})
		}
	}
	return
}

func partTwoPointsBetween(interval Interval) (points []Point) {
	// include diagonals as well
	if interval.intervalType == HORIZONTAL || interval.intervalType == VERTICAL {
		return partOnePointsBetween(interval)
	}

	// align the points such that we always go from left to right
	if interval.start.x > interval.end.x {
		interval.start, interval.end = interval.end, interval.start
	}

	x, y := interval.start.x, interval.start.y
	if interval.start.y < interval.end.y {
		// low to high
		for y <= interval.end.y {
			points = append(points, Point{x, y})
			x++
			y++
		}
	} else {
		// high to low
		for y >= interval.end.y {
			points = append(points, Point{x, y})
			x++
			y--
		}
	}
	return
}

/*
For each interval, get all the points that overlap

How to find all points that overlap?
	Use a map where the key is the Point.key

Increment the overlap count for each point
At the end, just go through all the points, and see which ones have overlaps >= 2
*/
func main() {
	intervals := parseInput("input")

	// Part One
	pointsMap := make(map[Point]int)
	for _, interval := range intervals {
		for _, point := range partOnePointsBetween(interval) {
			pointsMap[point]++
		}
	}

	ans := 0
	for _, count := range pointsMap {
		if count >= 2 {
			ans++
		}
	}
	fmt.Println("Part One:", ans)

	// Part Two
	pointsMap = make(map[Point]int)
	for _, interval := range intervals {
		for _, point := range partTwoPointsBetween(interval) {
			pointsMap[point]++
		}
	}

	ans = 0
	for _, count := range pointsMap {
		if count >= 2 {
			ans++
		}
	}
	fmt.Println("Part Two:", ans)
}
