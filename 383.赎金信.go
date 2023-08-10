/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	var tmp = make(map[rune]int)
	for _, s := range ransomNote {
		if _, ok := tmp[s]; ok {
			tmp[s] += 1
		} else {
			tmp[s] = 1
		}
	}
	for _, s := range magazine {
		if len(tmp) <= 0 {
			break
		}
		if _, ok := tmp[s]; ok {
			tmp[s] -= 1
			if tmp[s] <= 0 {
				delete(tmp, s)
			}
		}
	}
	return len(tmp) == 0
}

// @lc code=end

