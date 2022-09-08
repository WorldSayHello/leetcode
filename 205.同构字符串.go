/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	var table1 = make(map[byte]byte, 0)
	var table2 = make(map[byte]byte, 0)
	for i := range s {
		if v, ok := table1[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else if v, ok = table2[t[i]]; ok {
			return false
		} else {
			table1[s[i]] = t[i]
			table2[t[i]] = s[i]
		}
	}
	return true
}

// @lc code=end

