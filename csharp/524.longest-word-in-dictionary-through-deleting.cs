/*
 * @lc app=leetcode id=524 lang=csharp
 *
 * [524] Longest Word in Dictionary through Deleting
 */

// @lc code=start

using System.Collections.Generic;

public class Solution 
{
    private Dictionary<char, int[]> occurrences = new Dictionary<char, int[]>();

    public string FindLongestWord(string s, IList<string> dictionary) 
    {
        this.FindOccurrences(s);

        string result = string.Empty;
        foreach (string word in dictionary)
        {
            if (this.IsWordPossible(word))
            {
                if (word.Length > result.Length)
                {
                    result = word;
                }
                else if (word.Length == result.Length)
                {
                    if (string.Compare(word, result) < 0)
                    {
                        result = word;
                    }
                }
            }
        }

        return result;
    }

    public void FindOccurrences(string s)
    {
        Dictionary<char, IList<int>> occurrences = new Dictionary<char, IList<int>>();

        for (int i = 0; i < s.Length; i++)
        {
            if (occurrences.TryGetValue(s[i], out IList<int> occurrence))
            {
                occurrence.Add(i);
            }
            else
            {
                occurrences[s[i]] = new List<int>() { i };
            }
        }

        foreach (var occurrence in occurrences)
        {
            this.occurrences[occurrence.Key] = occurrence.Value.ToArray();
        }
    }

    public bool IsWordPossible(string word)
    {
        int start = -1;
        foreach (char c in word)
        {
            if (!this.occurrences.TryGetValue(c, out int[] occurrence))
            {
                return false;
            }
            else 
            {
                // could be reduced to O(nlog(n)) by binary search
                bool found = false;
                foreach (int index in occurrence)
                {
                    if (index > start)
                    {
                        found = true;
                        start = index;
                        break;
                    }
                }

                if (!found)
                {
                    return false;
                }
            }
        }

        return true;
    }
}
// @lc code=end

