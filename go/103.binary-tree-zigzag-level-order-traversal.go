/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)
	Traversal(root, 0, m)

	ret := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		if i%2 == 1 {
			for j := len(m[i]) - 1; j >= 0; j-- {
				ret[i] = append(ret[i], m[i][j])
			}
		} else {
			ret[i] = m[i]
		}
	}

	return ret
}

func Traversal(root *TreeNode, level int, m map[int][]int) {
	if root != nil {
		Traversal(root.Left, level+1, m)
		Traversal(root.Right, level+1, m)
		if _, exists := m[level]; exists {
			m[level] = append(m[level], root.Val)
		} else {
			m[level] = []int{root.Val}
		}
	}
}

// @lc code=end

