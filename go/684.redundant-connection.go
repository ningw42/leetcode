/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	sets := make([]int, n+1)
	for i := 1; i < len(sets); i++ {
		sets[i] = i
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		fromRoot := find(sets, from)
		toRoot := find(sets, to)
		if fromRoot != toRoot {
			union(sets, fromRoot, toRoot)
		} else {
			return edge
		}
	}
	return nil
}

func find(sets []int, i int) int {
	for i != sets[i] {
		i = sets[i]
	}
	return i
}

func union(sets []int, a, b int) {
	sets[b] = a
}
// @lc code=end

