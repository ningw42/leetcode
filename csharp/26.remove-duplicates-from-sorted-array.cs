/*
 * @lc app=leetcode id=26 lang=csharp
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
public class Solution 
{
    public int RemoveDuplicates(int[] nums) 
    {
        if (nums.Length <= 1)
        {
            return nums.Length;
        }

        int newIndex = 1;

        for (var i = 1; i < nums.Length; i++)
        {
            if (nums[i-1] != nums[i])
            {
                nums[newIndex++] = nums[i];
            }
        }

        return newIndex;
    }
}
// @lc code=end

