/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root != nil {
		return append(
			[]int{root.Val},
			append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)...,
		)
	} else {
		return []int{}
	}
}
