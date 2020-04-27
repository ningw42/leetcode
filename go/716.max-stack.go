type Item struct {
    stackIndex int
    heapIndex int
    value int
}

type MaxHeap []*Item
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Swap(i, j int) {
    h[i], h[j], h[i].heapIndex, h[j].heapIndex = h[j], h[i], h[j].heapIndex, h[i].heapIndex
}
func (h MaxHeap) Less(i, j int) bool {
    if h[i].value > h[j].value {
        return true
    } else if h[i].value < h[j].value {
        return false
    } else {
        return h[i].stackIndex > h[j].stackIndex
    }
}
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(*Item))
}
func (h *MaxHeap) Pop() interface{} {
    length := h.Len()
    x := (*h)[length-1]
    *h = (*h)[:length-1]
    return x
}
func newMaxHeap() *MaxHeap {
    var container []*Item
    h := MaxHeap(container)
    return &h
}

type MaxStack struct {
    h *MaxHeap
    s []*Item
}

/** initialize your data structure here. */
func Constructor() MaxStack {
    return MaxStack{
        h: newMaxHeap(),
    }
}


func (this *MaxStack) Push(x int) {
    item := &Item{
        stackIndex: len(this.s),
        heapIndex: this.h.Len(),
        value: x,
    }
    this.s = append(this.s, item)
    heap.Push(this.h, item)
}


func (this *MaxStack) Pop() int {
    elem := this.s[len(this.s)-1]
    this.s = this.s[:len(this.s)-1]
    heap.Remove(this.h, elem.heapIndex)
    return elem.value
}


func (this *MaxStack) Top() int {
    return this.s[len(this.s)-1].value
}


func (this *MaxStack) PeekMax() int {
    return (*(this.h))[0].value
}


func (this *MaxStack) PopMax() int {
    elem := heap.Pop(this.h).(*Item)
    this.s = append(this.s[:elem.stackIndex], this.s[elem.stackIndex+1:]...)
    for i := elem.stackIndex; i < len(this.s); i++ {
        this.s[i].stackIndex--
    }
    return elem.value
}


/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */