/*
 * @lc app=leetcode id=2 lang=csharp
 *
 * [2] Add Two Numbers
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
    // add to l1
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) 
    {
        ListNode head = l1;
        ListNode prev = null;

        int carry = 0;
        int sum = 0;
        while (l1 != null && l2 != null)
        {
            sum = l1.val + l2.val + carry;
            l1.val = sum % 10;
            carry = sum / 10;

            prev = l1;
            l1 = l1.next;
            l2 = l2.next;
        }

        if (l1 != null && l2 == null)
        {
            // still add to l1
            this.Add(l1, carry);
        }
        else if (l1 == null && l2 != null)
        {
            // add to l2 and link l1 and the rest of l2
            prev.next = this.Add(l2, carry);
        }
        else 
        {
            // create carry node when l1 and l2 have the same length
            if (carry != 0)
            {
                prev.next = new ListNode(carry, null);
            }
        }

        return head;
    }

    // add carry to list and return head
    public ListNode Add(ListNode head, int carry)
    {
        if (carry == 0)
        {
            return head;
        }

        ListNode h = head;
        ListNode prev = null;

        int sum = 0;
        while (head != null)
        {
            sum = head.val + carry;
            head.val = sum % 10;
            carry = sum / 10;

            prev = head;
            head = head.next;
        }

        if (carry != 0)
        {
            prev.next = new ListNode(carry, null);
        }

        return h;
    }
}
// @lc code=end

