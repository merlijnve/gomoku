package main

import "fmt"

func horizontalThree(x int, y int) int {
	boardSize := getBoardsize()

	if x+6 < boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == board[x+1][y] &&
			board[x+3][y] == board[x+1][y] &&
			board[x+4][y] == 0 {
			fmt.Println("found -***- horizontal")
			return board[x+1][y]
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == 0 &&
			board[x+3][y] == board[x+1][y] &&
			board[x+4][y] == board[x+1][y] &&
			board[x+5][y] == 0 {
			fmt.Println("found -*-**- horizontal")

			return board[x+1][y]
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == board[x+1][y] &&
			board[x+3][y] == 0 &&
			board[x+4][y] == board[x+1][y] &&
			board[x+5][y] == 0 {
			fmt.Println("found -**-*- horizontal")
			return board[x+1][y]
		}
	}
	return 0
}

func verticalThree(x int, y int) int {
	boardSize := getBoardsize()

	if y+6 < boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == board[x][y+1] &&
			board[x][y+3] == board[x][y+1] &&
			board[x][y+4] == 0 {
			fmt.Println("found -***- vertical")
			return board[x][y+1]
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == 0 &&
			board[x][y+3] == board[x][y+1] &&
			board[x][y+4] == board[x][y+1] &&
			board[x][y+5] == 0 {
			fmt.Println("found -*-**- vertical")
			return board[x][y+1]
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == board[x][y+1] &&
			board[x][y+3] == 0 &&
			board[x][y+4] == board[x][y+1] &&
			board[x][y+5] == 0 {
			fmt.Println("found -**-*- vertical")
			return board[x][y+1]
		}
	}
	return 0
}

func diagonalRightThree(x int, y int) int {
	boardSize := getBoardsize()

	if y+6 < boardSize && x+6 < boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == board[x+1][y+1] &&
			board[x+3][y+3] == board[x+1][y+1] &&
			board[x+4][y+4] == 0 {
			fmt.Println("found -***- diagonal")
			return board[x+1][y+1]
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == 0 &&
			board[x+3][y+3] == board[x+1][y+1] &&
			board[x+4][y+4] == board[x+1][y+1] &&
			board[x+5][y+5] == 0 {
			fmt.Println("found -*-**- diagonal")
			return board[x+1][y+1]
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == board[x+1][y+1] &&
			board[x+3][y+3] == 0 &&
			board[x+4][y+4] == board[x+1][y+1] &&
			board[x+5][y+5] == 0 {
			fmt.Println("found -**-*- diagonal")
			return board[x+1][y+1]
		}
	}
	return 0
}

func diagonalLeftThree(x int, y int) int {
	boardSize := getBoardsize()

	if y+6 < boardSize && x-6 > 0 {
		// -***-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == board[x-1][y+1] &&
			board[x-3][y+3] == board[x-1][y+1] &&
			board[x-4][y+4] == 0 {
			fmt.Println("found -***- diagonal")
			return board[x-1][y+1]
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == 0 &&
			board[x-3][y+3] == board[x-1][y+1] &&
			board[x-4][y+4] == board[x-1][y+1] &&
			board[x-5][y+5] == 0 {
			fmt.Println("found -*-**- diagonal")
			return board[x-1][y+1]
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == board[x-1][y+1] &&
			board[x-3][y+3] == 0 &&
			board[x-4][y+4] == board[x-1][y+1] &&
			board[x-5][y+5] == 0 {
			fmt.Println("found -**-*- diagonal")
			return board[x-1][y+1]
		}
	}
	return 0
}

func checkFreeThrees() int {
	boardSize := getBoardsize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			if horizontalThree(x, y) != 0 {
				return horizontalThree(x, y)
			}
			if verticalThree(x, y) != 0 {
				return verticalThree(x, y)
			}
			if diagonalRightThree(x, y) != 0 {
				return diagonalRightThree(x, y)
			}
			if diagonalLeftThree(x, y) != 0 {
				return diagonalLeftThree(x, y)
			}
		}
	}
	return 0
}
