/*
 * @lc app=leetcode id=47 lang=csharp
 *
 * [47] Permutations II
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    private SortedDictionary<int, int> candidates;

    private Dictionary<int, int> usedCandidates;

    public IList<IList<int>> PermuteUnique(int[] nums) 
    {
        this.candidates = new SortedDictionary<int, int>();
        this.usedCandidates = new Dictionary<int, int>();

        foreach (var num in nums)
        {
            if (this.candidates.TryGetValue(num, out int count))
            {
                this.candidates[num] = count + 1;
            }
            else 
            {
                this.candidates[num] = 1;
            }
            this.usedCandidates[num] = 0;
        }

        return this.GetUniquePermutation();
    }

    public IList<IList<int>> GetUniquePermutation()
    {
        List<IList<int>> results = new List<IList<int>>();
        // duplicate candidate produces exact same sub-permutation.
        HashSet<int> localUsedCandiates = new HashSet<int>();

        foreach (var candidate in candidates)
        {
            if (localUsedCandiates.Contains(candidate.Key) || this.usedCandidates[candidate.Key] == candidate.Value)
            {
                continue;
            }

            localUsedCandiates.Add(candidate.Key);
            this.usedCandidates[candidate.Key]++;
            foreach (var subPermutation in this.GetUniquePermutation())
            {
                subPermutation.Add(candidate.Key);
                results.Add(subPermutation);
            }
            this.usedCandidates[candidate.Key]--;
        }

        if (results.Count == 0)
        {
            results.Add(new List<int>());
        }

        return results;
    }
}
// @lc code=end

