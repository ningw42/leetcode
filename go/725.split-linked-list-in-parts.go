import "fmt"

/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	if root == nil {
		var ret []*ListNode
		for i := 0; i < k; i++ {
			ret = append(ret, nil)
		}
		return ret
	}
	length := 0
	curr := root
	for curr != nil {
		curr = curr.Next
		length++
	}

	distribution := distribute(length, k)
	fmt.Println(distribution)

	var ret []*ListNode
	currentHead := root
	currentCount := 0
	currentTotal := distribution[0]
	distribution = distribution[1:]
	for root != nil {
		currentCount++
		if currentCount == currentTotal {
			ret = append(ret, currentHead)
			next := root.Next
			root.Next = nil
			root = next
			currentCount = 0
			if len(distribution) > 0 {
				currentTotal = distribution[0]
				distribution = distribution[1:]
			}
			currentHead = next
		} else {
			root = root.Next
		}
	}
	if length < k {
		for i := 0; i < len(distribution)+1; i++ {
			ret = append(ret, nil)
		}
	}

	return ret
}

func distribute(length int, k int) []int {
	result := length / k
	reminder := length % k

	var ret []int
	for i := 0; i < k; i++ {
		count := result
		if reminder > 0 {
			count += 1
			reminder--
		}
		ret = append(ret, count)
	}

	return ret
}

// @lc code=end

