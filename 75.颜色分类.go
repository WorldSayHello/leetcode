/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {

	left := 0
	right := len(nums) - 1
	for i := left; i <= right; i++ {
		for nums[left] == 0 && left < right {
			left++
		}
		for nums[right] == 2 && left < right {
			right--
		}
		if nums[i] == 2 && i < right {
			tmp := nums[i]
			nums[i] = nums[right]
			nums[right] = tmp
		}
		if nums[i] == 0 && i > left {
			tmp := nums[i]
			nums[i] = nums[left]
			nums[left] = tmp
		}
	}
}

// @lc code=end

