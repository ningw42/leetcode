/*
 * @lc app=leetcode id=684 lang=csharp
 *
 * [684] Redundant Connection
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    public int[] FindRedundantConnection(int[][] edges) 
    {
        return this.Kruskal(edges);
    }

    #region Kruskal

    // use Kruskal's MST algorithm to find the edge. 
    // Prim should work too, but less intuitive.
    // edges are weighted by its index.
    public int[] Kruskal(int[][] edges)
    {
        int n = edges.Length;

        // disjoint set for cycle detection
        DisjointSet set = new DisjointSet(n + 1);

        // we need N-1 edge to connect all vertexes
        int edgeCount = 0;
        int nextEdge = 0;
        while (edgeCount < n - 1)
        {
            int from = edges[nextEdge][0];
            int to = edges[nextEdge][1];

            // cycle detection after adding current edge
            int rootFrom = set.Find(from);
            int rootTo = set.Find(to);
            if (rootFrom != rootTo)
            {
                set.Union(from, to);
            }
            else 
            {
                // it is guaranteed to have exact one answer, so we could do this.
                return edges[nextEdge];
            }

            nextEdge++;
        }

        // shouldn't reach here.
        return null;
    }

    // DisjointSet without any optimization
    public class DisjointSet
    {
        private int[] set; 

        public DisjointSet(int size)
        {
            this.set = new int[size];
            for (int i = 0; i < size; i++)
            {
                this.set[i] = i;
            }
        }

        public void Union(int x, int y)
        {
            int rootX = this.Find(x);
            int rootY = this.Find(y);

            this.set[rootX] = rootY;
        }

        public int Find(int x)
        {
            while (this.set[x] != x)
            {
                x = this.set[x];
            }

            return x;
        }
    }

    #endregion Kruskal
}
// @lc code=end

