/*
 * @lc app=leetcode id=83 lang=csharp
 *
 * [83] Remove Duplicates from Sorted List
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
    public ListNode DeleteDuplicates(ListNode head) 
    {
        if (head == null)
        {
            return null;
        }

        ListNode prev, cur, next;
        prev = head;
        cur = head.next;

        while (cur != null)
        {
            if (cur.val == prev.val)
            {
                prev.next = cur.next;
                cur = cur.next;
            }
            else 
            {
                prev = cur;
                cur = cur.next;
            }
        }

        return head;
    }
}
// @lc code=end

