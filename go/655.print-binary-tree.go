

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
  