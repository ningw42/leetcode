/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    } else if len(nums) == 1 {
        return &TreeNode{Val: nums[0], Left:nil, Right: nil}
    } else {
        max := nums[0]
        pos := 0
        for i, v := range nums[0:] {
            if v > max {
                max = v
                pos = i
            }
        }
        
        return &TreeNode{Val:max, Left: constructMaximumBinaryTree(nums[0:pos]), Right: constructMaximumBinaryTree(nums[pos+1:])}
    }
}