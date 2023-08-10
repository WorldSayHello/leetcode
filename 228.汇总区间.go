/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	length := len(nums)
	var i, j int
	var out []string
	var tmp string
	for i < length && j < length {
		if nums[i] == nums[j] {
			tmp = strconv.Itoa(nums[i])
		} else {
			tmp = strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j])
		}
		j++
		if j >= length || nums[j]-nums[j-1] > 1 {
			out = append(out, tmp)
			i = j
		}
	}
	return out
}

// @lc code=end

