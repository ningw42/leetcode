/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var diameter int
func diameterOfBinaryTree(root *TreeNode) int {
    diameter = 0
    traversal(root)
    return diameter
}

func traversal(root *TreeNode) int {
    if root == nil {
        return 0
    } else {
        left := traversal(root.Left)
        right := traversal(root.Right)
        diameter = max(diameter, left + right)
        return max(left, right) + 1
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}