package main

func isValidSudoku(board [][]byte) bool {
	// Проверяем строки
	hashRow := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		hashRow = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				val := board[i][j]
				_, ok := hashRow[val]
				if ok {
					return false
				}
				hashRow[val] = true
			}
		}
	}

	// Проверяем колонные
	hashCol := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		hashCol = make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				val := board[j][i]
				_, ok := hashCol[val]
				if ok {
					return false
				}
				hashCol[val] = true
			}
		}
	}

	hashBox := make(map[byte]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			hashBox = make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for b := 0; b < 3; b++ {
					if board[k+i*3][b+j*3] != '.' {
						val := board[k+i*3][b+j*3]
						_, ok := hashBox[val]
						if ok {
							return false
						}
						hashBox[val] = true
					}
				}
			}
		}
	}

	return true
}
