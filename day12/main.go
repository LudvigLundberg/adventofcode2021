package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/LudvigLundberg/adventofcode2021/parseinput"
)

func main() {
	input, err := parseinput.ParseFile("input")

	if err != nil {
		panic(err)
	}

	partOne(input)
	partTwo(input)
}

func partOne(input []string) {
	fmt.Println("-------- Part One -------")

	edges := createEdges(input)
	graph := createGraph(edges)
	paths := numberOfPaths(graph, "start", "end", map[string]bool{}, true)

	fmt.Printf("number of paths %v\n", paths)
}

func partTwo(input []string) {
	fmt.Println("-------- Part Two -------")

	edges := createEdges(input)
	graph := createGraph(edges)
	paths := numberOfPaths(graph, "start", "end", map[string]bool{}, false)

	fmt.Printf("number of paths %v\n", paths)
}

type Graph = map[string][]string
type Edge struct {
	start, end string
}

func createEdges(input []string) []Edge {
	edges := make([]Edge, len(input))

	for i, line := range input {
		edge := strings.Split(line, "-")
		edges[i] = Edge{start: edge[0], end: edge[1]}
	}
	return edges
}

func createGraph(edges []Edge) Graph {
	graph := make(map[string][]string)
	for _, edge := range edges {
		vertices, ok := graph[edge.start]
		if !ok {
			graph[edge.start] = []string{edge.end}
		} else {
			graph[edge.start] = append(vertices, edge.end)
		}

		vertices, ok = graph[edge.end]
		if !ok {
			graph[edge.end] = []string{edge.start}
		} else {
			graph[edge.end] = append(vertices, edge.start)
		}
	}

	return graph
}

func numberOfPaths(g Graph, start string, end string, visited map[string]bool, smallCaveVisited bool) int {
	if start == end {
		return 1
	}
	paths := 0

	if !visited[start] || !smallCaveVisited {
		if visited[start] && start != "start" {
			smallCaveVisited = true
		}
		for _, node := range g[start] {
			if node != "start" {
				nextVisited := make(map[string]bool)
				for k, v := range visited {
					nextVisited[k] = v
				}
				if !unicode.IsUpper(rune(start[0])) {
					nextVisited[start] = true
				}
				paths += numberOfPaths(g, node, end, nextVisited, smallCaveVisited)
			}
		}
	}

	return paths
}
