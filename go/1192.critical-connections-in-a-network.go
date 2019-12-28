/*
 * @lc app=leetcode id=1192 lang=golang
 *
 * [1192] Critical Connections in a Network
 */

// @lc code=start
var id int
var bridges [][]int
var visited []bool
var ids, low []int
var graph [][]int

func criticalConnections(n int, connections [][]int) [][]int {
	id = 0
	bridges = nil
	visited = make([]bool, n) // visited
	ids = make([]int, n) // id array
	low = make([]int, n) // low-link value array
	graph = buildGraph(n, connections)

	// tarjan's bridge finding algorithm
	dfs(0, -1)

	return bridges
}

func dfs(from int, parent int) {
	visited[from] = true
	ids[from] = id
	low[from] = id
	id++

	for _, to := range graph[from] {
		if to != parent {
			if visited[to] {
				// if 'to' is already visited
				// update low-link value of 'from' if possible
				low[from] = minInt(low[from], ids[to])
			} else {
				dfs(to, from)
				low[from] = minInt(low[from], low[to])
				if ids[from] < low[to] {
					bridges = append(bridges, []int{from, to})
				}
			}
		}
	}
}

func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
// @lc code=end

