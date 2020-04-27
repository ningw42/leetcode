/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// tuple containing node and its vertical index
type Tuple struct {
    node *TreeNode
    x int
    y int
}

// bfs to preserve the order
func verticalTraversal(root *TreeNode) [][]int {
    var queue, next []*Tuple
    queue = append(queue, &Tuple{
        node: root,
        x: 0,
        y: 0,
    })
    // map of vertical indexed nodes' value
    m := make(map[int][]*Tuple)
    // left and right bound of vertical index
    var left, right int
    for len(queue) > 0 {
        for len(queue) > 0 {
            // pop from queue
            tuple := queue[0]
            queue = queue[1:]
            
            // ignore nil node
            if tuple.node == nil {
                continue
            }
            
            // index node by its vertical index
            m[tuple.x] = append(m[tuple.x], tuple)
            // keep track of the left and right bound
            left = min(left, tuple.x)
            right = max(right, tuple.x)
            
            // add children of current node to queue
            next = append(next, &Tuple{
                node: tuple.node.Left,
                x: tuple.x - 1,
                y: tuple.y - 1,
            })
            next = append(next, &Tuple{
                node: tuple.node.Right,
                x: tuple.x + 1,
                y: tuple.y - 1,
            })
        }
        queue, next = next, queue
    }
  
    // relaxing result
    // root == nil
    if len(m) == 0 {
        return nil
    }
    
    var results [][]int
    for i := left; i <= right; i++ {
        result := m[i]
        sort.Slice(result, func(i, j int) bool {
            if result[i].y > result[j].y {
                return true
            } else if result[i].y < result[j].y {
                return false
            } else {
                return result[i].node.Val < result[j].node.Val
            }
        })
        var values []int
        for _, tuple := range result {
            values = append(values, tuple.node.Val)
        }
        results = append(results, values)
    }
    
    return results
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
