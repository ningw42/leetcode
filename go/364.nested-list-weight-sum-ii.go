/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func depthSumInverse(nestedList []*NestedInteger) int {
    leveled := make(map[int][]int)
    depth := sumup(nestedList, 1, leveled)
    sum := 0
    for level, values := range leveled {
        for _, value := range values {
            sum += value * (depth - level)
        }
    }
    return sum
}

func sumup(list []*NestedInteger, level int, leveled map[int][]int) int {
    if len(list) == 0 {
        return 1
    }
    var newList []*NestedInteger
    for _, node := range list {
        if node.IsInteger() {
            leveled[level] = append(leveled[level], node.GetInteger())
        } else {
            newList = append(newList, node.GetList()...)
        }
    }
    return sumup(newList, level+1, leveled) + 1
}