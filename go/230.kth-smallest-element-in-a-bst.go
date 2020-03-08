/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	return iteration(root, k)
}

func iteration(root *TreeNode, k int) int {
	var stack []*TreeNode
	var count int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if count == k-1 {
			return root.Val
		}
		root = root.Right
		count++
	}
}

func recursive(root *TreeNode, k int) int {
	ordered := inorder(root)
	return ordered[k-1]
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		return append(inorder(root.Left), append([]int{root.Val}, inorder(root.Right)...)...)
	}
}
// @lc code=end

