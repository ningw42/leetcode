/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		stack := [][]*TreeNode{[]*TreeNode{root.Left, root.Right}}
		for len(stack) > 0 {
			var nodes []*TreeNode
			// fmt.Println(stack)
			nodes, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
			l := nodes[0]
			r := nodes[1]
			if l == nil && r == nil {
				continue
			} else if l != nil && r != nil {
				if l.Val == r.Val {
					stack = append(stack, []*TreeNode{l.Left, r.Right})
					stack = append(stack, []*TreeNode{l.Right, r.Left})
				} else {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
}


