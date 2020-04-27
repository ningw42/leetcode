/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    leftBoundary := left(root.Left, nil)
    rightBoundary := right(root.Right)
    leftLeaves := leaf(root.Left)
    rightLeaves := leaf(root.Right)
   
    boundary := []int{root.Val}
    boundary = append(boundary, leftBoundary...)
    if len(leftBoundary) > 0 {
        boundary = append(boundary, leftLeaves[1:]...)
    }
    boundary = append(boundary, rightLeaves...)
    if len(rightLeaves) > 0 {
        boundary = append(boundary, rightBoundary[1:]...)
    }
   
    return boundary
}

func left(root *TreeNode, ret []int) []int {
    if root == nil {
        return ret
    }
    ret = append(ret, root.Val)
    if root.Left != nil {
        return left(root.Left, ret)
    } else {
        return left(root.Right, ret)
    }
}

func right(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var ret []int
    if root.Right != nil {
        ret = right(root.Right)
    } else {
        ret = right(root.Left)
    }
    ret = append(ret, root.Val)
    return ret
}

func leaf(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    } else {
        return append(leaf(root.Left), leaf(root.Right)...)
    }
}