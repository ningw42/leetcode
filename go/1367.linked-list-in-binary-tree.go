/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
    // find candidates
    var cs []*TreeNode
  
    // iterative inorder traversal
    var stack []*TreeNode
    current := root
    for current != nil || len(stack) > 0 {
        // left
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        // root
        current = stack[len(stack)-1] // top
        stack = stack[:len(stack)-1] // pop
        if current.Val == head.Val {
            cs = append(cs, current)
        }
        // right
        current = current.Right
    }
 
    // try out each candidates
    for _, c := range cs {
        if isPath(head, c) {
            return true
        }
    } 
    return false
}

func isPath(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    } else {
        if root == nil {
            return false 
        } else {
            if head.Val == root.Val {
                return isPath(head.Next, root.Left) || isPath(head.Next, root.Right)
            } else {
                return false
            }
        }
    }
}
