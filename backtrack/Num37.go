package backtrack

func solveSudoku(board [][]byte) {
	solveSudokuBacktrack(0, 0, board)
}

func solveSudokuBacktrack(i, j int, board [][]byte) bool {
	if i == 9 {
		return true
	}
	if j == 9 {
		return solveSudokuBacktrack(i+1, 0, board)
	}
	if board[i][j] != '.' {
		return solveSudokuBacktrack(i, j+1, board)
	}

	for k := '1'; k <= '9'; k++ {
		if !solveSudokuIsValid(i, j, byte(k), board) {
			continue
		}
		board[i][j] = byte(k)
		if solveSudokuBacktrack(i, j+1, board) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func solveSudokuIsValid(row, col int, num byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == byte(num) {
			return false
		}
		if board[i][col] == byte(num) {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == byte(num) {
			return false
		}
	}
	return true
}
