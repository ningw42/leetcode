/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

var Ancestor *TreeNode
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	Ancestor = nil
	containsPQ(root, p, q)
	return Ancestor
}
 
func containsPQ(root, p, q *TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	} else {
		P := root == p
		Q := root == q
		leftP, leftQ := containsPQ(root.Left, p, q)
		rightP, rightQ := containsPQ(root.Right, p, q)
 
		P = P || leftP || rightP
		Q = Q || leftQ || rightQ
 
		if P && Q && Ancestor == nil {
			Ancestor = root
		}
 
		return P, Q
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    } else {
        if root == p || root == q {
            return root
        } else {
            leftHasPorQ := lowestCommonAncestor(root.Left, p, q)
            rightHasPorQ := lowestCommonAncestor(root.Right, p, q)
            if leftHasPorQ != nil && rightHasPorQ != nil {
                return root
            } else if leftHasPorQ != nil && rightHasPorQ == nil {
                return leftHasPorQ
            } else if leftHasPorQ == nil && rightHasPorQ != nil {
                return rightHasPorQ
            } else {
                return nil
            }
        }
    }
}
// @lc code=end

