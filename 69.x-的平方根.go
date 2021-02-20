/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x
	ret := left
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			ret = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ret
}

// @lc code=end

