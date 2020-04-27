/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
    parent := parentMin(node)
    child := childMin(node.Right)
    if parent == nil && child == nil {
        return nil
    } else if parent != nil && child == nil {
        return parent
    } else if parent == nil && child != nil {
        return child
    } else {
        if child.Val < parent.Val {
            return child
        } else {
            return parent
        }
    }
}

func parentMin(root *Node) *Node {
    curr := root
    for curr != nil && curr.Val <= root.Val {
        curr = curr.Parent
    }
    return curr
}

func childMin(root *Node) *Node {
    var curr, prev *Node
    curr = root
    for curr != nil {
        curr, prev = curr.Left, curr
    }
    return prev
}
