type MyCircularQueue struct {
    capacity int
    length int
    head int
    tail int
    container []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        capacity: k,
        container: make([]int, k),
    }
}

func (this *MyCircularQueue) next(cursor int) int {
    return (cursor + 1) % this.capacity
}
func (this *MyCircularQueue) previous(cursor int) int {
    prev := cursor - 1
    if prev < 0 {
        return (prev + this.capacity) % this.capacity
    } else {
        return prev
    }
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    this.container[this.tail] = value
    this.length++
    this.tail = this.next(this.tail)
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.length--
    this.head = this.next(this.head)
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.container[this.head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.container[this.previous(this.tail)]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return this.length == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */