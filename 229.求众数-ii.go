/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums) / 3
	var out = make([]int, 0)
	var last = -13679
	var count = 1
	for _, num := range nums {
		if num == last {
			count++
		} else {
			last = num
			count = 1
		}
		if count == n+1 {
			out = append(out, num)
		}
	}
	return out
}

// @lc code=end

