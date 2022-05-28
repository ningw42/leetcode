/*
 * @lc app=leetcode id=138 lang=csharp
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/*
// Definition for a Node.
public class Node {
    public int val;
    public Node next;
    public Node random;
    
    public Node(int _val) {
        val = _val;
        next = null;
        random = null;
    }
}
*/

using System.Collections.Generic;

public class Solution 
{
    public Node CopyRandomList(Node head) 
    {
        Dictionary<Node, Node> mapping = new Dictionary<Node, Node>();

        Node cur, dummy, prev;
        dummy = new Node(int.MaxValue);
        prev = dummy;
        cur = head;

        // deep copy list
        while (cur != null)
        {
            prev.next = new Node(cur.val);
            mapping[cur] = prev.next;
            prev = prev.next;
            cur = cur.next;
        }

        // link random
        cur = head;
        while (cur != null)
        {
            if (cur.random == null)
            {
                mapping[cur].random = null;
            }
            else 
            {
                mapping[cur].random = mapping[cur.random];
            }
            cur = cur.next;
        }

        return dummy.next;
    }
}
// @lc code=end

