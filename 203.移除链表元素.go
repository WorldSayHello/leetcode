/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	for tmp, last := head, head; tmp != nil; tmp = tmp.Next {
		if tmp.Val == val {
			last.Next = tmp.Next
			if tmp == head {
				head = head.Next
			}
			continue
		}
		last = tmp
	}
	return head
}

// @lc code=end

