/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count int
var target int
func pathSum(root *TreeNode, sum int) int {
	count = 0
	target = sum
	if root == nil {
		return 0
	}

	sums(root, nil)
	return count
}

func sums(root *TreeNode, upperSums map[int]int) {
	if root != nil {
		newUpperSums := addValueToUpperSums(upperSums, root.Val)
		count += newUpperSums[target]
		sums(root.Left, replicateUpperSums(newUpperSums))
		sums(root.Right, replicateUpperSums(newUpperSums))
	}
}

func addValueToUpperSums(upperSums map[int]int, value int) map[int]int {
	newUpperSums := make(map[int]int)
	for sum, count := range upperSums {
		newUpperSums[sum+value] += count
	}
	newUpperSums[value] += 1

	return newUpperSums
}

func replicateUpperSums(upperSums map[int]int) map[int]int {
	replica := make(map[int]int)
	for sum, count := range upperSums {
		replica[sum] = count
	}
	return replica
}
// @lc code=end

