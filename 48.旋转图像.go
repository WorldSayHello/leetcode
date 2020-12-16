/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	var tmp, i, j int
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			tmp = matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}

	for i = 0; i < n; i++ {
		for j = 0; j < n/2; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
}

// @lc code=end

