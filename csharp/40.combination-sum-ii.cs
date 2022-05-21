/*
 * @lc app=leetcode id=40 lang=csharp
 *
 * [40] Combination Sum II
 */

// @lc code=start

using System;
using System.Collections.Generic;
using System.Linq;

public class Solution 
{
    private int[] uniqueCandidates;

    private Dictionary<int, int> candidateMap;

    private Dictionary<string, IList<int>> results;

    public IList<IList<int>> CombinationSum2(int[] candidates, int target) 
    {
        this.results = new Dictionary<string, IList<int>>();
        this.InitializeCandidates(candidates);

        this.Combination(0, target, new Stack<int>());

        return this.results.Values.ToList();
    }

    // start: dedup
    public void Combination(int start, int target, Stack<int> result)
    {
        for (var i = start; i < this.uniqueCandidates.Length; i++)
        {
            var candidate = this.uniqueCandidates[i];
            if (this.candidateMap[candidate] > 0)
            {
                var subTarget = target - candidate;
                if (subTarget < 0)
                {
                    break;
                }
                else if (subTarget == 0)
                {
                    result.Push(candidate);
                    this.results[string.Join(",", result)] = result.ToArray();
                    result.Pop();
                }
                else 
                {
                    result.Push(candidate);
                    this.candidateMap[candidate]--;
                    this.Combination(i, subTarget, result);
                    this.candidateMap[candidate]++;
                    result.Pop();
                }
            }
        }
    }

    public void InitializeCandidates(int[] candidates)
    {
        this.candidateMap = new Dictionary<int, int>();

        foreach (var candidate in candidates)
        {
            if (this.candidateMap.TryGetValue(candidate, out int count))
            {
                this.candidateMap[candidate] = count + 1;
            }
            else 
            {
                this.candidateMap[candidate] = 1;
            }
        }

        this.uniqueCandidates = this.candidateMap.Keys.ToArray();
        Array.Sort(this.uniqueCandidates);
    }
}
// @lc code=end

