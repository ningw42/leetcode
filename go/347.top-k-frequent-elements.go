func topKFrequent(nums []int, k int) []int {
    return priorityQueue(nums, k)
}

func priorityQueue(nums []int, k int) []int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }
    pq := NewPQ()
    for num, c := range count {
        heap.Push(pq, Pair{
            num: num,
            count: c,
        })
    }
    ret := make([]int, k)
    for i := 0; i < k; i++ {
        ret[i] = heap.Pop(pq).(Pair).num
    }
    return ret
}

type Pair struct {
    num int
    count int
}
type PQ []Pair
func NewPQ() *PQ {
    var container []Pair
    pq := PQ(container)
    return &pq
}
func (p *PQ) Len() int {return len(*p)}
func (p *PQ) Less(i, j int) bool {return (*p)[i].count > (*p)[j].count}
func (p *PQ) Swap(i, j int) {(*p)[i], (*p)[j] = (*p)[j], (*p)[i]}
func (p *PQ) Push(x interface{}) {(*p) = append((*p), x.(Pair))}
func (p *PQ) Pop() interface{} {
    length := p.Len()
    x := (*p)[length-1]
    (*p) = (*p)[:length-1]
    return x
}

