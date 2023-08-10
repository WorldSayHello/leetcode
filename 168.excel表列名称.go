/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	s := []byte{}

	for columnNumber > 0 {
		tmp := columnNumber % 26
		if tmp == 0 {
			tmp = 26
		}
		columnNumber = columnNumber / 26
		s = append(s, byte('A'+tmp-1))
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}

// @lc code=end

