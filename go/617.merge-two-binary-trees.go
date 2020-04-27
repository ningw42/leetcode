/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return buildNewTree(t1, t2)
}

func buildNewTree(t1, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    } else if t1 != nil && t2 == nil {
        return t1
    } else if t1 == nil && t2 != nil {
        return t2
    } else {
        return &TreeNode{
            Val: t1.Val+t2.Val,
            Left: buildNewTree(t1.Left, t2.Left),
            Right: buildNewTree(t1.Right, t2.Right),
        }
    }
}