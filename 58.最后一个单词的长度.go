/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}
	size := 0
	for i := len(s) - 1; i >= 0; i-- {
		if size != 0 && s[i] == ' ' {
			break
		}
		if s[i] != ' ' {
			size++
		}
	}
	return size
}

// @lc code=end

