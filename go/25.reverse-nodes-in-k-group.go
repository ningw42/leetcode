/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	// collect all nodes
	stash := make([]*ListNode, 0)

	current := head
	length := 0
	for current != nil {
		stash = append(stash, current)
		current = current.Next
		length++
	}

	if k > length {
		return head
	}

	var newHead *ListNode
	var lastKTail *ListNode
	i := 0
	// reverse every k elements
	for ; i + k <= length; i = i + k {
		kHead, kTail := helper(stash[i:i+k])
		if newHead == nil {
			newHead = kHead
		} else {
			lastKTail.Next = kHead
		}
		lastKTail = kTail
	}

	// append tail if necessary
	if length % k != 0 {
		lastKTail.Next = stash[i]
	}

	return newHead
}

func helper(nodes []*ListNode) (head *ListNode, tail *ListNode) {
	length := len(nodes)
	first := nodes[length - 1]
	last := nodes[0]

	head = first
	tail = last

	current := head
	for i := length - 2; i > 0; i-- {
		current.Next = nodes[i]
		current = current.Next
	}
	current.Next = tail	
	tail.Next = nil

	return
}

func printList(head *ListNode) {
	fmt.Println()
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func printSlice(s []*ListNode) {
	fmt.Println()
	for _, n := range s {
		fmt.Println(n.Val)
	}
}


