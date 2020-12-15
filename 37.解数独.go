/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	backtrace(board)
}

func backtrace(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := 0; k < 9; k++ {
				if isInvalid(i, j, byte(k+'1'), board) {
					board[i][j] = byte(k + '1')
					if backtrace(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}

	return true
}

func isInvalid(row int, col int, val byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[j][col] == val {
			return false
		}
	}

	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

// @lc code=end

