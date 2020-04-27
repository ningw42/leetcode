/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    } else {
        if root.Left == nil && root.Right == nil {
            return []int{root.Val}
        } else if root.Left != nil && root.Right == nil {
            return append([]int{root.Val}, rightSideView(root.Left)...)
        } else if root.Left == nil && root.Right != nil {
            return append([]int{root.Val}, rightSideView(root.Right)...)
        } else {
            l := rightSideView(root.Left)
            r := rightSideView(root.Right)
            if len(l) > len(r) {
                return append(append([]int{root.Val}, r...), l[len(r):]...)
            } else {
                return append([]int{root.Val}, r...)
            }
        }
    }
}