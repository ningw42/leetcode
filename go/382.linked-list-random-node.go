/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    capacity int
    reservoir []*ListNode
    cursor *ListNode
    head *ListNode
    count int
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    capacity := 10
    count := 1
    var reservoir []*ListNode
    cursor := head
    for i := 0; cursor != nil && i < capacity; i++ {
        reservoir = append(reservoir, cursor)
        cursor = cursor.Next
        count++
    }
    // fmt.Println(reservoir)
    return Solution{
        capacity: capacity,
        reservoir: reservoir,
        cursor: cursor,
        head: head,
        count: count,
    }
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    if this.capacity > len(this.reservoir) {
        return this.reservoir[rand.Intn(len(this.reservoir))].Val
    } else {
        if this.cursor == nil {
            this.cursor = this.head
        }
        index := rand.Intn(this.capacity)
        value := this.reservoir[index].Val
        random := rand.Intn(this.count) // [0, count]
        if random < this.capacity {
            this.reservoir[random] = this.cursor
        }
        this.count++ // increase count by 1
        this.cursor = this.cursor.Next // advance pointer
        return value
    }
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */