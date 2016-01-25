package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// The top row is represented by indexes 0, 1, 2, the middle row by 3, 4, 5, and the bottom row by 6, 7, 8.
	// As players make their moves, give the appropriate indexes the values "X" and "O".
	// For example, when player X selects the top-left corner:
	//  	board[0] = "X"
	board := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// init the reader
	stdin := bufio.NewReader(os.Stdin)

	player_won := false
	player := "o"
	turns := 0
	for !player_won {
		turns += 1

		// Before each turn, display the current state of the board with printBoard()
		printBoard(board)

		// Prompt player to make a move by entering a number 1 through 9
		fmt.Print("Make a move by entering a number 1 through 9: ")
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error when reading from STDIN")
		}

		// Clean input
		input = strings.TrimSpace(input)

		// Check valid input
		if !validMove(input) {
			fmt.Println("Error: invalid number!")
			continue
		}

		// Player move
		index, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		index -= 1
		if board[index] == "x" || board[index] == "o" {
			fmt.Println("Spot already taken, try a different number")
			continue
		} else {
			board[index] = player
		}

		// The game need at least 3 moves to have a winner
		if turns > 2 {
			// After player makes a move, check if a player has won or if the game is a tied stalemate.
			winner := winner(&board)
			if winner != "" {
				fmt.Println("Winner: ", winner)
				return
			}

			if turns > 7 {
				fmt.Println("TIE!")
				return
			}
		}

		// Update player
		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}
	}
}

// Display board
func printBoard(board [9]string) {
	fmt.Println("")
	fmt.Println("    ", board[0], board[1], board[2])
	fmt.Println("    ", board[3], board[4], board[5])
	fmt.Println("    ", board[6], board[7], board[8])
	fmt.Println("")
}

// Check if move is valid
func validMove(input string) bool {
	board_index, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return board_index > 0 && board_index < 10
}

// return "X" or "O" if x or o player wins
// return "tie" if no squares are empty and there is no winner
// return "" if some squares are still empty and there is no winner
// (Bonus objective: return "tie" if the board still has empty squares but neither player can win)
func winner(board *[9]string) string {
	if xWon(board) {
		return "X"
	} else if oWon(board) {
		return "O"
	} else {
		return ""
	}
}

// Check if "x" won
func xWon(board *[9]string) bool {
	return won(board, "x")
}

// Check if "o" won
func oWon(board *[9]string) bool {
	return won(board, "o")
}

// Check if player won
func won(board *[9]string, player string) bool {
	return (board[0] == player && board[3] == player && board[6] == player) ||
		(board[1] == player && board[4] == player && board[7] == player) ||
		(board[2] == player && board[5] == player && board[8] == player) ||
		(board[2] == player && board[4] == player && board[6] == player) ||
		(board[0] == player && board[4] == player && board[8] == player)
}
