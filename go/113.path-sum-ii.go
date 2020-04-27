/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	} else {
		restSum := sum - root.Val
		var result [][]int
		if root.Left == nil && root.Right == nil {
			if restSum == 0 {
				return [][]int{[]int{root.Val}}
			} else {
				return nil
			}
		} else if root.Left != nil && root.Right == nil { 
			result = pathSum(root.Left, restSum)
		} else if root.Left == nil && root.Right != nil { 
			result = pathSum(root.Right, restSum)
		} else {
			result = append(pathSum(root.Left, restSum), pathSum(root.Right, restSum)...)
		}
		if result == nil {
			return nil
		} else {
			var ret [][]int
			for _, path := range result {
				ret = append(ret, append([]int{root.Val}, path...))
			}
			return ret
		}
	} 
}