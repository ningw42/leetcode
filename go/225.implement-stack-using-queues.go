// a standard queue implemented by slice
type Queue []int
func (q *Queue) Push(e int) {
    *q = append(*q, e)
}
func (q *Queue) Pop() int {
    e := (*q)[0]
    *q = (*q)[1:]
    return e
}
func (q Queue) Peek() int {
    return q[0]
}
func (q Queue) Size() int {
    return len(q)
}
func (q Queue) Empty() bool {
    return q.Size() == 0
}

type MyStack struct {
    q *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    queue := Queue(make([]int, 0))
    return MyStack{q: &queue}
}

// O(N)
/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    // push element to the end of queue
    this.q.Push(x)
    // move previous elements
    for i := 0; i < this.q.Size() - 1; i++ {
        this.q.Push(this.q.Pop())
    }
}

// O(1)
/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    return this.q.Pop()
}


// O(1)
/** Get the top element. */
func (this *MyStack) Top() int {
    return this.q.Peek()
}


// O(1)
/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.q.Empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */