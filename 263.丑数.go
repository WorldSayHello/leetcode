/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	var list = []int{2, 3, 5}
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for _, i := range list {
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

// @lc code=end

