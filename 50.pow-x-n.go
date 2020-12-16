/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / quickPow(x, n)
	}
	return quickPow(x, n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x

}

// @lc code=end

