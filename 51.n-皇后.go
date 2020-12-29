/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	var queens [][]string

	var isInvalid func([]string, int, int) bool
	isInvalid = func(board []string, col int, row int) bool {
		if len(board) != n {
			return false
		}
		// 检查列
		for i := 0; i < n; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		// 检查左上角
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		// 检查右上角
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	board := make([]string, 0)

	var generateBoard func()
	generateBoard = func() {
		for i := 0; i < n; i++ {
			tmp := make([]byte, n)
			for j := 0; j < n; j++ {
				tmp[j] = '.'
			}
			board = append(board, string(tmp))
		}
	}

	var backtrace func(int)
	backtrace = func(row int) {
		if row == n {
			b := make([]string, n)
			copy(b, board)
			queens = append(queens, b)
			return
		}

		for col := 0; col < n; col++ {
			if !isInvalid(board, col, row) {
				continue
			}
			tmp := []byte(board[row])
			tmp[col] = 'Q'
			board[row] = string(tmp)
			backtrace(row + 1)
			tmp[col] = '.'
			board[row] = string(tmp)
		}
	}

	generateBoard()
	backtrace(0)

	return queens
}

// @lc code=end

