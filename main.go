package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readGraph reads a graph from STDIN
// It expects:
//   - First non-empty line: "n m" (number of vertices and edges)
//   - Next m lines: "u v" (each edge; vertices may be 0-indexed or 1-indexed)
//
// It auto-detects indexing by examining the maximum vertex label.
func readGraph(scanner *bufio.Scanner) *Graph {
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line != "" {
			break
		}
	}
	parts := strings.Fields(line)
	if len(parts) < 2 {
		fmt.Println("Error: expected first line to contain 'n m'")
		os.Exit(1)
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Error: invalid number of vertices")
		os.Exit(1)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error: invalid number of edges")
		os.Exit(1)
	}

	var edges [][2]int
	maxLabel := 0
	for i := 0; i < m; i++ {
		if !scanner.Scan() {
			fmt.Printf("Error: expected %d edges, got %d\n", m, i)
			os.Exit(1)
		}
		line = strings.TrimSpace(scanner.Text())
		if line == "" {
			i--
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Printf("Error: invalid edge format on line %d\n", i+2)
			os.Exit(1)
		}
		u, err1 := strconv.Atoi(parts[0])
		v, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error: invalid edge values on line %d\n", i+2)
			os.Exit(1)
		}
		edges = append(edges, [2]int{u, v})
		if u > maxLabel {
			maxLabel = u
		}
		if v > maxLabel {
			maxLabel = v
		}
	}

	adjust := false
	if maxLabel == n {
		adjust = true // 1-indexed
	} else if maxLabel == n-1 {
		adjust = false // 0-indexed
	} else {
		fmt.Printf("Warning: max vertex label %d doesn't equal n-1 (%d); assuming 1-indexed input.\n", maxLabel, n-1)
		adjust = true
	}

	g := NewGraph(n)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		if adjust {
			u--
			v--
		}
		g.AddEdge(u, v)
	}
	return g
}

func main() {
	fmt.Println("Graph Isomorphism Checker with Automorphism-based Symmetry Breaking (Go)")
	fmt.Println("Input format for each graph:")
	fmt.Println("  First line: n m (number of vertices and edges)")
	fmt.Println("  Next m lines: u v (each edge; vertices can be 0-indexed or 1-indexed)")
	fmt.Println("Enter Graph 1, then Graph 2 (each in the same format).")
	fmt.Println("---------------------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Reading Graph 1:")
	g1 := readGraph(scanner)
	fmt.Println("\nGraph 1:")
	fmt.Println(g1)

	fmt.Println("Reading Graph 2:")
	g2 := readGraph(scanner)
	fmt.Println("\nGraph 2:")
	fmt.Println(g2)

	if areIsomorphicSym(g1, g2) {
		fmt.Println("\nThe graphs are isomorphic (using symmetry breaking).")
		if mapping, ok := findIsomorphismSym(g1, g2); ok {
			fmt.Println("A valid mapping from Graph 1 to Graph 2:")
			for i, v := range mapping {
				fmt.Printf("  Vertex %d in Graph 1 -> Vertex %d in Graph 2\n", i, v)
			}
		}
	} else {
		fmt.Println("\nThe graphs are not isomorphic.")
	}
}
