/*
 * @lc app=leetcode id=20 lang=csharp
 *
 * [20] Valid Parentheses
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    public static Dictionary<char, char> map = new Dictionary<char, char>()
    {
        {'(', ')'},
        {'[', ']'},
        {'{', '}'},
    };

    public bool IsValid(string s) 
    {
        Stack<char> stack = new Stack<char>();

        foreach (var c in s)
        {
            if (c == '(' || c == '[' || c == '{')
            {
                stack.Push(c);
            }
            else 
            {
                if (stack.Count == 0)
                {
                    return false;
                }

                var top = stack.Pop();
                if (map[top] != c)
                {
                    return false;
                }
            }
        }

        return stack.Count == 0;
    }
}
// @lc code=end

