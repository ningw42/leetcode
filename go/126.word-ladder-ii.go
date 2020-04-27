// djikstra single sourced shortest path
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    m := makeMap(wordList)
    if _, exists := m[endWord]; !exists {
        return nil
    }
    m[beginWord] = len(wordList)
    wordList = append(wordList, beginWord)
    g := makeGraph(wordList)
    begin := m[beginWord]
    end := m[endWord]
    visited := make([]bool, len(wordList))
    dist := make([]int, len(wordList))
    prev := make([][]int, len(wordList))
    for i := 0; i < len(wordList); i++ {
        dist[i] = math.MaxInt64
    }
    dist[begin] = 0
    pq := NewPriorityQueue()
    heap.Push(pq, Item{Node: begin, Dist: 0})
    for pq.Len() > 0 {
        item := heap.Pop(pq).(Item)
        from := item.Node
        visited[from] = true
        for to, edgeExists := range g[from] {
            if edgeExists {
                if newDist := item.Dist + 1; newDist < dist[to] {
                    dist[to] = newDist
                    prev[to] = append(prev[to], from)
                    heap.Push(pq, Item{Node: to, Dist: newDist})
                } else if newDist == dist[to] {
                    prev[to] = append(prev[to], from)
                }
            }
        }
    }
    if dist[end] == math.MaxInt64 {
        return nil
    } else {
        return reconstructPath(begin, end, prev, wordList)
    }
}

type Item struct {
    Node int
    Dist int
}

type PriorityQueue []Item
func NewPriorityQueue() *PriorityQueue {
    var items []Item
    pq := PriorityQueue(items)
    return &pq
}
func (pq *PriorityQueue) Len() int {
    return len(*pq)
}
func (pq *PriorityQueue) Swap(i, j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *PriorityQueue) Less(i, j int) bool {
    return (*pq)[i].Dist < (*pq)[j].Dist
}
func (pq *PriorityQueue) Push(item interface{}) {
    (*pq) = append((*pq), item.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
    length := pq.Len()
    item := (*pq)[length-1]
    (*pq) = (*pq)[:length-1]
    return item
}

func reconstructPath(from, to int, prevs [][]int, list []string) [][]string {
    if from == to {
        return [][]string{[]string{list[from]}}
    } else {
        var pathes [][]string
        for _, prev := range prevs[to] {
            for _, path := range reconstructPath(from, prev, prevs, list) {
                pathes = append(pathes, append(path, list[to]))
            }
        }
        return pathes
    }
}

func makeGraph(list []string) [][]bool {
    g := make([][]bool, len(list))
    for i := 0; i < len(list); i++ {
        g[i] = make([]bool, len(list))
    }
    for i := 0; i < len(list); i++ {
        for j := i + 1; j < len(list); j++ {
            if stringDiff(list[i], list[j]) == 1 {
                g[i][j] = true
                g[j][i] = true
            }
        }
    }
    return g
}

func stringDiff(a, b string) int {
    var diffBytes int
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            diffBytes++
        }
    }
    return diffBytes
}

func makeMap(list []string) map[string]int {
    m := make(map[string]int)
    for i, s := range list {
        m[s] = i
    }
    return m
}