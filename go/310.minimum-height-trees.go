import (
	"math"
)

/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

func findMinHeightTrees(n int, edges [][]int) []int {
	return Clever(n, edges)
}

func Clever(n int, edges [][]int) []int {
	// https://www.geeksforgeeks.org/roots-tree-gives-minimum-height/
	// Start pointers from all leaf nodes and move one step inside each time,
	// keep combining pointers which overlap while moving,
	// at the end only one pointer will remain on some vertex or two pointers will remain at one distance away.
	// Those node represent the root of the vertex which will minimize the height of the tree.

	if n == 1 {
		return []int{0}
	}
	// build graph
	var graph [][]bool
	degrees := make([]int, n)
	for i := 0; i < n; i++ {
		graph = append(graph, make([]bool, n))
	}
	for _, edge := range edges {
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
		degrees[edge[0]] = degrees[edge[0]] + 1
		degrees[edge[1]] = degrees[edge[1]] + 1
	}

	// BFS
	var queue []int
	// add all leaf nodes into queue
	for node, degree := range degrees {
		if degree == 1 {
			queue = append(queue, node)
		}
	}
	for n > 2 {
		numberOfLeafNodes := len(queue)
		for i := 0; i < numberOfLeafNodes; i++ {
			v1 := queue[0]
			queue = queue[1:]
			n--
			for v2, connected := range graph[v1] {
				if connected {
					// remove edge
					graph[v1][v2] = false
					graph[v2][v1] = false
					degrees[v1] = degrees[v1] - 1
					degrees[v2] = degrees[v2] - 1
					// if v2 becomes leaf node after removing edge from v1
					// add v2 to queue
					if degrees[v2] == 1 {
						queue = append(queue, v2)
					}
				}
			}
		}
	}

	return queue
}

func Naive(n int, edges [][]int) []int {
	// build graph
	var g [][]bool
	for i := 0; i < n; i++ {
		g = append(g, make([]bool, n))
	}
	for _, edge := range edges {
		g[edge[0]][edge[1]] = true
		g[edge[1]][edge[0]] = true
	}

	// find all possible height and find the min
	var ret []int
	minHeight := math.MaxInt32
	for i := 0; i < n; i++ {
		if height := getHeight(copyGraph(g), i); height < minHeight {
			minHeight = height
			ret = []int{i}
		} else if height == minHeight {
			ret = append(ret, i)
		}
	}

	return ret
}

func copyGraph(graph [][]bool) [][]bool {
	duplicate := make([][]bool, len(graph))
	for i := range graph {
		duplicate[i] = make([]bool, len(graph[i]))
		copy(duplicate[i], graph[i])
	}

	return duplicate
}

func getHeight(graph [][]bool, root int) int {
	height := 0
	for node, connected := range graph[root] {
		if connected {
			graph[root][node] = false
			graph[node][root] = false
			if h := getHeight(graph, node) + 1; h > height {
				height = h
			}
		}
	}

	return height
}

