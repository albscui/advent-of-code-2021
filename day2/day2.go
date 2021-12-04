package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	amount    int
}

type Submarine struct {
	horizontalPosition int
	depth              int
	aim                int
}

func (s *Submarine) exec(cmd Command) {
	switch cmd.direction {
	case "forward":
		s.horizontalPosition += cmd.amount
	case "up":
		s.depth -= cmd.amount
	case "down":
		s.depth += cmd.amount
	}

}

func (s *Submarine) exec2(cmd Command) {
	switch cmd.direction {
	case "down":
		s.aim += cmd.amount
	case "up":
		s.aim -= cmd.amount
	case "forward":
		s.horizontalPosition += cmd.amount
		s.depth += (s.aim * cmd.amount)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// parse input
	// commands := []Command{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}
	commands := []Command{}
	f, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		direction := words[0]
		amount, err := strconv.Atoi(words[1])
		check(err)
		commands = append(commands, Command{direction, amount})
	}

	// Part One
	submarine := Submarine{0, 0, 0}
	for _, cmd := range commands {
		submarine.exec(cmd)
	}
	fmt.Println("Part One:", submarine.horizontalPosition*submarine.depth)

	// Part Two
	// reset submarine
	submarine.horizontalPosition = 0
	submarine.depth = 0
	submarine.aim = 0
	for _, cmd := range commands {
		submarine.exec2(cmd)
	}
	fmt.Println("Part Two:", submarine.horizontalPosition*submarine.depth)

}
