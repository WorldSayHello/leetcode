/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	var s = make([]byte, 0)
	for columnNumber > 0 {
		columnNumber--
		tmp := columnNumber % 26
		columnNumber /= 26
		s = append(s, 'A'+byte(tmp))
	}
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return string(s)
}

// @lc code=end

