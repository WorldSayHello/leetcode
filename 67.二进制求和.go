/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	l := aLen
	if bLen > l {
		l = bLen
	}
	aa := reverse(a)
	bb := reverse(b)
	flag := 0
	result := ""
	for i := 0; i < l; i++ {
		tmp := flag
		if i < aLen {
			tmp += int(aa[i] - '0')
		}
		if i < bLen {
			tmp += int(bb[i] - '0')
		}
		flag = tmp / 2
		d := tmp % 2
		if d == 0 {
			result += "0"
		} else {
			result += "1"
		}
	}
	if flag != 0 {
		result += "1"
	}

	return reverse(result)
}

func reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

// @lc code=end

