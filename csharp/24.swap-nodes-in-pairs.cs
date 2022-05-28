/*
 * @lc app=leetcode id=24 lang=csharp
 *
 * [24] Swap Nodes in Pairs
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
    public ListNode SwapPairs(ListNode head) 
    {
        ListNode dummy = new ListNode(int.MaxValue, head);

        ListNode prev = dummy;
        ListNode cur = head;
        ListNode next = null;

        while (cur != null && cur.next != null)
        {
            next = cur.next.next;
            prev.next = cur.next;
            cur.next.next = cur;
            cur.next = next;

            prev = cur;
            cur = next;
        }

        return dummy.next;
    }
}
// @lc code=end

