/*
 * @lc app=leetcode id=341 lang=csharp
 *
 * [341] Flatten Nested List Iterator
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool IsInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     int GetInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return null if this NestedInteger holds a single integer
 *     IList<NestedInteger> GetList();
 * }
 */

using System;
using System.Collections.Generic;

public class NestedIterator 
{
    private IEnumerator<NestedInteger> enumerator;

    private NestedIterator currentNestedIterator;

    private bool first;

    public NestedIterator(IList<NestedInteger> nestedList) 
    {
        this.enumerator = nestedList.GetEnumerator();
        this.currentNestedIterator = null;
        this.first = true;
    }

    public bool HasNext() 
    {
        if (this.first)
        {
            // the very first call to HasNext, move to next.
            this.first = false;
            bool hasNext = this.enumerator.MoveNext();
            if (!hasNext)
            {
                return false;
            }
        }
        else
        {
            if (this.enumerator.Current.IsInteger())
            {
                // current cursor points to a integer, move to next.
                // could be merged into the previous 'first' case
                bool hasNext = this.enumerator.MoveNext();
                if (!hasNext)
                {
                    return false;
                }
            }
            else
            {
                // current cursor points to a nested list
                // try to advance the inner cursor
                bool hasNestedNext = this.currentNestedIterator.HasNext();
                if (hasNestedNext)
                {
                    return true;
                }

                // or move to next element at current level if inner move fails.
                bool hasNext = this.enumerator.MoveNext();
                if (!hasNext)
                {
                    return false;
                }
            }
        }

        // current (level) cursor moved to next if code doesn't bait out eariler.
        // we need to inspect next element until there is a concrete integer(integer or nested list with available next element).
        while (true)
        {
            if (this.enumerator.Current.IsInteger())
            {
                return true;
            }

            this.currentNestedIterator = new NestedIterator(this.enumerator.Current.GetList());
            bool hasNestedNext = this.currentNestedIterator.HasNext();
            if (hasNestedNext)
            {
                return true;
            }

            bool hasNext = this.enumerator.MoveNext();
            if (!hasNext)
            {
                return false;
            }
        }
    }

    public int Next() 
    {
        if (this.enumerator.Current.IsInteger())
        {
            return this.enumerator.Current.GetInteger();
        }
        else 
        {
            return this.currentNestedIterator.Next();
        }
    }
}

/**
 * Your NestedIterator will be called like this:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.HasNext()) v[f()] = i.Next();
 */
// @lc code=end

