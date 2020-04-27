/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	return iteration(root, k)
}

func iteration(root *TreeNode, k int) int {
	var stack []*TreeNode
	var count int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if count == k-1 {
            break
		}
		root = root.Right
		count++
	}
    return root.Val
}