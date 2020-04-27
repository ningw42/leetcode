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