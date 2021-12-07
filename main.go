package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/webview/webview"
)

var board [][]int

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

//create an ui with webview
func initWindow() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Gomoku")
	w.SetSize(800, 1250, webview.HintNone)
	w.Bind("getBoard", func() [][]int {
		return board
	})
	w.Bind("makeMove", makeMove)
	w.Bind("checkWin", checkWin)
	print(getCurrentDirectory())
	w.Navigate("file://" + getCurrentDirectory() + "/web/index.html")
	w.Run()
}

func makeMove(col int, row int, player int) [][]int {
	if board[row][col] == 0 {
		board[row][col] = player
		return [][]int{{col, row}, {col + 1, row}}
	}
	return nil
}

// minimax algorithm
// maybe add alpha-beta pruning?
func minimax(board [][]int, depth int, isMax bool) int {
	if depth == 0 {
		return 0
	}
	if isMax {
		var best int = -1
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 0 {
					board[i][j] = 2
					best = max(best, minimax(board, depth-1, !isMax))
					board[i][j] = 0
				}
			}
		}
		return best
	} else {
		var best int = 1
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 0 {
					board[i][j] = 1
					best = min(best, minimax(board, depth-1, !isMax))
					board[i][j] = 0
				}
			}
		}
		return best
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getBoardsize() int {
	return 19
}

func main() {
	for i := 0; i < getBoardsize(); i++ {
		board = append(board, make([]int, getBoardsize()))
	}
	initWindow()
}
