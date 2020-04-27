/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    list []*NestedInteger   // containing list
    cursor int              // cursor points to next item to access
    itor *NestedIterator    // if the item that cursor points to is a list, itor is
                            // the iterator of that list
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        list: nestedList,
        cursor: 0,
        itor: nil,
    }
}

func (this *NestedIterator) Next() int {
    if this.list[this.cursor].IsInteger() {
        // next item is integer
        value := this.list[this.cursor].GetInteger()
        // advance the cursor for next iteration
        this.cursor++
        return value
    } else {
        // next item is in subordinate list
        return this.itor.Next()
    }
}

func (this *NestedIterator) HasNext() bool {
    // cursor exceeds list length
    if this.cursor >= len(this.list) {
        return false
    }
    if this.list[this.cursor].IsInteger() {
        // next item to access is an integer
        return true
    } else {
        // next item to access is a subordinate list
        if this.itor == nil {
            // very first time to access the subordinate list
            // initialize iterator for that list
            this.itor = Constructor(this.list[this.cursor].GetList())
        }
        if this.itor.HasNext() {
            // subordinate list has next
            return true
        } else {
            // subordinate list has no next
            // reset iterator of last subordinate list
            this.itor = nil
            // try the next item in current list
            this.cursor++
            return this.HasNext()
        }
    }
}