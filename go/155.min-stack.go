/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (34.76%)
 * Total Accepted:    251.3K
 * Total Submissions: 722.9K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * 
 * 
 * pop() -- Removes the element on top of the stack.
 * 
 * 
 * top() -- Get the top element.
 * 
 * 
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 * 
 * 
 */
 type MinStack struct {
    s []int
    m int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        s:[]int{},
        m:math.MaxInt32,
    }
}


func (this *MinStack) Push(x int)  {
    if x < this.m {
        this.m = x
    }
    this.s = append(this.s, x)
}


func (this *MinStack) Pop()  {
    fmt.Println(this.s)
    if this.s[len(this.s)-1] == this.m {
        this.m = min(this.s[0:len(this.s)-1])
    }
    this.s = this.s[0:len(this.s)-1]
}


func (this *MinStack) Top() int {
    return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
    return this.m
}

func min(s []int) int {
    if len(s) == 0 {
        return math.MaxInt32
	}
	// perform sort on copied slice, prevent from modification to the underlying array
    c := make([]int, len(s))
    copy(c, s)
    sort.Ints(c)
    return c[0]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
