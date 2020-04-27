type Position struct {
    X int
    Y int
    Value int
}
func (p Position) GetNeighbour() [][]int {
    return [][]int{
        []int{p.X+1, p.Y},
        []int{p.X-1, p.Y},
        []int{p.X, p.Y+1},
        []int{p.X, p.Y-1},
    }
}

type PriorityQueue []Position
func (pq PriorityQueue) Len() int {
    return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Value > pq[j].Value
}
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Position))
}
func (pq *PriorityQueue) Pop() interface{} {
    length := pq.Len()
    element := (*pq)[length-1]
    *pq = (*pq)[:length-1]
    return element
}
func NewPriorityQueue() *PriorityQueue {
    var data []Position
    queue := PriorityQueue(data)
    return &queue
}

func maximumMinimumPath(A [][]int) int {
    visited := make([][]bool, len(A))
    for i := 0; i < len(A); i++ {
        visited[i] = make([]bool, len(A[0]))
    }
    queue := NewPriorityQueue()
    heap.Init(queue)
    
    queue.Push(Position{
        X: 0,
        Y: 0,
        Value: A[0][0],
    })
    score := A[0][0]
    
    for queue.Len() > 0 {
        pos := heap.Pop(queue).(Position)
        visited[pos.X][pos.Y] = true
        if pos.Value < score {
            score = pos.Value
        }
        if pos.X == len(A)-1 && pos.Y == len(A[0])-1 {
            return score
        }
        neighbours := pos.GetNeighbour()
        for _, neighbour := range neighbours {
            a, b := neighbour[0], neighbour[1]
            if 0 <= a && a < len(A) && 0 <= b && b < len(A[0]) && !visited[a][b] {
                heap.Push(queue, Position{
                    X: a,
                    Y: b,
                    Value: A[a][b],
                })
            }
        }
    }
    return score
}