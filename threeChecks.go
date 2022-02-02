package main

import (
	"fmt"
	"reflect"
)

// returns the coordinates of free threes found
func horizontalThree(x int, y int, freeThrees [][][]int) [][][]int {
	boardSize := getBoardsize()

	if x+5 <= boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == board[x+1][y] &&
			board[x+3][y] == board[x+1][y] &&
			board[x+4][y] == 0 {
			fmt.Println("found -***- horizontal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 4, y}, {1}})
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == 0 &&
			board[x+3][y] == board[x+1][y] &&
			board[x+4][y] == board[x+1][y] &&
			board[x+5][y] == 0 {
			fmt.Println("found -*-**- horizontal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 5, y}, {1}})
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x+1][y] != 0 &&
			board[x+2][y] == board[x+1][y] &&
			board[x+3][y] == 0 &&
			board[x+4][y] == board[x+1][y] &&
			board[x+5][y] == 0 {
			fmt.Println("found -**-*- horizontal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 5, y}, {1}})
		}
	}
	return freeThrees
}

func verticalThree(x int, y int, freeThrees [][][]int) [][][]int {
	boardSize := getBoardsize()

	if y+5 < boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == board[x][y+1] &&
			board[x][y+3] == board[x][y+1] &&
			board[x][y+4] == 0 {
			fmt.Println("found -***- vertical")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x, y + 4}, {2}})
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == 0 &&
			board[x][y+3] == board[x][y+1] &&
			board[x][y+4] == board[x][y+1] &&
			board[x][y+5] == 0 {
			fmt.Println("found -*-**- vertical")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x, y + 5}, {2}})
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x][y+1] != 0 &&
			board[x][y+2] == board[x][y+1] &&
			board[x][y+3] == 0 &&
			board[x][y+4] == board[x][y+1] &&
			board[x][y+5] == 0 {
			fmt.Println("found -**-*- vertical")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x, y + 5}, {2}})
		}
	}
	return freeThrees
}

func diagonalRightThree(x int, y int, freeThrees [][][]int) [][][]int {
	boardSize := getBoardsize()

	if y+5 < boardSize && x+6 < boardSize {
		// -***-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == board[x+1][y+1] &&
			board[x+3][y+3] == board[x+1][y+1] &&
			board[x+4][y+4] == 0 {
			fmt.Println("found -***- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 4, y + 4}, {3}})
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == 0 &&
			board[x+3][y+3] == board[x+1][y+1] &&
			board[x+4][y+4] == board[x+1][y+1] &&
			board[x+5][y+5] == 0 {
			fmt.Println("found -*-**- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 5, y + 5}, {3}})
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x+1][y+1] != 0 &&
			board[x+2][y+2] == board[x+1][y+1] &&
			board[x+3][y+3] == 0 &&
			board[x+4][y+4] == board[x+1][y+1] &&
			board[x+5][y+5] == 0 {
			fmt.Println("found -**-*- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x + 5, y + 5}, {3}})
		}
	}
	return freeThrees
}

func diagonalLeftThree(x int, y int, freeThrees [][][]int) [][][]int {
	boardSize := getBoardsize()

	if y+5 < boardSize && x > 5 {
		// -***-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == board[x-1][y+1] &&
			board[x-3][y+3] == board[x-1][y+1] &&
			board[x-4][y+4] == 0 {
			fmt.Println("found -***- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x - 4, y + 4}, {3}})
		}
		// -*-**-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == 0 &&
			board[x-3][y+3] == board[x-1][y+1] &&
			board[x-4][y+4] == board[x-1][y+1] &&
			board[x-5][y+5] == 0 {
			fmt.Println("found -*-**- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x - 5, y + 5}, {3}})
		}
		// -**-*-
		if board[x][y] == 0 &&
			board[x-1][y+1] != 0 &&
			board[x-2][y+2] == board[x-1][y+1] &&
			board[x-3][y+3] == 0 &&
			board[x-4][y+4] == board[x-1][y+1] &&
			board[x-5][y+5] == 0 {
			fmt.Println("found -**-*- diagonal")
			freeThrees = append(freeThrees, [][]int{{x, y}, {x - 5, y + 5}, {3}})
		}
	}
	return freeThrees
}

func checkFreeThrees(existingFreeThrees [][][]int) [][][]int {
	boardSize := getBoardsize()

	freeThrees := make([][][]int, 0)
	newFreeThrees := make([][][]int, 0)
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			freeThrees = horizontalThree(x, y, freeThrees)
			freeThrees = verticalThree(x, y, freeThrees)
			freeThrees = diagonalLeftThree(x, y, freeThrees)
			freeThrees = diagonalRightThree(x, y, freeThrees)
		}
	}
	for j := 0; j < len(freeThrees); j++ {
		exists := false
		for i := 0; i < len(existingFreeThrees); i++ {
			if reflect.DeepEqual(existingFreeThrees[i], freeThrees[j]) {
				exists = true
			}
		}
		if exists == false {
			newFreeThrees = append(newFreeThrees, freeThrees[j])
		}
	}
	if len(newFreeThrees) > 1 {
		fmt.Println("ILLEGAL MOVE")
	}
	return freeThrees
}
