package backtrack

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	var ans [][]string
	// 构造初始棋盘
	board := make([]string, n)
	initialRow := strings.Repeat(".", n)
	for i := 0; i < n; i++ {
		board[i] = initialRow
	}

	var backtrack func(int)
	backtrack = func(row int) {
		// 到最后一行，将结果加入答案
		if row == n {
			tmp := make([]string, n)
			copy(tmp, board)
			ans = append(ans, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			// 选择当前位置
			board[row] = generatorBoard(n, col)
			backtrack(row + 1)
			// 撤销选择
			board[row] = initialRow
		}
	}

	backtrack(0)
	return ans
}

func isValid(board []string, row, col int) bool {
	n := len(board)
	// 检查列冲突
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查右上对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查左上对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func generatorBoard(size, col int) string {
	var sb strings.Builder
	for i := 0; i < size; i++ {
		if i == col {
			sb.WriteByte('Q')
		} else {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}
