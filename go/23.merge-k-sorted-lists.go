import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (30.70%)
 * Total Accepted:    282.9K
 * Total Submissions: 921.5K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	allNil := true
	for _, list := range lists {
		if list != nil {
			allNil = false
			break
		}
	}
	if allNil {
		return nil
	}

	head := &ListNode{}
	previous := head

	for end, indexes := getIndexesOfMinValue(lists); !end; {
		nextNodes := make(map[int]*ListNode)
		// record the next node
		for _, index := range indexes {
			nextNodes[index] = lists[index].Next
		}
		// link all nodes
		cursor := lists[indexes[0]]
		previous.Next = cursor
		for _, index := range indexes[1:] {
			cursor.Next = lists[index]
			cursor = cursor.Next
		}
		previous = cursor
		// move cursor to the next node
		for _, index := range indexes {
			lists[index] = nextNodes[index]
		}
		end, indexes = getIndexesOfMinValue(lists)
	}

	return head.Next
}

// find all the indexes of the minium value
func getIndexesOfMinValue(lists []*ListNode) (bool, []int) {
	m := make(map[int]int)

	for index, node := range lists {
		if node != nil {
			m[index] = node.Val
		}
	}
	if len(m) == 0 {
		return true, []int{}
	}

	min := 0x7FFFFFFF
	invertedMap := make(map[int][]int)
	for index, value := range m {
		if value < min {
			min = value
		}
		invertedMap[value] = append(invertedMap[value], index)
	}

	return false, invertedMap[min]
}

func PrintLists(lists []*ListNode) {
	result := "["
	for _, list := range lists {
		result += "[" + PrintList(list) + "],"
	}

	result = strings.Trim(result, ",") + "]"

	fmt.Println(result)
}

func PrintList(l *ListNode) string {
	result := ""
	for c := l; c != nil; c = c.Next {
		result = result + (strconv.Itoa(c.Val) + ",")
	}

	return strings.Trim(result, ",")
}
