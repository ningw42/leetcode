/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    return useSlice(head)
}

func useSlice(head *ListNode) *ListNode {
    var slice []int
    for head != nil {
        slice = append(slice, head.Val)
        head = head.Next
    }
    
    slice = removeZeroSubslice(slice)
    
    dummy := &ListNode{}
    tail := dummy
    for _, value := range slice {
        tail.Next = &ListNode{Val:value}
        tail = tail.Next
    }
    
    return dummy.Next
}

func removeZeroSubslice(slice []int) []int {
    if len(slice) == 0 {
        return nil
    } 
    if len(slice) == 1 {
        if slice[0] == 0 {
            return nil
        } else {
            return slice
        }
    }
    
    for i := 0; i < len(slice); i++ {
        sum := 0
        for j := i; j < len(slice); j++ {
            sum += slice[j]
            if sum == 0 {
                return append(removeZeroSubslice(slice[:i]), removeZeroSubslice(slice[j+1:])...)
            }
        }
    }
    return slice
}