/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return FollowUp(head)
}

// O(n) time complexity and O(1) space complexity
func FollowUp(head *ListNode) bool {
    // 1. get length O(n)
    current := head
    length := 0
    for current != nil {
        length++
        current = current.Next
    }
    if length == 0 || length == 1 {
        return true
    }
    
    // 2. split list O(n)
    var left, right *ListNode
    left = head
    current = left
    for i := 0; i < length / 2 - 1; i++ {
        current = current.Next
    }
    if length % 2 == 1 {
        right = current.Next.Next
    } else {
        right = current.Next
    }
    current.Next = nil
    
    // 3. reverse the second list O(n)
    right = reverse(right)
    
    // 4. compare node by node
    for left != nil && right != nil {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
    // empty list
    if head == nil {
        return nil
    }
    
    current := head.Next
    previous := head
   
    // single node
    if current == nil {
        return head
    }
   
    previous.Next = nil
    for current != nil {
        current, previous, current.Next = current.Next, current, previous
    }
    
    return previous
}

func Origin(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i, j := 0, len(list)-1; i < j; {
		if list[i] != list[j] {
			return false
		}
		i++
		j--
	}

	return true
}

