/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || target < nums[0] || target > nums[n-1] {
		return []int{-1, -1}
	}
	t := binarySearch(nums, 0, n-1, target)
	if t == -1 {
		return []int{-1, -1}
	}
	left, right := t, t
	for i := t + 1; i < n; i++ {
		if nums[i] == target {
			right = i
		} else {
			break
		}
	}
	for i := t - 1; i >= 0; i-- {
		if nums[i] == target {
			left = i
		} else {
			break
		}
	}
	return []int{left, right}
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] < target {
		return binarySearch(nums, mid+1, right, target)
	}
	if nums[mid] > target {
		return binarySearch(nums, left, mid-1, target)
	}

	return mid
}

// @lc code=end

