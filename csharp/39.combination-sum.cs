/*
 * @lc app=leetcode id=39 lang=csharp
 *
 * [39] Combination Sum
 */

// @lc code=start

using System;
using System.Collections.Generic;

public class Solution 
{
    private int[] candidates;

    private IList<IList<int>> results;

    public IList<IList<int>> CombinationSum(int[] candidates, int target) 
    {
        this.results = new List<IList<int>>();

        // sort candidates ascending
        this.candidates = candidates;
        Array.Sort(this.candidates);

        this.Combination(0, target, new Stack<int>());

        return this.results;
    }

    public void Combination(int start, int target, Stack<int> result)
    {
        // Console.WriteLine($"start: {start}, target: {target}");
        for (int i = start; i < this.candidates.Length; i++)
        {
            var candidate = this.candidates[i];
            var subTarget = target - candidate;
            if (subTarget < 0)
            {
                return;
            }
            else if (subTarget == 0)
            {
                result.Push(candidate);
                // each result in reverse order, doesn't matter in this case.
                this.results.Add(result.ToList()); 
                result.Pop();
            }
            else 
            {
                result.Push(candidate);
                this.Combination(i, subTarget, result);
                result.Pop();
            }
        }
    }
}
// @lc code=end

