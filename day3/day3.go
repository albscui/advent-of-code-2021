package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(filename string) (lines []string) {
	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func main() {
	// iterate through the input line by line
	// for each line of the input, get the length of the number
	// create a map called freq
	// where freq[i][0] is the count of all 0's at position i
	//       freq[i][1] is the count of all 1's at position i
	// then to determine the most common bit at any position i
	// we can simply compare the two
	lines := readLines("input")
	freq := make(map[int]*[2]int)
	for _, line := range lines {
		for i, c := range line {
			j, err := strconv.Atoi(string(c))
			check(err)
			if _, prs := freq[i]; !prs {
				freq[i] = &[2]int{0, 0}
			}
			freq[i][j]++
		}
	}

	gammaRateSlice := []string{}
	epsilonRateSlice := []string{}
	for i := 0; i < len(freq); i++ {
		if freq[i][0] < freq[i][1] {
			gammaRateSlice = append(gammaRateSlice, "1")
			epsilonRateSlice = append(epsilonRateSlice, "0")
		} else {
			gammaRateSlice = append(gammaRateSlice, "0")
			epsilonRateSlice = append(epsilonRateSlice, "1")
		}
	}

	// Part One
	gammaRateStr := strings.Join(gammaRateSlice, "")
	epsilonRateStr := strings.Join(epsilonRateSlice, "")
	gammaRate, err := strconv.ParseInt(gammaRateStr, 2, 32)
	check(err)
	epsilonRate, err := strconv.ParseInt(epsilonRateStr, 2, 32)
	check(err)
	fmt.Println("Part One:", gammaRate*epsilonRate)

	// Part Two
	// iterate on lines, continuously update lines with filtered results
	// use i to track which index we are filtering on
	// each iteration: calculate the most frequent value at index i
	// all the while sorting each line into zero_basket, one_basket
	// then just update lines to equal to the basket that represents the most frequent value
	oxygenGeneratorRatings := lines
	N := len(gammaRateStr)
	for i := 0; i < N; i++ {
		if len(oxygenGeneratorRatings) == 1 {
			break
		}
		zeros, ones := 0, 0
		zerosBucket, onesBucket := []string{}, []string{}
		for _, line := range oxygenGeneratorRatings {
			if line[i] == '0' {
				zeros++
				zerosBucket = append(zerosBucket, line)
			} else {
				ones++
				onesBucket = append(onesBucket, line)
			}
		}
		if zeros > ones {
			oxygenGeneratorRatings = zerosBucket
		} else {
			oxygenGeneratorRatings = onesBucket
		}
	}

	co2ScrubberRatings := lines
	for i := 0; i < N; i++ {
		if len(co2ScrubberRatings) == 1 {
			break
		}
		zeros, ones := 0, 0
		zerosBucket, onesBucket := []string{}, []string{}
		for _, line := range co2ScrubberRatings {
			if line[i] == '0' {
				zeros++
				zerosBucket = append(zerosBucket, line)
			} else {
				ones++
				onesBucket = append(onesBucket, line)
			}
		}
		if zeros <= ones {
			co2ScrubberRatings = zerosBucket
		} else {
			co2ScrubberRatings = onesBucket
		}
	}

	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatings[0], 2, 32)
	check(err)
	co2ScrubberRating, err := strconv.ParseInt(co2ScrubberRatings[0], 2, 32)
	check(err)
	fmt.Println("Part Two:", oxygenGeneratorRating*co2ScrubberRating)
}
