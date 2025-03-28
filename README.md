

### Graph Isomorphism

Two graphs $G_1$ and $G_2$ are isomorphic if there exists a bijection $f: V_1 \to V_2$ such that
$$(u, v) \in E_1 \iff (f(u), f(v)) \in E_2.$$
This condition ensures that the structure of the graphs is preserved under the mapping.

### Automorphism Group

An automorphism of a graph $G = (V, E)$ is an isomorphism from $G$ to itself. The set of all automorphisms forms a group, denoted as $\text{Aut}(G)$, under function composition. For any $\phi, \psi \in \text{Aut}(G)$,

$$\phi \circ \psi \in \text{Aut}(G) \quad \text{and} \quad \phi^{-1} \in \text{Aut}(G).$$
The identity automorphism $\text{id}$ satisfies:
$$\text{id}(v) = v \quad \forall v \in V.$$

### Symmetry Reduction

The automorphism group can be used to partition the vertex set $V$ into orbits. For a vertex $v \in V$, its orbit under $\text{Aut}(G)$ is defined as:

$$\mathcal{O}(v) = \{\phi(v) \mid \phi \in \text{Aut}(G)\}$$

By selecting one representative from each orbit, we reduce the number of candidate mappings when checking for isomorphism. This is particularly useful in backtracking algorithms, where symmetry reduction can lead to significant improvements in performance.

### Invariant Properties

Graph invariants are properties that remain unchanged under any isomorphism. Common invariants include:
- **Degree Sequence:** The list of vertex degrees.
- **Eigenvalues of the Adjacency Matrix:** If $A$ is the adjacency matrix of $G$, its eigenvalues $\lambda_1, \lambda_2, \ldots, \lambda_n$ remain invariant under isomorphism.
- **Cycle Counts:** The number of cycles of a given length in the graph.

These invariants are used as preliminary checks to quickly rule out non-isomorphic graph pairs before attempting full isomorphism testing.


# Graph Isomorphism with Automorphism-based Symmetry Breaking in Go

## Motivation

Graph isomorphism is a challenging problem with many real-world applications, including network analysis, chemistry, and fraud detection. Traditional backtracking approaches suffer from factorial worst-case complexity. However, if a graph has many symmetries (automorphisms), we can exploit these by computing vertex orbits and imposing symmetry-breaking constraints during the isomorphism search. This project demonstrates such an approach in Go.

## Mathematical Theory

### Graph Isomorphism
  Two graphs $G_1$ and $G_2$ are isomorphic if there exists a bijection $f: V_1 \to V_2$ such that
  
  $$(u, v) \in E_1 \iff (f(u), f(v)) \in E_2.$$

  This condition ensures that the structure of the graphs is preserved under the mapping.

### Automorphism Group

An automorphism of a graph $G = (V, E)$ is an isomorphism from $G$ to itself. The set of all automorphisms forms a group, denoted as $\text{Aut}(G)$, under function composition. For any $\phi, \psi \in \text{Aut}(G)$,

$$\phi \circ \psi \in \text{Aut}(G) \quad \text{and} \quad \phi^{-1} \in \text{Aut}(G).$$
The identity automorphism $\text{id}$ satisfies:
$$\text{id}(v) = v \quad \forall v \in V.$$

### Orbits

The automorphism group can be used to partition the vertex set $V$ into orbits. For a vertex $v \in V$, its orbit under $\text{Aut}(G)$ is defined as:

$$\mathcal{O}(v) = \{\phi(v) \mid \phi \in \text{Aut}(G)\}$$

By selecting one representative from each orbit, we reduce the number of candidate mappings when checking for isomorphism. This is particularly useful in backtracking algorithms, where symmetry reduction can lead to significant improvements in performance.

- An **automorphism** of a graph is an isomorphism from the graph to itself.
- The set of all automorphisms forms a group, $\mathrm{Aut}(G)$.
- This group partitions the vertex set into **orbits**. Vertices in the same orbit are symmetric.
- By fixing the mapping for a representative (e.g., the smallest vertex) in each orbit during the isomorphism search, we can reduce the number of permutations to consider.

