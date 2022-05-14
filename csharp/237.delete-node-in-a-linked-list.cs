/*
 * @lc app=leetcode id=237 lang=csharp
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int x) { val = x; }
 * }
 */
public class Solution 
{
    public void DeleteNode(ListNode node) 
    {
        ListNode prev = null;
        while (node.next != null)
        {
            node.val = node.next.val;
            prev = node;
            node = node.next;
        }

        prev.next = null;
    }
}
// @lc code=end

