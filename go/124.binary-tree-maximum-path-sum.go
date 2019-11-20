/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

var max, maxVal int
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	} 
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	max = math.MinInt64
	maxVal = math.MinInt64
	maxSum(root)
	if max == 0 {
		return maxVal
	} else {
		return max
	}
}

func maxSum(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		if root.Val > maxVal {
			maxVal = root.Val
		}
		left := maxSum(root.Left)
		right := maxSum(root.Right)
		if left > max {
			max = left
		}
		if right > max {
			max = right
		}
		if sum := right + left + root.Val; sum > max {
			max = sum
		}

		maxSubSum := left
		if right > maxSubSum {
			maxSubSum = right
		}

		if sum := maxSubSum + root.Val; sum > 0 {
			return sum
		} else {
			return 0
		}
	}
}
