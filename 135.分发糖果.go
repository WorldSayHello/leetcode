/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	condies := make([]int, n)
	condies[0] = 1
	// 1 1 2
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			condies[i] = condies[i-1] + 1
		} else {
			condies[i] = 1
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sum := condies[n-1]
	//  2 1 2
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			condies[i] = max(condies[i], condies[i+1]+1)
		} else {
			condies[i] = max(condies[i], 1)
		}
		sum += condies[i]
	}

	return sum
}

// @lc code=end

