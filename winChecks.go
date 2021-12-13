package main

func horizontalWin(x int, y int) int {
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

func verticalWin(x int, y int) int {
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

func diagonalRightWin(x int, y int) int {
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

func diagonalLeftWin(x int, y int) int {
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
			if horizontalWin(x, y) != 0 {
				return horizontalWin(x, y)
			}
			if verticalWin(x, y) != 0 {
				return verticalWin(x, y)
			}
			if diagonalLeftWin(x, y) != 0 {
				return diagonalLeftWin(x, y)
			}
			if diagonalRightWin(x, y) != 0 {
				return diagonalRightWin(x, y)
			}
		}
	}
	return 0
}
