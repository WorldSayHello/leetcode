/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	leftMax := 0
	for i := 0; i < n; i++ {
		if i <= leftMax {
			leftMax = max(leftMax, i+nums[i])
			if leftMax >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

