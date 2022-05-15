/*
 * @lc app=leetcode id=542 lang=csharp
 *
 * [542] 01 Matrix
 */

// @lc code=start

using System;
using System.Collections.Generic;

public class Solution 
{
    public static (int, int)[] diff = new (int, int)[]
    {
        (-1, 0),
        (1, 0),
        (0, 1),
        (0, -1),
    };

    // BFS start with zero-cell
    public int[][] UpdateMatrix(int[][] mat) 
    {
        // BFS queue
        Queue<(int, int)> q = new Queue<(int, int)>();

        // Mark and enqueue
        for (var i = 0; i < mat.Length; i++)
        {
            for (var j = 0; j < mat[i].Length; j++)
            {
                if (mat[i][j] == 1)
                {
                    mat[i][j] = -1;
                }
                else 
                {
                    q.Enqueue((i, j));
                }
            }
        }

        // BFS
        while (q.Count > 0)
        {
            var (i, j) = q.Dequeue();
            foreach (var (di, dj) in diff)
            {
                var ii = i + di;
                var jj = j + dj;

                // handle one-cells
                if (0 <= ii && ii < mat.Length && 0 <= jj && jj < mat[0].Length)
                {
                    if (mat[ii][jj] == -1)
                    {
                        mat[ii][jj] = mat[i][j] + 1;
                        q.Enqueue((ii, jj));
                    }
                }
            }
        }

        return mat;
    }
}
// @lc code=end

