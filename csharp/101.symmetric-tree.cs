/*
 * @lc app=leetcode id=101 lang=csharp
 *
 * [101] Symmetric Tree
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

using System.Collections.Generic;

public class Solution {
    public bool IsSymmetric(TreeNode root) 
    {
        if (root == null)
        {
            return true;
        }
        // else 
        // {
        //     return this.IsMirrorRecursive(root.left, root.right);
        // }
        else 
        {
            return this.IsMirrorIterative(root.left, root.right);
        }
    }

    public bool IsMirrorRecursive(TreeNode a, TreeNode b)
    {
        if (a == null && b == null)
        {
            return true;
        }
        else if (a == null && b != null)
        {
            return false;
        }
        else if (a != null && b == null)
        {
            return false;
        }
        else 
        {
            return a.val == b.val
                && this.IsMirrorRecursive(a.left, b.right)
                && this.IsMirrorRecursive(a.right, b.left);
        }
    }

    public bool IsMirrorIterative(TreeNode a, TreeNode b)
    {
        Stack<TreeNode> s1 = new Stack<TreeNode>();
        Stack<TreeNode> s2 = new Stack<TreeNode>();

        s1.Push(a);
        s2.Push(b);

        while (s1.Count > 0 && s2.Count > 0)
        {
            var p = s1.Pop();
            var q = s2.Pop();
            if (p == null && q == null)
            {
                continue;
            }
            else if (p == null && q != null)
            {
                return false;
            }
            else if (p != null && q == null)
            {
                return false;
            }
            else 
            {
                if (p.val == q.val)
                {
                    s1.Push(p.left);
                    s1.Push(p.right);
                    s2.Push(q.right);
                    s2.Push(q.left);
                }
                else 
                {
                    return false;
                }
            }
        }

        return s1.Count == s2.Count; // == 0
    }
}
// @lc code=end

