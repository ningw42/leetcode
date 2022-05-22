/*
 * @lc app=leetcode id=104 lang=csharp
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

using System;
using System.Collections.Generic;

public class Solution 
{
    public int MaxDepth(TreeNode root) 
    {
        return this.MaxDepthIterative(root);
    }

    public int MaxDepthIterative(TreeNode root)
    {
        var s = new Stack<(TreeNode, int)>();
        s.Push((root, 0));

        int max = 0;
        while (s.Count > 0)
        {
            var (cur, h) = s.Pop();
            if (cur == null)
            {
                continue;
            }
            s.Push((cur.left, h + 1));
            s.Push((cur.right, h + 1));
            max = Math.Max(max, h + 1);
        }

        return max;
    }

    public int MaxDepthRecursive(TreeNode root)
    {
        if (root == null)
        {
            return 0;
        }
        else 
        {
            return Math.Max(this.MaxDepthRecursive(root.left), this.MaxDepthRecursive(root.right)) + 1;
        }
    }
}
// @lc code=end

