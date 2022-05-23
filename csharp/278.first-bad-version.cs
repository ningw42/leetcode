/*
 * @lc app=leetcode id=278 lang=csharp
 *
 * [278] First Bad Version
 */

// @lc code=start
/* The isBadVersion API is defined in the parent class VersionControl.
      bool IsBadVersion(int version); */

using System;

public class Solution : VersionControl 
{
    // int overflow
    public int FirstBadVersion(int n) 
    {
        uint left = 1;
        uint right = (uint)n;

        while (left < right)
        {
            int mid = (int)((left + right) / 2);
            if (this.IsBadVersion(mid))
            {
                // mid is a bad version
                right = (uint)mid;
            }
            else 
            {
                // mid is not a bad version
                left = (uint)(mid + 1);
            }
        }

        return (int)right;
    }
}
// @lc code=end

