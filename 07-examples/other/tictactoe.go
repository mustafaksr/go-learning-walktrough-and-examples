package main

import (
	"fmt"
)

func printBoard(board [3][3]string) {
	for i, row := range board {
		for j, cell := range row {
			if cell == "" {
				fmt.Printf(" %d ", i*3+j+1)
			} else {
				fmt.Printf(" %s ", cell)
			}

			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()

		if i < 2 {
			fmt.Println("-----------")
		}
	}
}

func checkWin(board [3][3]string, player string) bool {
	// Check rows, columns, and diagonals for a win
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func main() {
	var board [3][3]string
	currentPlayer := "X"
	var move int

	fmt.Println("Welcome to Tic-Tac-Toe!")

	for {
		fmt.Printf("\nPlayer %s's turn:\n", currentPlayer)
		printBoard(board)

		fmt.Print("Enter your move (1-9): ")
		fmt.Scan(&move)

		if move < 1 || move > 9 {
			fmt.Println("Invalid move. Please enter a number between 1 and 9.")
			continue
		}

		row := (move - 1) / 3
		col := (move - 1) % 3

		if board[row][col] != "" {
			fmt.Println("Invalid move. Cell already occupied. Try again.")
			continue
		}

		board[row][col] = currentPlayer

		if checkWin(board, currentPlayer) {
			printBoard(board)
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		// Check for a draw
		isDraw := true
		for _, row := range board {
			for _, cell := range row {
				if cell == "" {
					isDraw = false
					break
				}
			}
			if !isDraw {
				break
			}
		}
		if isDraw {
			printBoard(board)
			fmt.Println("It's a draw!")
			break
		}

		// Switch players
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}

	fmt.Println("Game over.")
}
