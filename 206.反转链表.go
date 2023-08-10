/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for travel := head; travel.Next != nil; {
		tmp := travel.Next     // 2
		travel.Next = tmp.Next // 1->3
		tmp.Next = head        // 2->1
		head = tmp             // 2
	}

	return head
}

// @lc code=end

