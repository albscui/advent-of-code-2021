package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input []Puzzle

type Puzzle struct {
	observations []string
	outputs      []string
	cipher       map[string]string // maps single encrypted to single decrypted letter
	numberCodes  [10]string        // maps each unique number to its code
}

func (p *Puzzle) fillNumberCodes() {
	for _, obs := range p.observations {
		switch len(obs) {
		case 2:
			p.numberCodes[1] = obs
		case 4:
			p.numberCodes[4] = obs
		case 3:
			p.numberCodes[7] = obs
		case 7:
			p.numberCodes[8] = obs
		}
	}
}

// return all the strings that contains a but not b
func (p *Puzzle) exclusiveContains(a, b string) (ans []string) {
	a, b = sortString(a), sortString(b)
	for _, obs := range p.observations {
		if strings.Contains(obs, a) && !strings.Contains(obs, b) {
			ans = append(ans, obs)
		}
	}
	return
}

func (p *Puzzle) knownLetters() string {
	keys := make([]string, len(p.cipher))
	i := 0
	for k := range p.cipher {
		keys[i] = k
		i++
	}
	return strings.Join(keys, "")
}

/*
a:
	Get codes for #1 and #7, the diff maps to "a"
c/f:
	Get the code for #1 -> (x, y)
	Get the numbers that only has x but not y
	if only one number exists, that must be 2, which means
		x -> "c"
		y -> "f"
	otherwise swap
b/d:
	Get the code for #4, diff with #1 -> (x, y)
	Find the number that only has x or y
	if found:
		x -> "b"
		y -> "d"
	otherwise swap
e/g:
	Get code for #8, diff with all existing known letters -> (x, y)
	Find the numbers that only has x or y
	if there are 3 numbers
		x -> "e"
		y -> "g"
	otherwise swap
*/
func (p *Puzzle) Solve() int {
	// setup
	p.fillNumberCodes()
	// solve for "a"
	one := p.numberCodes[1]
	seven := p.numberCodes[7]
	d := diff(one, seven)
	p.cipher[d] = "a"

	// solve for "c" and "f"
	x, y := string(one[0]), string(one[1])
	if len(p.exclusiveContains(x, y)) == 1 {
		p.cipher[x] = "c"
		p.cipher[y] = "f"
	} else {
		p.cipher[y] = "c"
		p.cipher[x] = "f"
	}

	// solve for "b" and "d"
	four := p.numberCodes[4]
	d = diff(one, four)
	x, y = string(d[0]), string(d[1])
	if len(p.exclusiveContains(x, y)) == 1 {
		p.cipher[x] = "b"
		p.cipher[y] = "d"
	} else {
		p.cipher[y] = "b"
		p.cipher[x] = "d"
	}

	// solve for "e" and "g"
	eight := p.numberCodes[8]
	d = diff(eight, p.knownLetters())
	x, y = string(d[0]), string(d[1])
	if len(p.exclusiveContains(x, y)) == 3 {
		p.cipher[x] = "g"
		p.cipher[y] = "e"
	} else {
		p.cipher[y] = "g"
		p.cipher[x] = "e"
	}

	outputs := make([]string, len(p.outputs))
	for i, o := range p.outputs {
		outputs[i] = strconv.Itoa(p.decode(o))
	}
	ans, _ := strconv.Atoi(strings.Join(outputs, ""))
	return ans
}

func (p *Puzzle) decode(s string) int {
	a := make([]string, len(s))
	for i, c := range s {
		_c := p.cipher[string(c)]
		a[i] = _c
	}
	sort.Strings(a)
	return toNumber(strings.Join(a, ""))
}

func diff(s1, s2 string) string {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	d := []string{}
	for _, c := range s2 {
		_c := string(c)
		if !strings.Contains(s1, string(_c)) {
			d = append(d, _c)
		}
	}
	return strings.Join(d, "")
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func toNumber(k string) int {
	legend := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}
	if v, prs := legend[sortString(k)]; prs {
		return v
	}
	panic(fmt.Sprintf("Could not convert %s to a number\n", k))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(path string) (out Input) {
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.Split(line, " | ")
		observations := strings.Split(lineSplitted[0], " ")
		output := strings.Split(lineSplitted[1], " ")
		puzzle := Puzzle{observations, output, make(map[string]string), [10]string{}}
		out = append(out, puzzle)
	}
	return
}

func partOne(input Input) (ans int) {
	for _, puzzle := range input {
		for _, out := range puzzle.outputs {
			switch len(out) {
			case 2, 3, 4, 7:
				ans += 1
			}
		}
	}
	return
}

func partTwo(input Input) (ans int) {
	for i := range input {
		ans += input[i].Solve()
	}
	return
}

func main() {
	input := parseInput(os.Args[1])
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