### Equations

1. **Isomorphism Condition:** $\forall u,v \in V, \quad A_{uv} = B_{f(u)f(v)}$
2. **Automorphism Condition:** $\forall u,v \in V, \quad A_{uv} = A_{\phi(u)\phi(v)}$
3. **Orbit Representative:** $\text{rep}(O) = \min(O)$

## Algorithm Pseudocode

```
function AUTOMORPHISMS(G):
    results = []
    backtrackAuto(G, mapping, used, 0, results)
    return results

function COMPUTE_ORBITS(G, automorphisms):
    uf = new UnionFind(n)
    for each automorphism mapping in automorphisms:
        for i from 0 to n-1:
            uf.union(i, mapping[i])
    orbits = group vertices by uf.find(vertex)
    return orbits

function REPRESENTATIVE_MAP(orbits, n):
    rep = array of booleans of size n
    for each orbit O in orbits:
         rep[min(O)] = true
    return rep

function BACKTRACK_SYM(G1, G2, mapping, used, index, rep):
    if index == n:
         return IS_MAPPING_VALID(G1, G2, mapping)
    if rep[index] is true:
         if not used[index]:
             mapping[index] = index
             mark index as used
             if BACKTRACK_SYM(G1, G2, mapping, used, index+1, rep) returns true:
                 return true
             unmark index
         return false
    else:
         for candidate in 0 to n-1:
              if candidate is not used:
                  mapping[index] = candidate
                  mark candidate as used
                  if BACKTRACK_SYM(G1, G2, mapping, used, index+1, rep) returns true:
                        return true
                  unmark candidate
         return false

function ARE_ISOMORPHIC_SYM(G1, G2):
    auts = AUTOMORPHISMS(G1)
    orbits = COMPUTE_ORBITS(G1, auts)
    rep = REPRESENTATIVE_MAP(orbits, n)
    return BACKTRACK_SYM(G1, G2, mapping, used, 0, rep)
```

## Results

### Test Case 1: Complete Graph \(K_4\)

**Graph 1 (0-indexed):**
```
4 6
0 1
0 2
0 3
1 2
1 3
2 3
```

**Graph 2 (1-indexed):**
```
4 6
1 2
1 3
1 4
2 3
2 4
3 4
```

*Result:* The graphs are isomorphic.

### Test Case 2: Disconnected Graphs

**Graph 1:**
```
6 4
0 1
2 3
4 5
1 0
```

**Graph 2:**
```
6 4
3 4
5 6
1 2
2 1
```

*Result:* The graphs are isomorphic.

### Test Case 3: Different Label Ordering

**Graph 1:**
```
4 4
0 1
1 2
2 3
3 0
```

**Graph 2:**
```
4 4
2 3
3 0
0 1
1 2
```

*Result:* The graphs are isomorphic.

## Implementation

- **Graph Representation:**  
  Implemented in `graph.go`, graphs are stored as an adjacency matrix.
  
- **Isomorphism Testing:**  
  In `isomorphism.go`, a backtracking algorithm tests for graph isomorphism while utilizing basic invariants.
  
- **Automorphism Computation:**  
  In `automorphism.go`, a stub function returns the identity automorphism, which can later be expanded to include full symmetry reduction.
  
- **Main Application:**  
  `main.go` creates sample graphs, runs the isomorphism test, and displays automorphisms.

### Manim Visualization

- **Visualization Module:**  
  The `manim/visualization.py` script uses Manim to animate the graph structure (a triangle) and demonstrates a sample permutation (swapping vertices).

### Build & Run Instructions

#### Go Application

1. **Prerequisites:**  
   - Install [Go](https://golang.org/dl/)

2. **Building and Running:**
   ```bash
   cd GraphIsomorphismGo
   go build -o graph_iso 
   .\graph_iso < test_cases/[choose test case] 
   ``` 
---


