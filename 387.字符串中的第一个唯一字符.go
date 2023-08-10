/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	tmp := make([]int, 26)
	for i := range s {
		tmp[s[i]-'a'] += 1
	}
	for i := range s {
		if tmp[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

