
type Graph struct {
	Edges map[int]map[int]int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	g := buildGraph(n, flights)
	K = K + 1 // covert from # nodes to # edges

	// return DFS(g, src, dst, K)
	return DP(g, src, dst, n, K)
}

func buildGraph(n int, flights [][]int) *Graph {
	edges := make(map[int]map[int]int)

	for _, flight := range flights {
		from := flight[0]
		to := flight[1]
		weight := flight[2]
		if _, e1 := edges[from]; e1 {
			edges[from][to] = weight
		} else {
			edges[from] = map[int]int{to: weight}
		}
	}

	return &Graph{Edges: edges}
}

// F(i, j, k) = F(i, a, k-1) + F(a, j, 1)
func DP(g *Graph, src, dst, n, k int) int {
	var results [][][]int
	for i := 0; i <= k; i++ {
		results = append(results, makeEmpty(n))
	}
	// 0
	for i := 0; i < n; i++ {
		results[0][i][i] = 0
	}
	// 1
	for from, tos := range g.Edges {
		for to, weight := range tos {
			results[1][from][to] = weight
		}
	}

	for stop := 2; stop <= k; stop++ {
		for from, tos := range g.Edges {
			for to, weight := range tos {
				for final, restWeight := range results[stop-1][to] {
					if restWeight >= 0 {
						newResult := weight + restWeight
						if existingResult := results[stop][from][final]; existingResult < 0 {
							results[stop][from][final] = newResult
						} else {
							if newResult < existingResult {
								results[stop][from][final] = newResult
							}
						}
					}
				}
			}
		}
	}

	min := math.MaxInt32
	for i := 0; i <= k; i++ {
		if weight := results[i][src][dst]; weight >= 0 && weight < min {
			min = weight
		}
	}
	if min == math.MaxInt32 {
		min = -1
	}

	return min
}
func makeEmpty(n int) [][]int {
	var ret [][]int
	for i := 0; i < n; i++ {
		ret = append(ret, make([]int, n))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ret[i][j] = -1
		}
	}

	return ret
}

// wont AC, time limit exceeded
func DFS(g *Graph, src, dst, k int) int {
	if src == dst {
		return 0
	} else {
		if k == 0 {
			return -1
		}
	}
	min := math.MaxInt32
	for next, weight := range g.Edges[src] {
		if weight > 0 {
			w := DFS(g, next, dst, k-1)
			if w >= 0 && w+weight < min {
				min = w + weight
			}
		}
	}
	if min == math.MaxInt32 {
		min = -1
	}

	return min
}
