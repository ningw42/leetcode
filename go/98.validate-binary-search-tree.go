/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var prev *TreeNode

func isValidBST(root *TreeNode) bool {
    prev = nil
    return isBST(root)
}

func isBST(root *TreeNode) bool {
    if root == nil {
        return true
    } else {
        if !isBST(root.Left) {
            return false
        }
        if prev != nil && prev.Val >= root.Val {
            return false
        }
        prev = root
        return isBST(root.Right)
    }
}
