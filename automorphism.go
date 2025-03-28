package main

// backtrackAuto recursively generates all automorphisms via backtracking.
func backtrackAuto(g *Graph, mapping []int, used []bool, index int, results *[][]int) {
	n := g.NumVertices
	if index == n {
		if isMappingValid(g, g, mapping) {
			perm := make([]int, n)
			copy(perm, mapping)
			*results = append(*results, perm)
		}
		return
	}
	for candidate := 0; candidate < n; candidate++ {
		if !used[candidate] {
			mapping[index] = candidate
			used[candidate] = true
			backtrackAuto(g, mapping, used, index+1, results)
			used[candidate] = false
		}
	}
}

// automorphisms returns all automorphisms (permutations) of graph g.
func automorphisms(g *Graph) [][]int {
	n := g.NumVertices
	results := [][]int{}
	mapping := make([]int, n)
	used := make([]bool, n)
	backtrackAuto(g, mapping, used, 0, &results)
	return results
}

// -------------------- Union-Find for Orbit Computation --------------------

// UF is a simple union-find structure.
type UF struct {
	parent []int
}

func newUF(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{parent: parent}
}

func (uf *UF) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UF) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		uf.parent[rootY] = rootX
	}
}

// computeOrbits computes vertex orbits from the automorphism group.
func computeOrbits(g *Graph, auts [][]int) [][]int {
	n := g.NumVertices
	uf := newUF(n)
	for _, mapping := range auts {
		for i := 0; i < n; i++ {
			uf.union(i, mapping[i])
		}
	}
	orbitMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := uf.find(i)
		orbitMap[root] = append(orbitMap[root], i)
	}
	var orbits [][]int
	for _, orbit := range orbitMap {
		orbits = append(orbits, orbit)
	}
	return orbits
}

// representativeMap returns a boolean slice rep where rep[i] is true if vertex i is the smallest in its orbit.
func representativeMap(orbits [][]int, n int) []bool {
	rep := make([]bool, n)
	for _, orbit := range orbits {
		min := orbit[0]
		for _, v := range orbit {
			if v < min {
				min = v
			}
		}
		rep[min] = true
	}
	return rep
}
