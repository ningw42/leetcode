import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=655 lang=golang
 *
 * [655] Print Binary Tree
 *
 * https://leetcode.com/problems/print-binary-tree/description/
 *
 * algorithms
 * Medium (49.80%)
 * Total Accepted:    15.3K
 * Total Submissions: 30.7K
 * Testcase Example:  '[1,2]'
 *
 * Print a binary tree in an m*n 2D string array following these rules:
 *
 *
 * The row number m should be equal to the height of the given binary tree.
 * The column number n should always be an odd number.
 * The root node's value (in string format) should be put in the exactly middle
 * of the first row it can be put. The column and the row where the root node
 * belongs will separate the rest space into two parts (left-bottom part and
 * right-bottom part). You should print the left subtree in the left-bottom
 * part and print the right subtree in the right-bottom part. The left-bottom
 * part and the right-bottom part should have the same size. Even if one
 * subtree is none while the other is not, you don't need to print anything for
 * the none subtree but still need to leave the space as large as that for the
 * other subtree. However, if two subtrees are none, then you don't need to
 * leave space for both of them.
 * Each unused space should contain an empty string "".
 * Print the subtrees following the same rules.
 *
 *
 * Example 1:
 *
 * Input:
 * ⁠    1
 * ⁠   /
 * ⁠  2
 * Output:
 * [["", "1", ""],
 * ⁠["2", "", ""]]
 *
 *
 *
 *
 * Example 2:
 *
 * Input:
 * ⁠    1
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \
 * ⁠    4
 * Output:
 * [["", "", "", "1", "", "", ""],
 * ⁠["", "2", "", "", "", "3", ""],
 * ⁠["", "", "4", "", "", "", ""]]
 *
 *
 *
 * Example 3:
 *
 * Input:
 * ⁠     1
 * ⁠    / \
 * ⁠   2   5
 * ⁠  /
 * ⁠ 3
 * ⁠/
 * 4
 * Output:
 *
 * [["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 * ⁠["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 * ⁠["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 * ⁠["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
 *
 *
 *
 * Note:
 * The height of binary tree is in the range of [1, 10].
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

const MARK string = ""

func printTree(root *TreeNode) [][]string {
	return print(root, getHeight(root))
}

func getHeight(root *TreeNode) uint {
	if root == nil {
		return 0
	} else {
		left := getHeight(root.Left)
		right := getHeight(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func print(root *TreeNode, height uint) [][]string {
	if height == 1 {
		var val string
		if root != nil {
			val = strconv.Itoa(root.Val)
		} else {
			val = MARK
		}
		return [][]string{[]string{val}}
	} else {
		height--
		var result, left, right [][]string
		var row, rowPadding []string
		var val string
		if root != nil {
			val = strconv.Itoa(root.Val)
			left = print(root.Left, height)
			right = print(root.Right, height)
		} else {
			val = MARK
			left = print(nil, height)
			right = print(nil, height)
		}

		half := 1<<height - 1
		// row
		for i := 0; i < half; i++ {
			rowPadding = append(rowPadding, MARK)
		}
		row = append(rowPadding, val)
		row = append(row, rowPadding...)
		result = append(result, row)

		fmt.Println(left)
		fmt.Println(right)

		for i := 0; i < len(left); i++ {
			pRow := append(left[i], MARK)
			pRow = append(pRow, right[i]...)
			result = append(result, pRow)
		}

		return result
	}
}
  