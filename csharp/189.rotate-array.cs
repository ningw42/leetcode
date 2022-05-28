/*
 * @lc app=leetcode id=189 lang=csharp
 *
 * [189] Rotate Array
 */

// @lc code=start
public class Solution 
{
    // Space: O(n), Time: O(n)
    // for O(1) space, simply rotate right by 1 for n times.
    public void Rotate(int[] nums, int k) 
    {
        k = k % nums.Length;

        // keep stash in memory
        int[] stash = new int[k];
        for (int i = 0; i < k; i++)
        {
            stash[i] = nums[nums.Length - 1 - i];
        }

        int j = 0;
        for (int i = nums.Length - 1; i >= 0; i--)
        {
            j = i - k;
            if (j >= 0)
            {
                nums[i] = nums[j];
            }
            else
            {
                nums[i] = stash[-j - 1];
            }
        }
    }
}
// @lc code=end

