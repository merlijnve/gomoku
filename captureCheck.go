package main

import (
	"fmt"
)

func captureCheck() {
	boardSize := getBoardsize()

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			// from {x,y} to {x + 3,y}
			if x+3 < boardSize &&
				board[x][y] != 0 &&
				board[x+1][y] != 0 && board[x+1][y] != board[x][y] &&
				board[x+2][y] == board[x+1][y] &&
				board[x+3][y] == board[x][y] {
				fmt.Println("found capture of", board[x+1][y])
			}
			// from {x,y} to {x,y + 3}
			if y+3 < boardSize &&
				board[x][y] != 0 &&
				board[x][y+1] != 0 && board[x][y+1] != board[x][y] &&
				board[x][y+2] == board[x][y+1] &&
				board[x][y+3] == board[x][y] {
				fmt.Println("found capture of", board[x][y+1])
			}
			// from {x,y} to {x + 3,y + 3}
			if x+3 < boardSize && y+3 < boardSize &&
				board[x][y] != 0 &&
				board[x+1][y+1] != 0 && board[x+1][y+1] != board[x][y] &&
				board[x+2][y+2] == board[x+1][y+1] &&
				board[x+3][y+3] == board[x][y] {
				fmt.Println("found capture of", board[x+1][y+1])
			}
			// from {x,y} to {x - 3,y + 3}
			if x >= 3 && y+3 < boardSize &&
				board[x][y] != 0 &&
				board[x-1][y+1] != 0 && board[x-1][y+1] != board[x][y] &&
				board[x-2][y+2] == board[x-1][y+1] &&
				board[x-3][y+3] == board[x][y] {
				fmt.Println("found capture of", board[x-1][y+1])
			}
		}
	}
}
