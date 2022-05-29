/*
 * @lc app=leetcode id=560 lang=csharp
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    public int SubarraySum(int[] nums, int k) 
    {
        // return BruteForce(nums, k);
        return PrefixSum(nums, k);
    }

    public int PrefixSum(int[] nums, int k)
    {
        int count = 0;
        int[] prefixSums = new int[nums.Length];

        // prefixSums[i] = sum from nums[0] to nums[i] inclusive.
        prefixSums[0] = nums[0];
        for (int i = 1; i < nums.Length; i++)
        {
            prefixSums[i] = prefixSums[i - 1] + nums[i];
        }

        // (prefix_sum, number_of_occurrence(before current element))
        Dictionary<int, int> prefixSumCounts = new Dictionary<int, int>();
        foreach (int sum in prefixSums)
        {
            if (sum == k)
            {
                count++;
            }

            // add number of occurrences of the sub target prior to current element
            int subTarget = sum - k;
            if (prefixSumCounts.TryGetValue(subTarget, out int prefixSumCount))
            {
                count += prefixSumCount;
            }

            // put current prefix sum into map
            if (prefixSumCounts.TryGetValue(sum, out int c))
            {
                prefixSumCounts[sum] = c + 1;
            }
            else 
            {
                prefixSumCounts[sum] = 1;
            }
        }

        return count;
    }

    // may pass if you have good luck.
    public int BruteForce(int[] nums, int k)
    {
        int count = 0;
        int sum = 0;
        int[] sums = new int[nums.Length];

        for (int i = 0; i < nums.Length; i++)
        {
            for (int j = 0; j + i < nums.Length; j++)
            {
                sum = sums[j] + nums[j + i];
                sums[j] = sum;
                if (sum == k)
                {
                    count++;
                }
            }
        }

        return count;
    }
}
// @lc code=end

