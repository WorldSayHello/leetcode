import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	paths := strings.Split(path, "/")
	for i := 0; i < len(paths); i++ {
		if paths[i] == "." {
			continue
		}
		if paths[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if len(paths[i]) > 0 {
			stack = append(stack, paths[i])
		}
	}

	return string('/') + strings.Join(stack, "/")
}

// @lc code=end

