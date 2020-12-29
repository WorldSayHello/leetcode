import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
func getPermutation(n int, k int) string {

	var contains func([]int, int) bool
	contains = func(s []int, e int) bool {
		for _, t := range s {
			if t == e {
				return true
			}
		}
		return false
	}

	ret := make([]int, n)
	length := 0

	var backtrace func(int, []int)
	backtrace = func(m int, track []int) {
		if length == k {
			return
		}
		if len(track) == m {
			length++
			if length == k {
				copy(ret, track)
			}
			return
		}
		for i := 1; i <= m; i++ {
			if contains(track, i) {
				continue
			}
			track = append(track, i)
			backtrace(m, track)
			track = track[:len(track)-1]
		}
	}

	track := make([]int, 0)
	backtrace(n, track)

	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(ret[i])
	}

	return str
}

// @lc code=end

