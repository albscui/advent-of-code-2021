package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func threeMeasurementSlidingWindowSums(a []int) (sums []int) {
	// assume at least 3 elements
	sums = append(sums, a[0]+a[1]+a[2])
	for i := 3; i < len(a); i++ {
		sums = append(sums, sums[len(sums)-1]+a[i]-a[i-3])
	}
	return
}

func countNumberOfTimesDepthIncreases(depths []int) (count int) {
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return
}

func main() {
	fp, err := os.Open("input")
	check(err)
	defer fp.Close()

	input := []int{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		check(err)
		input = append(input, n)
	}

	// example := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	fmt.Println("Part One Answer:", countNumberOfTimesDepthIncreases(input))
	fmt.Println("Part Two Answer:", countNumberOfTimesDepthIncreases(threeMeasurementSlidingWindowSums(input)))
}
