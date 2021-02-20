/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {

}

func jiangjinfu(heights []int) int {
	n := len(heights)
	max := 0
	for i := 0; i < n; i++ {
		left, right := i, i
		for j := i - 1; j >= 0; j-- {
			if heights[i] <= heights[j] {
				left--
			} else {
				break
			}
		}
		for j := i + 1; j < n; j++ {
			if heights[i] <= heights[j] {
				right++
			} else {
				break
			}
		}
		if (right-left+1)*heights[i] > max {
			max = (right - left + 1) * heights[i]
		}
		// fmt.Printf("%d,%d,%d\n", left, right, max)
	}
	return max
}

// @lc code=end

