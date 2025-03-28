package main

// backtrackSym is similar to the plain backtracking algorithm but uses symmetry breaking.
// If a vertex has a forced mapping (from unique degree/orbit constraints), that mapping is used.
func backtrackSym(g1, g2 *Graph, mapping []int, used []bool, index int, constraints map[int]int) bool {
	n := g1.NumVertices
	if index == n {
		return isMappingValid(g1, g2, mapping)
	}
	// If there's a constraint for this vertex, enforce it.
	if val, ok := constraints[index]; ok {
		if used[val] {
			return false
		}
		mapping[index] = val
		used[val] = true
		if backtrackSym(g1, g2, mapping, used, index+1, constraints) {
			return true
		}
		used[val] = false
		return false
	}
	// Otherwise, try all candidates.
	for candidate := 0; candidate < n; candidate++ {
		if !used[candidate] {
			mapping[index] = candidate
			used[candidate] = true
			if backtrackSym(g1, g2, mapping, used, index+1, constraints) {
				return true
			}
			used[candidate] = false
		}
	}
	return false
}

// uniqueConstraints computes a map of forced vertex mappings based on unique degree values.
// For each degree that appears uniquely in both graphs, we fix the mapping.
func uniqueConstraints(g1, g2 *Graph) map[int]int {
	constraints := make(map[int]int)
	n := g1.NumVertices
	// Compute degree frequencies and store the unique vertex for that degree.
	degreeFreq1 := make(map[int]int)
	degreeFreq2 := make(map[int]int)
	vertexForDegree1 := make(map[int]int)
	vertexForDegree2 := make(map[int]int)
	for i := 0; i < n; i++ {
		d1 := g1.Degree(i)
		d2 := g2.Degree(i)
		degreeFreq1[d1]++
		degreeFreq2[d2]++
		vertexForDegree1[d1] = i
		vertexForDegree2[d2] = i
	}
	// For each degree that appears exactly once in both graphs, fix the mapping.
	for d, count := range degreeFreq1 {
		if count == 1 && degreeFreq2[d] == 1 {
			v1 := vertexForDegree1[d]
			v2 := vertexForDegree2[d]
			constraints[v1] = v2
		}
	}
	return constraints
}

// areIsomorphicSym returns true if g1 and g2 are isomorphic using symmetry breaking.
func areIsomorphicSym(g1, g2 *Graph) bool {
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
	constraints := uniqueConstraints(g1, g2)
	return backtrackSym(g1, g2, mapping, used, 0, constraints)
}

// findIsomorphismSym returns a valid mapping from g1 to g2 using symmetry breaking.
func findIsomorphismSym(g1, g2 *Graph) ([]int, bool) {
	n := g1.NumVertices
	mapping := make([]int, n)
	used := make([]bool, n)
	constraints := uniqueConstraints(g1, g2)
	if backtrackSym(g1, g2, mapping, used, 0, constraints) {
		res := make([]int, n)
		copy(res, mapping)
		return res, true
	}
	return nil, false
}
