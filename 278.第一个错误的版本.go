/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var left = 1
	var rigth = n
	for left < rigth {
		var mid = (left + rigth) / 2
		if isBadVersion(mid) {
			rigth = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

