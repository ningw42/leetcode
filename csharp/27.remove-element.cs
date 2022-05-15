/*
 * @lc app=leetcode id=27 lang=csharp
 *
 * [27] Remove Element
 */

// @lc code=start
public class Solution 
{
    public int RemoveElement(int[] nums, int val) 
    {
        int left = 0;
        int right = nums.Length - 1;

        while (left <= right)
        {
            if (nums[left] == val)
            {
                if (nums[right] == val)
                {
                    right--;
                }
                else 
                {
                    nums[left] = nums[right];
                    left++;
                    right--;
                }
            }
            else 
            {
                left++;
            }
        }

        return right + 1;
    }
}
// @lc code=end

