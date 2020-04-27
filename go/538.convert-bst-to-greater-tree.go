/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    } else {
        convert(root, 0)
        return root
    }
}

func convert(root *TreeNode, s int) {
    if root.Left == nil && root.Right == nil {
        root.Val += s
    } else if root.Left != nil && root.Right == nil {
        convert(root.Left, s + root.Val)
        root.Val += s
    } else if root.Left == nil && root.Right != nil {
        rs := sum(root.Right)
        convert(root.Right, s)
        root.Val += s + rs
    } else {
        rs := sum(root.Right)
        convert(root.Left, s + root.Val + rs)
        convert(root.Right, s)
        root.Val += s + rs
    }
}

func sum(n *TreeNode) int {
    if n == nil {
        return 0
    } else {
        return sum(n.Left) + sum(n.Right) + n.Val
    }
} 