type MinHeap []int
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    length := h.Len()
    x := (*h)[length-1]
    *h = (*h)[:length-1]
    return x
}

type KthLargest struct {
    pq *MinHeap
    k int
}

func Constructor(k int, nums []int) KthLargest {
    pq := MinHeap(nil)
    ret := KthLargest{
        k: k,
        pq: &pq,
    }
    for _, num := range nums {
        ret.Add(num)
    }
    return ret
}

func (this *KthLargest) Add(val int) int {
    if this.pq.Len() < this.k {
        heap.Push(this.pq, val)
    } else {
        if val > (*this.pq)[0] {
            heap.Push(this.pq, val)
            heap.Pop(this.pq)
        }
    }
    return (*this.pq)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */