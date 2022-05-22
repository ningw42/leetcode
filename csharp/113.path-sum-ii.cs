/*
 * @lc app=leetcode id=113 lang=csharp
 *
 * [113] Path Sum II
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
using System.Linq;

public class Solution 
{
    public IList<IList<int>> PathSum(TreeNode root, int targetSum) 
    {
        if (root == null)
        {
            return new List<IList<int>>();
        }
        else 
        {
            var pathes = this.PathSum2(root, targetSum);
            foreach (var path in pathes)
            {
                ((List<int>)path).Reverse();
            }
            return pathes;
        }
    }

    public IList<IList<int>> PathSum2(TreeNode root, int targetSum)
    {
        var subTarget = targetSum - root.val;
        List<IList<int>> result = new List<IList<int>>();

        if (root.left != null && root.right != null)
        {
            AddElement(result, this.PathSum2(root.left, subTarget), root.val);
            AddElement(result, this.PathSum2(root.right, subTarget), root.val);
        }
        else if (root.left != null && root.right == null)
        {
            AddElement(result, this.PathSum2(root.left, subTarget), root.val);
        }
        else if (root.left == null && root.right != null)
        {
            AddElement(result, this.PathSum2(root.right, subTarget), root.val);
        }
        else 
        {
            if (subTarget == 0)
            {
                result.Add(new List<int>() { root.val });
            }
        }

        return result;
    }

    public static void AddElement(IList<IList<int>> result, IList<IList<int>> pathes, int val)
    {
        foreach (var path in pathes)
        {
            path.Add(val);
            result.Add(path);
        }
    }
}
// @lc code=end

