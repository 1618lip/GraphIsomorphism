package main

// uniqueDegreeConstraints returns a map from vertex in g1 to vertex in g2 for vertices that have a unique degree.
func uniqueDegreeConstraints(g1, g2 *Graph) map[int]int {
	constraints := make(map[int]int)
	n := g1.NumVertices

	// Compute frequency maps for degrees.
	degreeFreq1 := make(map[int]int)
	degreeFreq2 := make(map[int]int)
	// Also store the vertex with that degree (if unique).
	vertexOfDegree1 := make(map[int]int)
	vertexOfDegree2 := make(map[int]int)
	for i := 0; i < n; i++ {
		d1 := g1.Degree(i)
		degreeFreq1[d1]++
		vertexOfDegree1[d1] = i

		d2 := g2.Degree(i)
		degreeFreq2[d2]++
		vertexOfDegree2[d2] = i
	}

	// For each degree that appears exactly once in both graphs, fix the mapping.
	for d, count1 := range degreeFreq1 {
		if count1 == 1 {
			if degreeFreq2[d] == 1 {
				constraints[vertexOfDegree1[d]] = vertexOfDegree2[d]
			}
		}
	}
	return constraints
}
