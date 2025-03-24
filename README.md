# Graph Isomorphism Testing via Symmetry Reduction

## Overview

This project demonstrates a graph isomorphism testing algorithm enhanced with symmetry reduction techniques using group theory. The project is primarily implemented in C++ for performance, with a Python-based visualization module to illustrate graph symmetries and the isomorphism testing process through interactive animations.

## Problem Definition

Graph isomorphism is the problem of determining whether two finite graphs are isomorphic. Two graphs $G_1$ and $G_2$ are isomorphic if there exists a bijection between their vertex sets such that edge connectivity is preserved. Formally, if

$$G_1 = (V_1, E_1) \quad \text{and} \quad G_2 = (V_2, E_2)$$
then $G_1\simeq G_2$ if there exists a bijective function
$$f: V_1 \rightarrow V_2$$
such that
$$\forall u, v \in V_1,\quad (u, v) \in E_1 \iff (f(u), f(v)) \in E_2.$$

## Motivation

- **Efficiency:** Testing graph isomorphism can be computationally intensive. Leveraging group theory to identify and exploit graph symmetries helps reduce redundant computations.
- **Visualization:** An interactive visual approach aids in understanding how the algorithm processes graphs and handles symmetries.
- **Interdisciplinary Learning:** This project bridges concepts from abstract algebra, algorithm design, and computer graphics, making it a rich educational tool.

## Theory

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

## Implementation

### C++ Components

- **Graph Representation:**  
  Implemented in `Graph.h` and `Graph.cpp`, graphs are stored using an adjacency matrix.

- **Isomorphism Testing:**  
  In `Isomorphism.h` and `Isomorphism.cpp`, a backtracking algorithm tests for isomorphism by trying all vertex mappings while using basic degree invariants as a preliminary filter.

- **Automorphism Computation:**  
  The functions in `Automorphism.h` and `Automorphism.cpp` currently return only the identity automorphism, serving as a starting point for future symmetry reductions.

- **Main Application:**  
  `main.cpp` creates sample graphs, runs the isomorphism test, and displays the automorphisms.

### Python Visualization

- **Visualization Module:**  
  `vis/visualization.py` uses NetworkX and Matplotlib to animate the layout of a sample graph, providing a visual understanding of graph structure and potential symmetry.



### Build & Run Instructions

1. **Prerequisites:**  
   - C++ compiler with C++11 support  
   - CMake  
   - Python 3 with `networkx` and `matplotlib` (install via `pip install networkx matplotlib`)

2. **Building the C++ Application:**
   ```bash
   mkdir build
   cd build
   cmake ..
   make
   ./GraphIsomorphismProject
