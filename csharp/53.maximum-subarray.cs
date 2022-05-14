/*
 * @lc app=leetcode id=53 lang=csharp
 *
 * [53] Maximum Subarray
 */

// @lc code=start

using System;

public class Solution 
{
    // dp[i] means max sum you can get with subarray ends at i-th element.
    // result should be max(dp)
    public int MaxSubArray(int[] nums) 
    {
        int[] dp = new int[nums.Length];

        dp[0] = nums[0];
        int max = nums[0];
        for (var i = 1; i < nums.Length; i++)
        {
            dp[i] = dp[i-1] > 0 ? dp[i-1] + nums[i] : nums[i];
            max = Math.Max(max, dp[i]);
        }

        return max;
    }
}
// @lc code=end

