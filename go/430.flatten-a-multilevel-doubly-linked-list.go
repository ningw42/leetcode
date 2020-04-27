/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(head *Node) *Node {
    do(head)
    return head
}

func do(head *Node) (*Node, *Node) {
    var prev *Node
    cursor := head
    for cursor != nil {
        if cursor.Child != nil {
            h, t := do(cursor.Child)
            cursor.Child = nil
            if cursor.Next == nil {
                cursor.Next, h.Prev, t.Next, cursor = h, cursor, nil, nil
            } else {
                cursor.Next, h.Prev, t.Next, cursor.Next.Prev, cursor = h, cursor, cursor.Next, t, cursor.Next
            }
            prev = t
        } else {
            cursor, prev = cursor.Next, cursor
        }
    }
    return head, prev
}
