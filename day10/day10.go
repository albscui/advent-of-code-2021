package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func matchingBracket(b rune) rune {
	switch b {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	}
	panic(fmt.Sprintf("%v not supported \n", b))
}

func illegalCharScore(b rune) int {
	switch b {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	panic(fmt.Sprintf("%v not supported\n", b))
}

func checkLegal(chunk string) (rune, string) {
	stack := []rune{}
	for _, c := range chunk {
		switch c {
		case '(', '[', '{', '<':
			stack = append(stack, c)
		case ')', ']', '}', '>':
			if matchingBracket(c) != stack[len(stack)-1] {
				// corrupted
				return c, string(stack)
			}
			stack = stack[:len(stack)-1]
		}
	}
	return rune(0), string(stack)
}

func completionString(open string) (close string) {
	completion := []rune{}
	for i := range open {
		completion = append(completion, matchingBracket(rune(open[len(open)-1-i])))
	}
	return string(completion)
}

func completionScore(completion string) (score int) {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	for _, c := range completion {
		score = score*5 + points[c]
	}
	return
}

func partOne(chunks []string) (ans int) {
	for _, chunk := range chunks {
		if illegalChar, _ := checkLegal(chunk); illegalChar != rune(0) {
			ans += illegalCharScore(illegalChar)
		}
	}
	return
}

func partTwo(chunks []string) (ans int) {
	scores := []int{}
	for _, chunk := range chunks {
		if illegalChar, stack := checkLegal(chunk); illegalChar == rune(0) && len(stack) > 0 {
			// not illegal but incomplete
			scores = append(scores, completionScore(completionString(stack)))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func parseInput(path string) (ret []string) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return
}

func main() {
	input := parseInput(os.Args[1])
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
