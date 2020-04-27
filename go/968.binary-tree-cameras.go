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
			if isCamera || !isCovered {
				return 1
			} else {
				return 0
			}
		} else if root.Left != nil && root.Right == nil || root.Left == nil && root.Right != nil {
			var child *TreeNode
			if root.Left != nil {
				child = root.Left
			} else {
				child = root.Right
			}
			if isCamera {
				return 1 + minInts(
					find(child, true, true),
					find(child, false, true),
				)
			} else {
				if isCovered {
					return minInts(
						find(child, true, true),
						find(child, false, false),
					)
				} else {
					return find(child, true, true)
				}
			}
		} else {
			if isCamera {
				leftIsCamera := find(root.Left, true, true)
				leftIsNotCamera := find(root.Left, false, true)
				rightIsCamera := find(root.Right, true, true)
				rightIsNotCamera := find(root.Right, false, true)
				return 1 + minInts(
					leftIsCamera + rightIsCamera,
					leftIsNotCamera + rightIsCamera,
					leftIsCamera + rightIsNotCamera,
					leftIsNotCamera + rightIsNotCamera,
				)
			} else {
				leftIsCamera := find(root.Left, true, true)
				leftIsNotCamera := find(root.Left, false, false)
				rightIsCamera := find(root.Right, true, true)
				rightIsNotCamera := find(root.Right, false, false)
				if isCovered {
					return minInts(
						leftIsCamera + rightIsCamera,
						leftIsNotCamera + rightIsCamera,
						leftIsCamera + rightIsNotCamera,
						leftIsNotCamera + rightIsNotCamera,
					)
				} else {
					return minInts(
						leftIsCamera + rightIsCamera,
						leftIsNotCamera + rightIsCamera,
						leftIsCamera + rightIsNotCamera,
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
