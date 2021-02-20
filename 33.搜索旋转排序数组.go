/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	return binarySearch(nums, 0, n-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	mid := (left + right) / 2
	if left > right {
		return -1
	}
	if nums[mid] == target {
		return mid
	}

	if nums[left] <= nums[mid] {
		if nums[left] < target && nums[right] > target {
			return binarySearch(nums, left, right, target)
		}
	}

	// 有序
	if (nums[mid] < target && nums[mid] < nums[right]) || (nums[mid] > target && nums[mid] < nums[0]) {
		return binarySearch(nums, mid+1, right, target)
	} else if (nums[mid] < target && nums[mid] > nums[right]) || (nums[mid] > target && nums[mid] > nums[0]) {
		return binarySearch(nums, left, mid-1, target)
	}
	return -1

}

// @lc code=end

