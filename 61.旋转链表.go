/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	helper := make(map[int]*ListNode)
	tmp := head
	i := 0
	for tmp != nil {
		i++
		helper[i] = tmp
		tmp = tmp.Next
	}
	k = k % i
	if k != 0 {
		oldHead := head
		head = helper[i-k+1]
		helper[i].Next = oldHead
		helper[i-k].Next = nil
	}
	return head
}

// @lc code=end

