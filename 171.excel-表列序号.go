/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	ans := 0
	for i := range columnTitle {
		ans = ans*26 + int(columnTitle[i]-'A') + 1
	}
	return ans
}

// @lc code=end

