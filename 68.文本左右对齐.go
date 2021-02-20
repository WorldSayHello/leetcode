import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	if n == 0 {
		return words
	}

	var sprintf func(s []string, spaces int) string
	sprintf = func(s []string, spaces int) string {
		sn := len(s)
		spaces = spaces + sn - 1
		space := spaces
		left := 0
		if sn != 1 {
			space = spaces / (sn - 1)
			left = spaces % (sn - 1)
		}
		ret := s[0]
		for i := 1; i < sn; i++ {
			numsSpace := space
			if i <= left {
				numsSpace++
			}
			ret += fmt.Sprintf("%s%s", strings.Repeat(" ", numsSpace), s[i])
		}
		if sn == 1 {
			ret = fmt.Sprintf("%s%s", ret, strings.Repeat(" ", space))
		}
		return ret
	}

	var sprintfLast func(s []string, spaces int) string
	sprintfLast = func(s []string, spaces int) string {
		sn := len(s)
		if sn == 0 {
			return ""
		}
		ret := s[0]
		for i := 1; i < sn; i++ {
			ret += fmt.Sprintf("%s%s", " ", s[i])
		}
		ret += fmt.Sprintf("%s", strings.Repeat(" ", spaces))
		return ret
	}

	var wordRows []string
	row := make([]string, 0, 1)
	row = append(row, words[0])
	nowLen := len(words[0])
	i := 1
	for i = 1; i < n; i++ {
		nowLen = nowLen + len(words[i]) + 1
		if nowLen > maxWidth {
			tmp := sprintf(row, maxWidth-nowLen+len(words[i])+1)
			wordRows = append(wordRows, tmp)
			row = make([]string, 0, 1)
			nowLen = len(words[i])
		}
		row = append(row, words[i])
	}
	if i < n {
		row = append(row, words[i])
	}

	tmp := sprintfLast(row, maxWidth-nowLen)
	wordRows = append(wordRows, tmp)
	return wordRows
}

// @lc code=end

