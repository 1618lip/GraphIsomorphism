package main

import (
	"fmt"
	"strings"
)

// Graph represents an undirected graph using an adjacency matrix.
type Graph struct {
	NumVertices int
	AdjMatrix   [][]int
}

// NewGraph creates a new graph with n vertices.
func NewGraph(n int) *Graph {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return &Graph{NumVertices: n, AdjMatrix: matrix}
}

// AddEdge adds an undirected edge between vertices u and v.
func (g *Graph) AddEdge(u, v int) {
	if u >= 0 && u < g.NumVertices && v >= 0 && v < g.NumVertices {
		g.AdjMatrix[u][v] = 1
		g.AdjMatrix[v][u] = 1
	}
}

// String returns a string representation of the graph.
func (g *Graph) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Graph with %d vertices:\n", g.NumVertices))
	for i, row := range g.AdjMatrix {
		sb.WriteString(fmt.Sprintf("%d: %v\n", i, row))
	}
	return sb.String()
}

// Degree returns the degree of vertex v.
func (g *Graph) Degree(v int) int {
	sum := 0
	for j := 0; j < g.NumVertices; j++ {
		sum += g.AdjMatrix[v][j]
	}
	return sum
}
