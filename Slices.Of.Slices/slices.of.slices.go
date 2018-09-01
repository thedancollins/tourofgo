package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	var s string
	s = ""
	for i := 0; i < len(board); i++ {
		if i == 0 {
			s = strings.Repeat("-", (len(board[i])*2)+1)
		}

		fmt.Printf("%s\n", strings.Repeat("-", (len(board[i])*2)+1))
		fmt.Printf("|%s|\n", strings.Join(board[i], " "))
	}
	fmt.Printf("%s\n", s)
}
