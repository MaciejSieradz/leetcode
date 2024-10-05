package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	var rows [9]map[byte]bool
	var cols [9]map[byte]bool
	var squares [9]map[byte]bool

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j]

				if rows[i][num] {
					return false
				}

				rows[i][num] = true

				if cols[j][num] {
					return false
				}

				cols[j][num] = true

				squareIndex := (i/3)*3 + j/3
				if squares[squareIndex][num] {
					return false
				}
				squares[squareIndex][num] = true
			}
		}
	}

	return true
}
