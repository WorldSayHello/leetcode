/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length && j <= i+k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// @lc code=end

