/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	var isInteger = func(s string) bool {
		if len(s) <= 0 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if !unicode.IsDigit(rune(s[i])) {
				return false
			}
		}
		return true
	}

	var isDecimal = func(s string) bool {
		if len(s) <= 0 {
			return false
		}
		if s[0] == '.' {
			if len(s) == 1 {
				return false
			}
			return isInteger(s[1:])
		}
		if s[len(s)-1] == '.' {
			if len(s) == 1 {
				return false
			}
			return isInteger(s[:len(s)-1])
		}
		nums := strings.Split(s, ".")
		if len(nums) != 2 {
			return false
		}
		return isInteger(nums[0]) && isInteger(nums[1])
	}

	var isSignNum = func(s string) bool {
		if len(s) <= 0 {
			return false
		}
		if s[0] == '+' || s[0] == '-' {
			if len(s) == 1 {
				return false
			}
			s = s[1:]
		}
		return isInteger(s) || isDecimal(s)
	}

	var isSignInteger = func(s string) bool {
		if len(s) <= 0 {
			return false
		}
		if s[0] == '+' || s[0] == '-' {
			if len(s) == 1 {
				return false
			}
			s = s[1:]
		}
		return isInteger(s)
	}

	s = strings.ToLower(s)
	nums := strings.Split(s, "e")
	if len(nums) == 1 {
		return isSignNum(nums[0])
	}
	if len(nums) != 2 {
		return false
	}

	return isSignNum(nums[0]) && isSignInteger(nums[1])
}

// @lc code=end

