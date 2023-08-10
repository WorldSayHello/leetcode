/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var zero = 0
	for i := 0; i < 32; i++ {
		flag := num & (1 << i)
		if flag > 0 {
			zero += 1
		}
	}
	return zero
}

// @lc code=end

