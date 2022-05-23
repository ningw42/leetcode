/*
 * @lc app=leetcode id=203 lang=csharp
 *
 * [203] Remove Linked List Elements
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
    public ListNode RemoveElements(ListNode head, int val) 
    {
        if (head == null)
        {
            return null;
        }

        ListNode dummy = new ListNode(0, head);

        ListNode prev, cur;
        prev = dummy;
        cur = head;
        while (cur != null)
        {
            if (cur.val == val)
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

        return dummy.next;
    }
}
// @lc code=end

