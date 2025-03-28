package main

// isMappingValid checks whether a given mapping is a valid isomorphism
// between g1 and g2. mapping[i] = j means vertex i in g1 is mapped to vertex j in g2.
func isMappingValid(g1, g2 *Graph, mapping []int) bool {
	n := g1.NumVertices
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g1.AdjMatrix[i][j] != g2.AdjMatrix[mapping[i]][mapping[j]] {
				return false
			}
		}
	}
	return true
}

// backtrack is a simple backtracking algorithm for graph isomorphism.
func backtrack(g1, g2 *Graph, mapping []int, used []bool, index int) bool {
	n := g1.NumVertices
	if index == n {
		return isMappingValid(g1, g2, mapping)
	}
	for candidate := 0; candidate < n; candidate++ {
		if !used[candidate] {
			mapping[index] = candidate
			used[candidate] = true
			if backtrack(g1, g2, mapping, used, index+1) {
				return true
			}
			used[candidate] = false
		}
	}
	return false
}

// areIsomorphic returns true if g1 and g2 are isomorphic (using plain backtracking).
func areIsomorphic(g1, g2 *Graph) bool {
	if g1.NumVertices != g2.NumVertices {
		return false
	}
	n := g1.NumVertices
	// Basic degree check.
	for i := 0; i < n; i++ {
		deg1, deg2 := 0, 0
		for j := 0; j < n; j++ {
			deg1 += g1.AdjMatrix[i][j]
			deg2 += g2.AdjMatrix[i][j]
		}
		if deg1 != deg2 {
			return false
		}
	}
	mapping := make([]int, n)
	used := make([]bool, n)
	return backtrack(g1, g2, mapping, used, 0)
}

// findIsomorphism returns a valid mapping if the graphs are isomorphic.
func findIsomorphism(g1, g2 *Graph) ([]int, bool) {
	n := g1.NumVertices
	mapping := make([]int, n)
	used := make([]bool, n)
	if backtrack(g1, g2, mapping, used, 0) {
		result := make([]int, n)
		copy(result, mapping)
		return result, true
	}
	return nil, false
}
