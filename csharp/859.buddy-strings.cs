/*
 * @lc app=leetcode id=859 lang=csharp
 *
 * [859] Buddy Strings
 */

// @lc code=start

using System;
using System.Collections.Generic;

public class Solution 
{
    public bool BuddyStrings(string s, string goal) 
    {
        // sanity check
        if (s.Length <= 1 || s.Length != goal.Length)
        {
            return false;
        }

        // chars distribution and diff counter
        Dictionary<char, int> sDistri = new Dictionary<char, int>();
        Dictionary<char, int> goalDistri = new Dictionary<char, int>();
        List<char> diffQueue = new List<char>();
        Stack<char> diffStack = new Stack<char>();
        for (var i = 0; i < s.Length; i++)
        {
            // bait out if diff counter larger than 2
            if (diffQueue.Count > 2)
            {
                return false;
            }

            if (s[i] != goal[i])
            {
                diffQueue.Add(s[i]);
                diffStack.Push(goal[i]);
            }

            if (sDistri.TryGetValue(s[i], out int v1))
            {
                sDistri[s[i]] = v1 + 1;
            }
            else 
            {
                sDistri[s[i]] = 1;
            }

            if (goalDistri.TryGetValue(goal[i], out int v2))
            {
                goalDistri[goal[i]] = v2 + 1;
            }
            else 
            {
                goalDistri[goal[i]] = 1;
            }
        }

        if (diffQueue.Count == 0)
        {
            // no diff between s and goal, check if there is a char with the same occurances
            // in both s and goal (>=2).
            foreach (var kv in sDistri)
            {
                if (kv.Value > 1 && goalDistri[kv.Key] == kv.Value)
                {
                    return true;
                }
            }
            return false;
        }
        else 
        {
            // check diff order
            var queue = diffQueue.ToArray();
            var stack = diffStack.ToArray();
            return queue[0] == stack[0] && queue[1] == stack[1];
        }
    }
}
// @lc code=end

