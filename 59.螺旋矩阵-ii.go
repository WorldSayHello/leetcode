/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	a := 0
	b := 0
	c := n - 1
	d := n - 1
	e := 1
	for a <= c && b <= d {
		e = printCircles(a, b, c, d, e, matrix)
		a++
		b++
		c--
		d--
	}
	return matrix
}

func printCircles(a int, b int, c int, d int, n int, matrix [][]int) int {
	if a == c {
		for i := b; i <= d; i++ {
			matrix[a][i] = n
			n++
		}
	}

	for i := b; i < d; i++ {
		matrix[a][i] = n
		n++

	}
	for i := a; i < c; i++ {
		matrix[i][d] = n
		n++
	}
	for i := d; i > b; i-- {
		matrix[c][i] = n
		n++
	}
	for i := c; i > a; i-- {
		matrix[i][b] = n
		n++
	}
	return n
}

// @lc code=end

