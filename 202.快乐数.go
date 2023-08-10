/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	var last = make(map[int]struct{})
	for n != 1 {
		if _, ok := last[n]; ok {
			break
		}
		last[n] = struct{}{}
		var tmp = n
		n = 0
		for tmp != 0 {
			n += (tmp % 10) * (tmp % 10)
			tmp /= 10
		}
	}

	return n == 1
}

// @lc code=end

