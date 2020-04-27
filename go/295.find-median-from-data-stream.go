type MedianFinder struct {
    max *Heap
    min *Heap
    count int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{
        max: &Heap{t: false},
        min: &Heap{t: true},
    }
}


func (this *MedianFinder) AddNum(num int) {
    this.count++
    if this.count & 1 == 1 {
        // odd, add to max heap
        if this.min.Len() > 0 && num > this.min.container[0] {
            heap.Push(this.min, num)
            heap.Push(this.max, heap.Pop(this.min))
        } else {
            heap.Push(this.max, num)
        }
    } else {
        // even
        if this.max.Len() > 0 && num < this.max.container[0] {
            heap.Push(this.max, num)
            heap.Push(this.min, heap.Pop(this.max))
        } else {
            heap.Push(this.min, num)
        }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.count == 0 {
        return 0
    }
    if this.count & 1 == 1 {
        // odd
        return float64(this.max.container[0])
    } else {
        // even
        return float64(this.max.container[0] + this.min.container[0]) / 2
    }
}

type Heap struct {
    container []int
    t bool // true for minheap, false for maxheap
}
func (h *Heap) Len() int {return len(h.container)}
func (h *Heap) Swap(i, j int) {h.container[i], h.container[j] = h.container[j], h.container[i]}
func (h *Heap) Less(i, j int) bool {
    if h.t {
        return h.container[i] < h.container[j]
    } else {
        return h.container[j] < h.container[i]
    }
}
func (h *Heap) Push(x interface{}) {
    h.container = append(h.container, x.(int))
}
func (h *Heap) Pop() interface{} {
    l := h.Len()
    x := h.container[l-1]
    h.container = h.container[:l-1]
    return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */