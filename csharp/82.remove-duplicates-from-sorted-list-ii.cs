/*
 * @lc app=leetcode id=82 lang=csharp
 *
 * [82] Remove Duplicates from Sorted List II
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

        ListNode prev, cur, dummy;
        dummy = new ListNode(-101, head); // dummy head
        prev = dummy;
        cur = head;

        while (cur != null && cur.next != null)
        {
            // consecutive duplicates count
            int duplicateCount = 0;
            while (cur.next != null && cur.val == cur.next.val)
            {
                cur = cur.next;
                duplicateCount++;
            }

            // basically to find a edge. (edge triggerred)
            if (duplicateCount == 0)
            {
                prev = cur;
                cur = cur.next;
            }
            else 
            {
                prev.next = cur.next;
                cur = cur.next;
            }
        }

        return dummy.next;
    }
}
// @lc code=end

