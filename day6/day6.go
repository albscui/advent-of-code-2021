package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type FishManager struct {
	mu        sync.Mutex
	wg        sync.WaitGroup
	fishCount int
	maxDay    int
}

// Global FishManager
var fm FishManager

type Fish struct {
	name  string
	timer int
	day   int
}

// StateFn represents the state of the fish
type StateFn func(*Fish) StateFn

func swimming(fish *Fish) StateFn {
	// fmt.Println("[swimming]", fish)
	fish.timer--
	fish.day++

	if fish.timer == 0 {
		return reproducing
	}
	return swimming
}

func reproducing(fish *Fish) StateFn {
	// fmt.Println("[reproducing]", fish)
	fish.timer = 6
	fish.day++

	// create a new Fish
	child := Fish{fish.name + "_child", 8, fish.day}
	fm.wg.Add(1)
	go func() {
		defer fm.wg.Done()
		child.Run()
	}()

	fm.mu.Lock()
	fm.fishCount++
	fm.mu.Unlock()

	return swimming
}

// run until the each fish reaches the max day allowed
func (fish *Fish) Run() {
	for state := swimming; fish.day < fm.maxDay; {
		state = state(fish)
	}
	// fmt.Println("[Done]", fish)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fm.maxDay = 80
	dat, err := os.ReadFile("input")
	check(err)
	initialTimers := strings.Split(string(dat), ",")

	for i, t := range initialTimers {
		timer, _ := strconv.Atoi(t)
		name := "fish" + strconv.Itoa(i)
		fish := Fish{name, timer, 0}

		fm.wg.Add(1)
		go func() {
			defer fm.wg.Done()
			fish.Run()
		}()

		fm.mu.Lock()
		fm.fishCount++
		fm.mu.Unlock()
	}

	fm.wg.Wait()

	fmt.Println(fm.fishCount)
}
