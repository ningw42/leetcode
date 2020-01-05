type Graph struct {
    AdjacencyList map[int][]int
    InDegree []int
}
func (g *Graph) AddEdge(from, to int) {
    g.InDegree[to]++
    g.AdjacencyList[from] = append(g.AdjacencyList[from], to)
}
func (g *Graph) TopologicalSort() []int {
    var queue []int
    for i, indegree := range g.InDegree {
        if indegree == 0 {
            queue = append(queue, i)
        }
    }
   
    var order []int
    for len(queue) > 0 {
        var newQueue []int
        for _, from := range queue {
            order = append(order, from)
            for _, to := range g.AdjacencyList[from] {
                g.InDegree[to]--
                if g.InDegree[to] == 0 {
                    newQueue = append(newQueue, to)
                }
            }
        }
        queue = newQueue
    }
    
    return order
}
func NewGraph(n int) *Graph {
    return &Graph{
        AdjacencyList: make(map[int][]int),
        InDegree: make([]int, n),
    }
}

// construct a directed map, then apply topological sort
// be aware of the invalid word sequence
func alienOrder(words []string) string {
    chars, mapping := getMapping(words)
    // one char
    if len(chars) == 1 {
        return string(chars)
    }
    graph := NewGraph(len(mapping))
    
    // add edge according to lexicographically order
    // if lexigraphical order is not valid
    if valid := addEdges(graph, words, mapping); !valid {
        return ""
    }
   
    order := graph.TopologicalSort()
    
    // cycle exists, not valid
    if len(order) < len(mapping) {
        return ""
    }
    
    // concat result in the topological order
    builder := strings.Builder{}
    for _, node := range order {
        builder.WriteByte(chars[node])
    }
    return builder.String()
}

func addEdges(g *Graph, words []string, mapping map[byte]int) bool {
    // first byte indicates the lexicographically order
    var order []byte
    m := make(map[byte][]string)
    var previousFirstChar, firstChar byte
    for _, word := range words {
        if len(word) == 0 {
            continue
        }
        firstChar = word[0]
        if _, exists := m[firstChar]; !exists {
            order = append(order, firstChar)
        }
        // make sure prefix is the first, see note #2
        if rest := word[1:]; rest == "" && len(m[firstChar]) > 0 {
            return false
        } else {
            // make sure the word with the same 'firstChar' is continuous
            if len(m[firstChar]) > 0 && previousFirstChar != firstChar {
                return false
            } else {
                m[firstChar] = append(m[firstChar], word[1:])
            }
        }
        previousFirstChar = firstChar
    }
    for i := 0; i < len(order); i++ {
        for j := i+1; j < len(order); j++ {
            g.AddEdge(mapping[order[i]], mapping[order[j]])
        }
    }
    for _, ws := range m {
        if len(ws) > 1 {
            if valid := addEdges(g, ws, mapping); !valid {
                return false
            }
        }
    }
    return true
}

func getMapping(words []string) ([]byte, map[byte]int) {
    idx := 0
    mapping := make(map[byte]int)
    var chars []byte
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if _, exists := mapping[word[i]]; !exists {
                chars = append(chars, word[i])
                mapping[word[i]] = idx
                idx++
            }
        }
    }
    return chars, mapping
}
