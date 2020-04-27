/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, balanced := getHeight(root)
	return balanced
}

func getHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	} else {
		var height int
		var balanced bool
		lh, lb := getHeight(root.Left)
		rh, rb := getHeight(root.Right)
		// height
		if lh > rh {
			height = lh + 1
		} else {
			height = rh + 1
		}
		// balanced
		if lb && rb && math.Abs(float64(lh - rh)) <= 1 {
			balanced = true
		}

		return height, balanced
	}
}