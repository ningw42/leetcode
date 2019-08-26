/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

 type DirectedGraph struct {
	NumberOfNodes int // number of nodes
	Matrix [][]bool // adjacency matrix
	In []int // inbound edge count
}

func NewDirectedGraph(numberOfNodes int) *DirectedGraph {
	var mat [][]bool
	for i := 0; i < numberOfNodes; i++ {
		mat = append(mat, make([]bool, numberOfNodes))
	}
	return &DirectedGraph{
		NumberOfNodes: numberOfNodes,
		Matrix: mat,
		In: make([]int, numberOfNodes),
	}
}

func (dg *DirectedGraph) AddEdge(from int, to int) {
	dg.Matrix[from][to] = true
	dg.In[to]++
}

// DFS to determine if there is a cycle
// see https://en.wikipedia.org/wiki/Topological_sorting
// L ← Empty list that will contain the sorted elements
// S ← Set of all nodes with no incoming edge
// while S is non-empty do
//     remove a node n from S
//     add n to tail of L
//     for each node m with an edge e from n to m do
//         remove edge e from the graph
//         if m has no other incoming edges then
//             insert m into S
// if graph has edges then
//     return error   (graph has at least one cycle)
// else 
//     return L   (a topologically sorted order)
func (dg DirectedGraph) TopologicalSort() []int {
	var s, l []int
	for node, incomingEdgeCount := range dg.In {
		if incomingEdgeCount == 0 {
			s = append(s, node)
		}
	}

	for len(s) > 0 {
		var node int
		node, s = s[0], s[1:]
		l = append(l, node)
		for to, exists := range dg.Matrix[node] {
			if exists {
				dg.Matrix[node][to] = false
				dg.In[to]--
				if dg.In[to] == 0 {
					s = append(s, to)
				}
			}
		}		
	}

	for _, adjacencyRow := range dg.Matrix {
		for _, exists := range adjacencyRow {
			if exists {
				return []int{}
			}
		}
	}

	return l
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	dg := NewDirectedGraph(numCourses)
	for _, edge := range prerequisites {
		dg.AddEdge(edge[1], edge[0])
	}

	return dg.TopologicalSort()
}

