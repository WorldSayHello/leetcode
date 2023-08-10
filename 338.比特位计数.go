/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[1] = 1
		if i%2 == 0 {
			bits[i] = bits[i/2]
		} else {
			bits[i] = bits[i/2] + 1
		}
	}
	return bits
}

// @lc code=end

