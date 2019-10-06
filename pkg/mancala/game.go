// Implements a playable game of Mancala using stdio.
// Game rules https://endlessgames.com/wp-content/uploads/Mancala_Instructions.pdf
package mancala

import (
	"fmt"
	"mancala/internal/board"
	"mancala/internal/util"

	"github.com/fatih/color"
)

// Starts game.
// Uses STDIO console for input and output.
// inputType gives the type of player for players 1 and 2, with 0 meaning the
// player inputs move from console, 1 being that a random move is always selected, and 2
// being that a minimax strategy is used to pick the next move. The game will end
// once there are no more moves and the winner will be announced.
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
		case 3:
			moveInput = board.GetMiniMaxMove_mt(player)
			//waitForContinue()
		}

		fmt.Println("Moving tile", moveInput+1)
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

// Print board to STDIO, takes advantage of board.String
func printBoard(board *board.Board) {
	fmt.Println(board)
}

// Gets move input from the player via STDIO
func getNextMoveFromPlayer(validMoves *[]int) int {
	for {
		fmt.Print("Enter move >")
		var moveInput int
		_, err := fmt.Scanf("%d\r\n", &moveInput)
		moveInput--
		if err != nil {
			color.Red("Invalid move: %v, input move by typing the number of a tile that has more than 0 pieces", err)
		} else {
			if !util.Contains(*validMoves, moveInput) {
				color.Red("Invalid move: tile is empty")
			} else {
				return moveInput
			}
		}
	}
}
