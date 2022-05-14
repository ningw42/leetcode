/*
 * @lc app=leetcode id=21 lang=csharp
 *
 * [21] Merge Two Sorted Lists
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
    public ListNode MergeTwoLists(ListNode list1, ListNode list2) 
    {
        // return MergeRecursively(list1, list2);
        return MergeIteratively(list1, list2);
    }

    public ListNode MergeIteratively(ListNode a, ListNode b)
    {
        if (a == null && b == null)
        {
            return null;
        }
        else if (a == null && b != null)
        {
            return b;
        }
        else if (a != null && b == null)
        {
            return a;
        }

        ListNode dummy = new ListNode();
        ListNode prev = dummy;

        while (a != null && b != null)
        {
            if (a.val < b.val)
            {
                prev.next = a;
                prev = a;
                a = a.next;
            }
            else 
            {
                prev.next = b;
                prev = b;
                b = b.next;
            }
        }

        if (a == null)
        {
            prev.next = b;
        }
        else 
        {
            prev.next = a;
        }

        return dummy.next;
    }

    public ListNode MergeRecursively(ListNode a, ListNode b)
    {
        if (a == null && b == null)
        {
            return null;
        }
        else if (a == null && b != null)
        {
            return b;
        }
        else if (a != null && b == null)
        {
            return a;
        }
        else 
        {
            if (a.val < b.val)
            {
                a.next = MergeRecursively(a.next, b);
                return a;
            }
            else 
            {
                b.next = MergeRecursively(a, b.next);
                return b;
            }
        }
    }
}
// @lc code=end

