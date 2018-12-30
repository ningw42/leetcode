import "sort"

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 *
 * https://leetcode.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (33.16%)
 * Total Accepted:    106.2K
 * Total Submissions: 319.2K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * Two elements of a binary search tree (BST) are swapped by mistake.
 *
 * Recover the tree without changing its structure.
 *
 * Example 1:
 *
 *
 * Input: [1,3,null,null,2]
 *
 * 1
 * /
 * 3
 * \
 * 2
 *
 * Output: [3,1,null,null,2]
 *
 * 3
 * /
 * 1
 * \
 * 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,4,null,null,2]
 *
 * ⁠ 3
 * ⁠/ \
 * 1   4
 * /
 * 2
 *
 * Output: [2,1,4,null,null,3]
 *
 * ⁠ 2
 * ⁠/ \
 * 1   4
 * /
 * ⁠ 3
 *
 *
 * Follow up:
 *
 *
 * A solution using O(n) space is pretty straight forward.
 * Could you devise a constant space solution?
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	raw := traversal(root)
	sorted := make([]int, len(raw))
	copy(sorted, raw)
	sort.Ints(sorted)

	var a, b *TreeNode
	for i, e := range raw {
		if e != sorted[i] {
			a = find(root, e)
			b = find(root, sorted[i])
			break
		}
	}

	t := a.Val
	a.Val = b.Val
	b.Val = t
}

func traversal(t *TreeNode) []int {
	if t == nil {
		return []int{}
	} else {
		return append(append(traversal(t.Left), t.Val), traversal(t.Right)...)
	}
}

func find(t *TreeNode, e int) *TreeNode {
	if t == nil {
		return nil
	} else {
		if t.Val == e {
			return t
		} else {
			n := find(t.Left, e)
			if n == nil {
				return find(t.Right, e)
			} else {
				return n
			}
		}
	}
}