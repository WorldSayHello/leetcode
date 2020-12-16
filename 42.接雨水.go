/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	first := 0
	second := 1
	if height[0] < height[1] {
		first = 1
		second = 0
	}
	for i := 2; i < n; i++ {
		if height[i] > height[first] {
			second = first
			first = i
		} else if height[i] > height[second] {
			second = i
		}
	}

	h := height[second]
	if first > second {
		tmp := first
		first = second
		second = tmp
	}
	mid := h * (second - first - 1)
	for i := first + 1; i < second; i++ {
		mid -= height[i]
	}

	return trap(height[0:first+1]) + trap(height[second:]) + mid
}

// @lc code=end

