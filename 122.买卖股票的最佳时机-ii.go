/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	incomes := make([]int, len(prices))
	last := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < last {
			incomes[i] = incomes[i-1]
		} else {
			incomes[i] += incomes[i-1] + (prices[i] - last)
		}
		last = prices[i]
	}
	return incomes[len(prices)-1]
}

// @lc code=end

