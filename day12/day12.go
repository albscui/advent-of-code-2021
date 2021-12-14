package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	START = "start"
	END   = "end"
)

func isBigCave(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

type Graph map[string][]string

type QueueElement struct {
	node    string
	path    []string
	visited map[string]int
}

func (g Graph) Paths(smallCaveLimit int) (paths [][]string) {
	q := []QueueElement{{START, []string{START}, map[string]int{START: 1}}}
	for len(q) > 0 {
		elem := q[0]
		q = q[1:]
		for _, nb := range g[elem.node] {
			// skip small cave if we've reached the limit
			if nb == START ||
				(!isBigCave(nb) &&
					(elem.visited[nb] == smallCaveLimit ||
						(elem.visited["small_cave_limit_reached"] == 1 && elem.visited[nb] == 1))) {
				continue
			} else if nb == END {
				// found it
				newPath := make([]string, len(elem.path))
				copy(newPath, elem.path)
				newPath = append(newPath, END)
				paths = append(paths, newPath)
			} else if isBigCave(nb) || elem.visited[nb] < smallCaveLimit {
				// make a copy of paths and add the new nb
				newPath := make([]string, len(elem.path))
				copy(newPath, elem.path)
				newPath = append(newPath, nb)
				// make a copy of visited
				newVisited := make(map[string]int)
				for key, value := range elem.visited {
					newVisited[key] = value
				}
				newVisited[nb]++
				if !isBigCave(nb) && newVisited[nb] == smallCaveLimit {
					newVisited["small_cave_limit_reached"] = 1
				}
				q = append(q, QueueElement{nb, newPath, newVisited})
			}
		}
	}
	return paths
}

func parseInput(path string) Graph {
	graph := make(map[string][]string)
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return graph
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := parseInput(os.Args[1])
	fmt.Println(len(input.Paths(1)))
	fmt.Println(len(input.Paths(2)))
}
