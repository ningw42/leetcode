/*
 * @lc app=leetcode id=210 lang=csharp
 *
 * [210] Course Schedule II
 */

// @lc code=start

using System;
using System.Collections.Generic;

public class Solution 
{
    private Dictionary<int, HashSet<int>> incoming;

    private Dictionary<int, HashSet<int>> outgoing;

    // Topological Sort
    public int[] FindOrder(int numCourses, int[][] prerequisites) 
    {
        this.BuildGraph(numCourses, prerequisites);

        Queue<int> nodeWithoutIncomingEdge = new Queue<int>();
        List<int> topologicalOrder = new List<int>();

        // 1. find all nodes without incoming edge
        for (var node = 0; node < numCourses; node++)
        {
            if (!this.incoming.TryGetValue(node, out _))
            {
                nodeWithoutIncomingEdge.Enqueue(node);
            }
        }

        // until there is no node without incoming edge
        while (nodeWithoutIncomingEdge.Count > 0)
        {
            // get a node without incoming edge and put it into result
            var from = nodeWithoutIncomingEdge.Dequeue();
            topologicalOrder.Add(from);

            // for all edges from 'from' to 'to', remove them
            // if there is a new node without incoming edges, add it to queue.
            if (this.outgoing.TryGetValue(from, out HashSet<int> tos))
            {
                foreach (var to in tos)
                {
                    this.incoming[to].Remove(from);
                    if (this.incoming[to].Count == 0)
                    {
                        nodeWithoutIncomingEdge.Enqueue(to);
                        this.incoming.Remove(to);
                    }
                }
                this.outgoing.Remove(from);
            }
        }

        if (this.incoming.Count == 0)
        {
            return topologicalOrder.ToArray();
        }
        else 
        {
            return Array.Empty<int>();
        }
    }

    public void BuildGraph(int numCourses, int[][] prerequisites)
    {
        this.incoming = new Dictionary<int, HashSet<int>>();
        this.outgoing = new Dictionary<int, HashSet<int>>();

        foreach (var edge in prerequisites)
        {
            var from = edge[1];
            var to = edge[0];

            if (this.incoming.TryGetValue(to, out HashSet<int> froms))
            {
                froms.Add(from);
            }
            else
            {
                this.incoming[to] = new HashSet<int>() { from };
            }

            if (this.outgoing.TryGetValue(from, out HashSet<int> tos))
            {
                tos.Add(to);
            }
            else 
            {
                this.outgoing[from] = new HashSet<int>() { to };
            }
        }
    }
}
// @lc code=end

