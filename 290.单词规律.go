/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	var table1 = make(map[byte]string)
	var table2 = make(map[string]byte)
	s1 := strings.Split(s, " ")
	if len(pattern) != len(s1) {
		return false
	}
	for i := range pattern {
		if v, ok := table1[pattern[i]]; ok {
			if v != s1[i] {
				return false
			}
		} else if v2, ok := table2[s1[i]]; ok {
			if v2 != pattern[i] {
				return false
			}
		} else {
			table1[pattern[i]] = s1[i]
			table2[s1[i]] = pattern[i]
		}
	}
	return true
}

// @lc code=end

