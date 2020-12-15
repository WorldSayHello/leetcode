import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	last := countAndSay(n - 1)
	start := 0
	ret := ""
	for i := 1; i < len(last)+1; i++ {
		if i == len(last) {
			ret += strconv.Itoa(i-start) + string(last[start])
		} else if last[i] != last[start] {
			ret += strconv.Itoa(i-start) + string(last[start])
			start = i
		}
	}
	return ret
}

// @lc code=end

