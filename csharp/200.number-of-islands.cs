/*
 * @lc app=leetcode id=200 lang=csharp
 *
 * [200] Number of Islands
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    private char[][] grid;

    public int NumIslands(char[][] grid) 
    {
        int count = 0;
        this.grid = grid;

        for (int i = 0; i < this.grid.Length; i++)
        {
            for (int j = 0; j < this.grid[i].Length; j++)
            {
                if (this.grid[i][j] == '1')
                {
                    count++;
                    this.DFS(i, j);
                    // this.BFS(i, j);
                }
            }
        }

        return count;
    }

    public static (int, int)[] offsets = new (int, int)[] 
    {
        (1, 0),
        (-1, 0),
        (0, 1),
        (0, -1),
    };

    public void BFS(int x, int y)
    {
        Queue<(int, int)> q = new Queue<(int, int)>();
        q.Enqueue((x, y));

        while (q.Count > 0)
        {
            var (i, j) = q.Dequeue();
            this.grid[i][j] = '0';
            foreach (var (di, dj) in offsets)
            {
                int ii = i + di;
                int jj = j + dj;
                if (
                    0 <= ii && ii < this.grid.Length && 0 <= jj && jj < this.grid[0].Length
                    && this.grid[ii][jj] == '1')
                {
                    q.Enqueue((ii, jj));
                }
            }
        }
    }

    public void DFS(int x, int y)
    {
        // out of bound
        if (!(0 <= x && x < this.grid.Length && 0 <= y && y < this.grid[0].Length))
        {
            return;
        }

        // water or traversed
        if (this.grid[x][y] == '0') 
        {
            return;
        }

        this.grid[x][y] = '0';
        this.DFS(x + 1, y);
        this.DFS(x - 1, y);
        this.DFS(x, y + 1);
        this.DFS(x, y - 1);
    }
}
// @lc code=end

