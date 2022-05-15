/*
 * @lc app=leetcode id=860 lang=csharp
 *
 * [860] Lemonade Change
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    public bool LemonadeChange(int[] bills) 
    {
        int fives = 0;
        int tens = 0;
        for (var i = 0; i < bills.Length; i++)
        {
            switch (bills[i])
            {
                case 5:
                    fives++;
                    break;
                case 10:
                    tens++;
                    fives--;
                    break;
                case 20:
                    if (tens > 0)
                    {
                        tens--;
                        fives--;
                    }
                    else 
                    {
                        fives -= 3;
                    }
                    break;
            }

            if (fives < 0 || tens < 0)
            {
                return false;
            }
        }

        return true;
    }
}
// @lc code=end

