func intervalIntersection(A [][]int, B [][]int) [][]int {
    return withMinHeap(A, B)
}

func bruteforce(A, B [][]int) [][]int {
    // TODO
    return nil
}

func withMinHeap(A, B [][]int) [][]int {
    // merge A and B
    merged := make([][]int, len(A) + len(B))
    for i, interval := range A {
        merged[i] = interval
    }
    for i, interval := range B {
        merged[len(A)+i] = interval
    }
    // sort by left bound ascending
    sort.Slice(merged, func (i, j int) bool {
        ii, jj := merged[i], merged[j]
        if ii[0] < jj[0] {
            return true
        } else if ii[0] > jj[0] {
            return false
        } else {
            return ii[1] < jj[1]
        }
    })

    pq := newMinIntervalHeap()
    var intersection [][]int
    for _, interval := range merged {
        left := interval[0]
        // remove interval from heap with smaller right bound compared to current interval's left bound
        for pq.Len() > 0 && (*pq)[0][1] < left {
            heap.Pop(pq)
        }
        if pq.Len() > 0 {
            // since there are two slice of disjoint intervals, heap should contain at most one interval which intersects with current interval
            right := min(interval[1], (*pq)[0][1])
            intersection = append(intersection, []int{left, right})
        }
        heap.Push(pq, interval)
    }
    
    return intersection
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

type MinIntervalHeap [][]int
func (h MinIntervalHeap) Len() int {return len(h)}
func (h MinIntervalHeap) Less(i, j int) bool {return h[i][1] < h[j][1]}
func (h MinIntervalHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinIntervalHeap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}
func (h *MinIntervalHeap) Pop() interface{} {
    length := len(*h)
    x := (*h)[length-1]
    (*h) = (*h)[:length-1]
    return x
}
func newMinIntervalHeap() *MinIntervalHeap {
    var container [][]int
    h := MinIntervalHeap(container)
    return &h
}