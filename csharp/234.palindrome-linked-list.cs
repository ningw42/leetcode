/*
 * @lc app=leetcode id=234 lang=csharp
 *
 * [234] Palindrome Linked List
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
    // O(n)
    public bool IsPalindrome(ListNode head) 
    {
        // get length
        int length = 0;
        ListNode cur = head;
        while (cur != null)
        {
            cur = cur.next;
            length++;
        }

        // single element list is a palindrome, length guaranteed no less than 1
        if (length == 1)
        {
            return true;
        }

        // reverse first half of the list
        var (left, right) = this.ReverseLinkedList(head, length / 2);
        // ignore the center element
        if (length % 2 != 0)
        {
            right = right.next;
        }

        // compare two (half) list
        while (left != null) // right != null
        {
            if (left.val != right.val)
            {
                return false;
            }
            left = left.next;
            right = right.next;
        }

        return true;
    }

    // Reverse first 'count' element in a list
    // (head of the reversed list, head of the rest list)
    public (ListNode, ListNode) ReverseLinkedList(ListNode head, int count)
    {
        int i = 0;
        
        ListNode prev, cur, next;
        next = null;
        prev = null;
        cur = head;
        while (i < count)
        {
            next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
            i++;
        }

        return (prev, cur);
    }
}
// @lc code=end

