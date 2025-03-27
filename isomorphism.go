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

// backtrack recursively tries all permutations (via backtracking)
// to find a valid mapping from g1 to g2.
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
			used[candidate] = false // backtrack
		}
	}
	return false
}

// areIsomorphic returns true if g1 and g2 are isomorphic.
func areIsomorphic(g1, g2 *Graph) bool {
	if g1.NumVertices != g2.NumVertices {
		return false
	}
	n := g1.NumVertices

	// Edge cases: empty graph or single vertex.
	if n == 0 || n == 1 {
		return true
	}

	// Compute and sort degrees for both graphs.
	deg1 := make([]int, n)
	deg2 := make([]int, n)
	for i := 0; i < n; i++ {
		sum1, sum2 := 0, 0
		for j := 0; j < n; j++ {
			sum1 += g1.AdjMatrix[i][j]
			sum2 += g2.AdjMatrix[i][j]
		}
		deg1[i] = sum1
		deg2[i] = sum2
	}
	// Sort the degree slices.
	// (Simple degree check: if the sorted degree multisets don't match, the graphs are not isomorphic.)
	quickSort(deg1)
	quickSort(deg2)
	for i := 0; i < n; i++ {
		if deg1[i] != deg2[i] {
			return false
		}
	}

	// Try to find an isomorphism via backtracking.
	mapping := make([]int, n)
	used := make([]bool, n)
	return backtrack(g1, g2, mapping, used, 0)
}

// findIsomorphism attempts to find and return a valid mapping from g1 to g2.
// Returns the mapping and true if found, or nil and false otherwise.
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

// quickSort is a simple quicksort implementation for []int.
func quickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	quickSort(a[:left])
	quickSort(a[left+1:])
}
