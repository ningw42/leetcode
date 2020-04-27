/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return sumup(root.Left, true) + sumup(root.Right, false)
	}
}

func sumup(r *TreeNode, l bool) int {
	if r == nil {
		return 0
	} else {
		if r.Left == nil && r.Right == nil {
			if l {
				return r.Val
			} else {
				return 0
			}
		} else {
			return sumup(r.Left, true) + sumup(r.Right, false)
		}
	}
}
