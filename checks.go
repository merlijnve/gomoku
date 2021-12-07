package main

func horizontal(x int, y int) int {
	boardSize := getBoardsize()

	player := board[x][y]
	if x+5 < boardSize && player != 0 &&
		player == board[x+1][y] &&
		player == board[x+2][y] &&
		player == board[x+3][y] &&
		player == board[x+4][y] {
		return player
	}
	return 0

}

func vertical(x int, y int) int {
	boardSize := getBoardsize()

	player := board[x][y]
	if y+5 < boardSize && player != 0 &&
		player == board[x][y+1] &&
		player == board[x][y+2] &&
		player == board[x][y+3] &&
		player == board[x][y+4] {
		return player
	}
	return 0

}

func diagonalRight(x int, y int) int {
	boardSize := getBoardsize()

	player := board[x][y]
	if y+5 < boardSize && x+5 < boardSize && player != 0 &&
		player == board[x+1][y+1] &&
		player == board[x+2][y+2] &&
		player == board[x+3][y+3] &&
		player == board[x+4][y+4] {
		return player
	}
	return 0

}

func diagonalLeft(x int, y int) int {
	boardSize := getBoardsize()

	player := board[x][y]
	if y+5 < boardSize && x-5 > 0 && player != 0 &&
		player == board[x-1][y+1] &&
		player == board[x-2][y+2] &&
		player == board[x-3][y+3] &&
		player == board[x-4][y+4] {
		return player
	}
	return 0

}

func checkWin() int {
	boardSize := getBoardsize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			if horizontal(x, y) != 0 {
				return horizontal(x, y)
			}
			if vertical(x, y) != 0 {
				return vertical(x, y)
			}
			if diagonalLeft(x, y) != 0 {
				return diagonalLeft(x, y)
			}
			if diagonalRight(x, y) != 0 {
				return diagonalRight(x, y)
			}
		}
	}
	return 0
}
