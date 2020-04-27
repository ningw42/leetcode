/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    length := 0
    curr := head
    for curr != nil {
        length++
        curr = curr.Next
    }
    curr = head
    num := 0
    for i := 1; i <= length; i++ {
        if curr.Val == 1 {
            num += int(math.Pow(2, float64(length-i)))
        }
        curr = curr.Next
    }
    return num
}