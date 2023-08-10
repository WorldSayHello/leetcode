/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	// 首先是2次幂，然后偶数位为1
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

// @lc code=end

