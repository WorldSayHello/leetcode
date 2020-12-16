/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	a := 0
	b := 0
	c := len(matrix) - 1
	d := len(matrix[0]) - 1

	result := make([]int, 0)
	for a <= c && b <= d {
		res := printCircles(a, b, c, d, matrix)
		a++
		b++
		c--
		d--
		result = append(result, res...)
	}
	return result
}

func printCircles(a int, b int, c int, d int, matrix [][]int) []int {

	result := make([]int, 0)
	if a == c {
		for i := b; i <= d; i++ {
			result = append(result, matrix[a][i])
		}
	} else if b == d {
		for i := a; i <= c; i++ {
			result = append(result, matrix[i][b])
		}
	} else {
		for i := b; i < d; i++ {
			result = append(result, matrix[a][i])
		}
		for i := a; i < c; i++ {
			result = append(result, matrix[i][d])
		}
		for i := d; i > b; i-- {
			result = append(result, matrix[c][i])
		}
		for i := c; i > a; i-- {
			result = append(result, matrix[i][b])
		}
	}
	return result
}

// @lc code=end

