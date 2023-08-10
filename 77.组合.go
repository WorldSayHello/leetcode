/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	var backtrace func(int)
	backtrace = func(cur int) {
		if cur > n+1 || len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			var cTmp = make([]int, k)
			copy(cTmp, tmp)
			res = append(res, cTmp)
		}
		for i := cur; i <= n; i++ {
			tmp = append(tmp, i)
			backtrace(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrace(1)
	return res
}

// @lc code=end

