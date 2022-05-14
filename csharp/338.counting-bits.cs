/*
 * @lc app=leetcode id=338 lang=csharp
 *
 * [338] Counting Bits
 */

// @lc code=start
public class Solution 
{
    public int[] CountBits(int n) 
    {
        int[] dp = new int[n+1];

        dp[0] = 0;
        for (var i = 1; i <= n; i++)
        {
            dp[i] = dp[i>>1] + (i & 1);
        }

        return dp;
    }
}
// @lc code=end

