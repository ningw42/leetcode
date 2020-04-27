/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    if root == nil {
        return 0
    } else {
        if root.Left == nil && root.Right == nil {
            return 0
        } else if root.Left != nil && root.Right == nil {
            return int(math.Abs(float64(sum(root.Left)))) + findTilt(root.Left)
        } else if root.Left == nil && root.Right != nil {
            return int(math.Abs(float64(sum(root.Right)))) + findTilt(root.Right)
        } else {
            return int(math.Abs(float64(sum(root.Right) - sum(root.Left)))) + findTilt(root.Right) + findTilt(root.Left)
        }
    }
}

func sum(root *TreeNode) int {
    if root == nil {
        return 0
    } else {
        return sum(root.Left) + sum(root.Right) + root.Val
    }
}