/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else {
		mid := len(nums) / 2
		return &TreeNode{
			Val: nums[mid],
			Left: sortedArrayToBST(nums[0:mid]),
			Right: sortedArrayToBST(nums[mid+1:]),
		}
	}
}