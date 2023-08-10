/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	last := 0
	for i := 0; i < len(nums); i++ {
		if last == nums[i] {
			return nums[i]
		}
		last = nums[i]
	}
	return last
}

// @lc code=end

