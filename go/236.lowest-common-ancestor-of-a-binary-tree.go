/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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