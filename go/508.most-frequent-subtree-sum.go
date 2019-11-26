/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
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

var frequentSums []int
var sumsCount map[int]int
var frequency int

func findFrequentTreeSum(root *TreeNode) []int {
	frequentSums = []int{}
	sumsCount = make(map[int]int)    
	frequency = 0

	treeSum(root)

	return frequentSums
}

func treeSum(root *TreeNode) int {
	if root != nil {
		var leftSum, rightSum int
		if root.Left != nil {
			leftSum = treeSum(root.Left)
		}
		if root.Right != nil {
			rightSum = treeSum(root.Right)
		}
		currentSum := leftSum + rightSum + root.Val
		addSum(currentSum)
		return currentSum
	} else {
		return 0
	}
}

func addSum(sum int) {
	sumsCount[sum]++
	if count := sumsCount[sum]; count > frequency {
		frequency = count
		frequentSums = []int{sum}
	} else if count == frequency {
		frequentSums = append(frequentSums, sum)
	}
}
// @lc code=end

