import (
	"math"
)

/*
 * @lc app=leetcode id=1104 lang=golang
 *
 * [1104] Path In Zigzag Labelled Binary Tree
 */

// @lc code=start
func pathInZigZagTree(label int) []int {
	var targetIndex int
	tree := []int{0}
	for i := 1; true; i++ {
		// build tree
		lower := int(math.Pow(2, float64(i-1)))
		upper := int(math.Pow(2, float64(i)))
		layer := make([]int, lower)
		var index int
		var found bool
		for j := lower; j < upper; j++ {
			if i&1 == 1 {
				index = j - lower
			} else {
				index = lower - 1 - (j - lower)
			}
			layer[index] = j
			if j == label {
				found = true
				break
			}
		}
		if found {
			targetIndex = len(tree) + index
			tree = append(tree, layer...)
			break
		}
		tree = append(tree, layer...)
	}

	// retrieve path element
	var reversed []int
	for targetIndex >= 1 {
		reversed = append(reversed, tree[targetIndex])
		targetIndex = targetIndex / 2
	}
	for i := 0; i < len(reversed)/2; i++ {
		reversed[i], reversed[len(reversed)-i-1] = reversed[len(reversed)-i-1], reversed[i]
	}

	return reversed
}

// @lc code=end

