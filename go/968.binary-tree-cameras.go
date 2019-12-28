/*
 * @lc app=leetcode id=968 lang=golang
 *
 * [968] Binary Tree Cameras
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	isCamera := find(root, true, true)
	isNotCamera := find(root, false, false)
	return minInts(isCamera, isNotCamera)
}

func find(root *TreeNode, isCamera bool, isCovered bool) int {
	if root == nil {
		return 0
	} else {
		if root.Left == nil && root.Right == nil {
			// leaf node
			// count camera if leaf node is a camera or is not a camera but not covered by a camera from its parent
			if isCamera || !isCovered {
				return 1
			} else {
				return 0
			}
		} else if root.Left != nil && root.Right == nil || root.Left == nil && root.Right != nil {
			// one child
			var child *TreeNode
			if root.Left != nil {
				child = root.Left
			} else {
				child = root.Right
			}
			if isCamera {
				// if current root is camera
				// child may or may not be camera
				return 1 + minInts(
					find(child, true, true), // child is camera
					find(child, false, true), // child is not camera
				)
			} else {
				// if current root is not camera
				if isCovered {
					// if current root is covered by a camera from its parent
					// child may or may not be camera
					return minInts(
						find(child, true, true), // child is camera
						find(child, false, false), // child is not camera
					)
				} else {
					// if current root is not covered by a camera from its parent
					// child is required to be a camera
					return find(child, true, true) 
				}
			}
		} else {
			// two children
			if isCamera {
				// if current root is camera
				// both children may or may not be camera
				leftIsCamera := find(root.Left, true, true)
				leftIsNotCamera := find(root.Left, false, true)
				rightIsCamera := find(root.Right, true, true)
				rightIsNotCamera := find(root.Right, false, true)
				return 1 + minInts(
					leftIsCamera + rightIsCamera, // both children are camera
					leftIsNotCamera + rightIsCamera, // only right child is camera
					leftIsCamera + rightIsNotCamera, // only left child is camera
					leftIsNotCamera + rightIsNotCamera, // both children are not camera
				)
			} else {
				// if current root is not camera
				leftIsCamera := find(root.Left, true, true)
				leftIsNotCamera := find(root.Left, false, false)
				rightIsCamera := find(root.Right, true, true)
				rightIsNotCamera := find(root.Right, false, false)
				if isCovered {
					// if current root is covered by a camera from its parent
					// both children may or may not be camera
					return minInts(
						leftIsCamera + rightIsCamera, // both children are camera
						leftIsNotCamera + rightIsCamera, // only right child is camera
						leftIsCamera + rightIsNotCamera, // only left child is camera
						leftIsNotCamera + rightIsNotCamera, // both children are not camera
					)
				} else {
					// if current root is not covered by a camera from its parent
					// at least one child is required to be camera to cover current root
					return minInts(
						leftIsCamera + rightIsCamera, // both children are camera
						leftIsNotCamera + rightIsCamera, // only right child is camera
						leftIsCamera + rightIsNotCamera, // only left child is camera
					)
				}
			}
		}
	}
}

func minInts(ints ...int) int {
	min := math.MaxInt64
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}

// @lc code=end

