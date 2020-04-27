type PQ [][]int
func (pq PQ) Len() int {return len(pq)}
// ascending order of right bound
func (pq PQ) Less(i, j int) bool {return pq[i][1] < pq[j][1]} 
func (pq PQ) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}
func (pq *PQ) Push(x interface{}) {
    *pq = append(*pq, x.([]int))
}
func (pq *PQ) Pop() interface{} {
    length := pq.Len()
    x := (*pq)[length-1]
    (*pq) = (*pq)[:length-1]
    return x
}
func NewPQ() *PQ {
    var container [][]int
    pq := PQ(container)
    return &pq
}

func merge(intervals [][]int) [][]int {
    // edge case
    if len(intervals) == 0 {
        return nil
    }
    
    // sort by left bound ascending
    sort.Slice(intervals, func (i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // priority queue of intervals ordered by right bound
    // priority queue holds all intervals currently not merged
    var ret [][]int
    var left, right int
    pq := NewPQ()
    // loop over intervals by ascending order of left bound
    for _, interval := range intervals {
        if pq.Len() == 0 {
            // nothing in priority queue
            // update left bound
            left = interval[0]
        } else {
            // pop from priority queue for all intervals with smaller right bound than
            // current interval's left bound
            for pq.Len() > 0 && (*pq)[0][1] < interval[0] {
                right = heap.Pop(pq).([]int)[1]
            }
            
            // all previous intervals have smaller right bound 
            // than current interval's left bound
            // merged interval should be generated
            if pq.Len() == 0 {
                ret = append(ret, []int{left, right})
                // update left bound for next merged interval
                left = interval[0]
            }
        }
        // put current interval into pq
        heap.Push(pq, interval)
    }
    // last merged interval's right bound
    for pq.Len() > 0 {
        right = heap.Pop(pq).([]int)[1]
    }
    ret = append(ret, []int{left, right})
    
    return ret
}