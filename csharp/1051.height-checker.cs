/*
 * @lc app=leetcode id=1051 lang=csharp
 *
 * [1051] Height Checker
 */

// @lc code=start

using System;

public class Solution 
{
    public int HeightChecker(int[] heights) 
    {
        // return BruteForce(heights);
        return CountingSort(heights);
    }

    public int CountingSort(int[] heights)
    {
        // calculate height distribution
        // height -> count
        int[] heightCount = new int[101];
        foreach (var h in heights)
        {
            heightCount[h]++;
        }

        int count = 0;
        int expectedHeight = 1;

        foreach (var h in heights)
        {
            // find next expect height in height distribution
            while (heightCount[expectedHeight] == 0)
            {
                expectedHeight++;
            }

            // count if actual height doesn't match the expected height
            if (h != expectedHeight)
            {
                count++;
            }

            // remove compared expected height from height distribution
            heightCount[expectedHeight]--;
        }

        return count;
    }

    // sort and compare
    public int BruteForce(int[] heights)
    {
        int[] expected = new int[heights.Length];
        for (var i = 0; i < heights.Length; i++)
        {
            expected[i] = heights[i];
        }
        Array.Sort(expected);

        int count = 0;
        for (var i = 0; i < heights.Length; i++)
        {
            if (heights[i] != expected[i])
            {
                count++;
            }
        }

        return count;
    }
}
// @lc code=end

