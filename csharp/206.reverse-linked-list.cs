/*
 * @lc app=leetcode id=206 lang=csharp
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution 
{
    public ListNode ReverseList(ListNode head) 
    {
        if (head == null)
        {
            return null;
        }

        ListNode prev, current, next;
        prev = null;
        current = head;
        next = head.next;

        while (next != null)
        {
            current.next = prev;
            prev = current;
            current = next;
            next = next.next;
        }
        current.next = prev;

        return current;
    }
}
// @lc code=end

