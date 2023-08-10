/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	var i, j int
	for j < length {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}

}

// @lc code=end

