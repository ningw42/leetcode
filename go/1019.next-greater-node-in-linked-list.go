/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    // reverse linked-list with monotone stack(descending)
    length, head := reverseLinkedList(head)
    if length == 0 {
        return nil
    } else if length == 1 {
        return []int{0}
    }
    
    var stack, result []int
    for head != nil {
        for len(stack) > 0 && stack[len(stack)-1] <= head.Val {
            stack = stack[:len(stack)-1] // pop from stack
        }
        if len(stack) == 0 {
            result = append(result, 0)
        } else {
            result = append(result, stack[len(stack)-1])
        }
        stack = append(stack, head.Val)
        head = head.Next
    }
    reverseSlice(result)
    return result
}

func reverseLinkedList(head *ListNode) (int, *ListNode) {
    // empty linked list
    if head == nil {
        return 0, nil
    }
    
    current := head.Next
    previous := head
    count := 1
   
    // single node
    if current == nil {
        return 1, head
    }
 
    // head becomes tail
    previous.Next = nil
    for current != nil {
        current, previous, current.Next = current.Next, current, previous
        count++
    }
    
    return count, previous
}

func reverseSlice(s []int) {
    var i, j int
    for ; i < len(s) / 2; i++ {
        j = len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}