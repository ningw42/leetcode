/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	} else {
		restSum := sum - root.Val
		if root.Left == nil && root.Right == nil {
			if restSum == 0 {
				return true
			} else {
				return false
			}
		} else if root.Left != nil && root.Right == nil { 
			return hasPathSum(root.Left, restSum)
		} else if root.Left == nil && root.Right != nil { 
			return hasPathSum(root.Right, restSum)
		} else {
			return hasPathSum(root.Left, restSum) || hasPathSum(root.Right, restSum)
		}
	}
}