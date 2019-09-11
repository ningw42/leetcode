/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	vals []int
	next int
}


func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		vals: []int{},
		next: 0,
	}
	iter.PreOrder(root)
	return iter
}

func (this *BSTIterator) PreOrder(root *TreeNode) {
	if root != nil {
		this.PreOrder(root.Left)
		this.vals = append(this.vals, root.Val)
		this.PreOrder(root.Right)
	}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	result := this.vals[this.next]
	this.next++
	return result
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.next < len(this.vals)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

