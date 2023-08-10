/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1 := len(num1)
	len2 := len(num2)
	sum := "0"
	for i := len1 - 1; i >= 0; i-- {
		tmp := "0"
		s1, _ := strconv.Atoi(string(num1[i]))
		for m := 0; m < s1; m++ {
			tmp = addTwoNum(tmp, num2)
		}
		tmp = tmp + strings.Repeat("0", len1-1-i)
		sum = addTwoNum(sum)
	}

}

func addTwoNum(num1 string, num2 string) string {

}

// @lc code=end

