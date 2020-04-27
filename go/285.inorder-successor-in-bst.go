/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    child := childMin(p.Right)
    parent := parentMin(root, p, nil)
    if child == nil && parent == nil {
        return nil
    } else if child != nil && parent == nil {
        return child
    } else if child == nil && parent != nil {
        return parent
    } else {
        if child.Val < parent.Val {
            return child
        } else {
            return parent
        }
    }
}

func parentMin(root *TreeNode, p *TreeNode, path []*TreeNode) *TreeNode {
    if root == p {
        for i := len(path) - 1; i >= 0; i-- {
            if path[i].Val > p.Val {
                return path[i]
            }
        }
        return nil
    } else {
        path = append(path, root)
        if root.Val > p.Val {
            return parentMin(root.Left, p, path)
        } else {
            return parentMin(root.Right, p, path)
        }
    }
}

func childMin(root *TreeNode) *TreeNode {
    var curr, prev *TreeNode
    curr = root
    for curr != nil {
        curr, prev = curr.Left, curr
    }
    return prev
}