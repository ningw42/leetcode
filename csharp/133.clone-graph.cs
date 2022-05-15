/*
 * @lc app=leetcode id=133 lang=csharp
 *
 * [133] Clone Graph
 */

// @lc code=start
/*
// Definition for a Node.
public class Node {
    public int val;
    public IList<Node> neighbors;

    public Node() {
        val = 0;
        neighbors = new List<Node>();
    }

    public Node(int _val) {
        val = _val;
        neighbors = new List<Node>();
    }

    public Node(int _val, List<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

using System.Collections.Generic;

public class Solution 
{
    public Dictionary<Node, Node> map = new Dictionary<Node, Node>();

    public Node CloneGraph(Node node) 
    {
        if (node == null)
        {
            return null;
        }

        Node newNode = null;

        if (!this.map.TryGetValue(node, out newNode))
        {
            newNode = new Node(node.val);
            this.map[node] = newNode;
            foreach (var neighbor in node.neighbors)
            {
                newNode.neighbors.Add(this.CloneGraph(neighbor));
            }
        }

        return newNode;
    }
}
// @lc code=end

