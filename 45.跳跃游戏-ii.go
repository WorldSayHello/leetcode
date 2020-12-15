/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	leftMax := 0
	steps := 0
	end := 0
	for i := 0; i < n-1; i++ {
		if i <= leftMax {
			leftMax = max(leftMax, i+nums[i])
			if i == end {
				end = leftMax
				steps++
			}
		}
	}
	return steps
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

