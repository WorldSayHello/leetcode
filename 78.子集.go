/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var backtrace func(int)
	backtrace = func(cur int) {
		if cur > len(nums) {
			return
		}
		var cTmp = make([]int, len(tmp))
		copy(cTmp, tmp)
		res = append(res, cTmp)
		for j := cur; j < len(nums); j++ {
			tmp = append(tmp, nums[j])
			backtrace(j + 1)
			tmp = tmp[:len(tmp)-1]
		}

	}

	backtrace(0)
	return res
}

// @lc code=end

