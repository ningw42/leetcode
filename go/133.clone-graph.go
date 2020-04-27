/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// 1. clone nodes and make a mapping from original nodes to copied nodes
// 2. connect edges
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    // 1. BFS to traverse graph
    m := make(map[*Node]*Node)
    var queue, nextQueue []*Node
    queue = []*Node{node}
    for len(queue) > 0 {
        for len(queue) > 0 {
            from := queue[0]
            queue = queue[1:]
            if _, exists := m[from]; exists {
                continue
            }
            m[from] = &Node{Val: from.Val}
            nextQueue = append(nextQueue, from.Neighbors...)
        }
        queue, nextQueue = nextQueue, queue
    }
    
    // connect edges
    for original, copied := range m {
        for _, to := range original.Neighbors {
            copied.Neighbors = append(copied.Neighbors, m[to])
        }
    }
    
    return m[node]
}
