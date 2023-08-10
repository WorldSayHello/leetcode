/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	table := make(map[int]struct{})
	ret := make([]int, 0)
	for _, item := range nums1 {
		table[item] = struct{}{}
	}
	for _, item := range nums2 {
		if _, ok := table[item]; ok {
			ret = append(ret, item)
			delete(table, item)
		}
	}
	return ret
}

// @lc code=end

