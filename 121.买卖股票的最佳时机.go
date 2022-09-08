/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	var profit = 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}
	return profit
}

// @lc code=end

