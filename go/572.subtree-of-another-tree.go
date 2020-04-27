/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    for _, subTree := range findPossible(s, t.Val) {
        if isIdentical(t, subTree) {
            return true
        }
    }
    return false
}

func findPossible(root *TreeNode, val int) []*TreeNode{
    if root == nil {
        return nil
    } else {
        var ret []*TreeNode
        if root.Val == val {
            ret = append(ret, root)
        }
        l := findPossible(root.Left, val)
        r := findPossible(root.Right, val)
        return append(ret, append(l, r...)...)
    }
}

func isIdentical(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    } else if a == nil && b != nil {
        return false
    } else if a != nil && b == nil {
        return false
    } else {
        return a.Val == b.Val && isIdentical(a.Left, b.Left) && isIdentical(a.Right, b.Right)
    }
}
