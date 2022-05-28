/*
 * @lc app=leetcode id=46 lang=csharp
 *
 * [46] Permutations
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    private int[] nums;

    private HashSet<int> usedCandidates;

    public IList<IList<int>> Permute(int[] nums) 
    {
        this.nums = nums;
        this.usedCandidates = new HashSet<int>();

        return this.GetPermutation();
    }

    public IList<IList<int>> GetPermutation()
    {
        List<IList<int>> results = new List<IList<int>>();

        foreach (var candidate in this.nums)
        {
            if (this.usedCandidates.Contains(candidate))
            {
                continue;
            }

            this.usedCandidates.Add(candidate);
            foreach (var subPermutation in this.GetPermutation())
            {
                subPermutation.Add(candidate);
                results.Add(subPermutation);
            }
            this.usedCandidates.Remove(candidate);
        }

        if (results.Count == 0)
        {
            results.Add(new List<int>());
        }

        return results;
    }
}
// @lc code=end

