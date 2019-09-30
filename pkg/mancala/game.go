// Implements a playable game of Mancala using stdio.
// Game rules https://endlessgames.com/wp-content/uploads/Mancala_Instructions.pdf
package mancala

import (
	"../../internal/board"
	"fmt"
	"github.com/fatih/color"
)

// FPrint Plays a game of Mancala
// Variable inputType
// type Code_blocks int
// 	fmt.Println("Hello")
func PlayGame(inputType [2]int) {
	board := board.GetNewBoard()
	player := 0

	for !board.IsGameOver() {
		printBoard(&board)

		// check for case where no moves are available
		validMoves := board.GetPlayerMoves(player)

		var numMoves = len(validMoves)
		if numMoves == 0 {
			color.Red("No moves available, ending game")
			board.Cleanup()
			break
		}

		// get next move depending on the input type of the current player
		fmt.Printf("Player %v's turn\n", player+1)
		var moveInput int
		switch inputType[player] {
		case 0:
			moveInput = getNextMoveFromPlayer(&validMoves)
		case 1:
			moveInput = board.SelectRandomMove(player)
		case 2:
			moveInput = board.GetMiniMaxMove(player)
			//waitForContinue()
		}

		fmt.Println("Moving tile", moveInput)
		moveAgain := board.Move(player, moveInput)

		if moveAgain {
			color.Green("Move again!")
		} else {
			player = (player + 1) % 2
		}
	}

	printBoard(&board)
	color.Green("Player %v wins!", board.GetWinner()+1)
}

func printBoard(board *board.Board) {
	fmt.Println(board)
}

func waitForContinue() {
	fmt.Println("Press enter to continue...")
	fmt.Scanf("\r\n")
}

func getNextMoveFromPlayer(validMoves *[]int) int {
	for {
		fmt.Print("Enter move >")
		var moveInput int
		_, err := fmt.Scanf("%d\r\n", &moveInput)
		if err != nil {
			color.Red("Invalid move: %v, input move by typing the number of a tile that has more than 0 pieces", err)
		} else {
			if !contains(*validMoves, moveInput) {
				color.Red("Invalid move: tile is empty")
			} else {
				return moveInput
			}
		}
	}
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
