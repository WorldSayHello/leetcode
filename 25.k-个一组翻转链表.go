/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	step := 0
	// var faster *ListNode = head
	var curr *ListNode = head
	var prev *ListNode = nil

	var start *ListNode = head
	var preStart *ListNode = nil

	var h *ListNode = head

	for curr != nil {
		if step%k == 0 {
			if preStart != nil {
				preStart.Next = prev
			}
			preStart = start
			start = curr
			prev = nil
		} else {

			start.Next = curr
		}
		if step == k-1 {
			h = curr
		}
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		step++
	}

	return h
}

// @lc code=end

