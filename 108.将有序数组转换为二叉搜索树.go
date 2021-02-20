/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	return inorderTraversal(nums, 0, len(nums)-1)
}

func inorderTraversal(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2

	root := &TreeNode{Val: nums[mid]}
	root.Left = inorderTraversal(nums, left, mid-1)
	root.Right = inorderTraversal(nums, mid+1, right)
	return root
}

// @lc code=end

