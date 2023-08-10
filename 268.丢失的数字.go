/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	length := len(nums)
	tmp := (1 + length) * length / 2
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	return tmp - sum
}

// @lc code=end

