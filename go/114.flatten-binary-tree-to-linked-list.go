/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    do(root)
}

func do(r *TreeNode) *TreeNode {
	if r == nil {
		return nil
	} else {
		if r.Left == nil && r.Right == nil {
			return r
		} else if r.Left == nil && r.Right != nil {
			return do(r.Right)
		} else if r.Left != nil && r.Right == nil {
			r.Right = r.Left
			r.Left = nil
			return do(r.Right)
		} else {
			rt := r.Right
			r.Right = r.Left
			r.Left = nil
			do(r.Right).Right = rt
			return do(rt)
		}
	}
}