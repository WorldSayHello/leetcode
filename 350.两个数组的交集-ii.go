/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1 := len(nums1)
	length2 := len(nums2)
	index1 := 0
	index2 := 0
	var ret = make([]int, 0)
	for index1 < length1 && index2 < length2 {
		if nums1[index1] == nums2[index2] {
			ret = append(ret, nums1[index1])
			index1++
			index2++
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}
	return ret
}

// @lc code=end

