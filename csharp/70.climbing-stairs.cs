/*
 * @lc app=leetcode id=70 lang=csharp
 *
 * [70] Climbing Stairs
 */

// @lc code=start
public class Solution 
{
    public int ClimbStairs(int n) 
    {
        int[] dp = new int[n+1];

        dp[0] = 1;
        dp[1] = 1;
        for (var i = 2; i <= n; i++)
        {
            dp[i] = dp[i-1] + dp[i-2];
        }

        return dp[n];
    }
}
// @lc code=end

