/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    // mapping from orginal nodes to copied nodes
    m := make(map[*Node]*Node)
    // copy singly linked list
    newHead := deepCopy(head, m)
    // use the mapping to fill Random pointer
    for head != nil {
        m[head].Random = m[head.Random]
        head = head.Next
    }
    return newHead
}

// deep copy a singly linked list(ignore the Random pointer)
// and keep track of a mapping from original node to copied node
func deepCopy(head *Node, m map[*Node]*Node) *Node {
    if head == nil {
        return nil
    } else {
        next := deepCopy(head.Next, m)
        node := &Node{
            Val: head.Val,
            Next: next,
        }
        m[head] = node
        return node
    }
}