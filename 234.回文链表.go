/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	last := head
	var reverse func(*ListNode) bool
	reverse = func(cur *ListNode) bool {
		if cur != nil {
			if !reverse(cur.Next) {
				return false
			}
			if cur.Val != last.Val {
				return false
			}
			last = last.Next
		}
		return true
	}
	return reverse(head)
}

// @lc code=end

