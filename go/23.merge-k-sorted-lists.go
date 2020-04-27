/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for {
        next := math.MaxInt32
        for i, list := range lists {
            if list != nil {
                if next == math.MaxInt32 {
                    next = i
                } else {
                    if list.Val < lists[next].Val {
                        next = i
                    }
                }
            }
        }
        if next == math.MaxInt32 {
            return dummy.Next
        } else {
            current.Next, lists[next], current = lists[next], lists[next].Next, lists[next]
            
        }
    }
}