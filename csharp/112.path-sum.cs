/*
 * @lc app=leetcode id=112 lang=csharp
 *
 * [112] Path Sum
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
public class Solution 
{
    public bool HasPathSum(TreeNode root, int targetSum) 
    {
        if (root == null)
        {
            return false;
        }
        else 
        {
            var subTarget = targetSum - root.val;
            if (root.left != null && root.right != null)
            {
                return this.HasPathSum(root.left, subTarget)
                    || this.HasPathSum(root.right, subTarget);
            }
            else if (root.left != null && root.right == null)
            {
                return this.HasPathSum(root.left, subTarget);
            }
            else if (root.left == null && root.right != null)
            {
                return this.HasPathSum(root.right, subTarget);
            }
            else 
            {
                return subTarget == 0;
            }
        }
    }
}
// @lc code=end

