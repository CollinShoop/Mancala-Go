package board

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var debug = false

type MoveResult struct {
	move, score int
}

// calculates best move using minimax algorithm
func (b *Board) GetMiniMaxMove(player int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	// find max score after applying all moves
	fmt.Println("Calculating move...")
	score, move := minimax(b.Copy(), player, player, 0, 5, true)
	fmt.Println("Minimax picked max move", move, "score", score)
	return move
}

// calculates score of the board using minimax algorithm
func minimax(board Board, player, scorePlayer, depth, maxDepth int, maximizing bool) (int, int) {
	moves := board.GetPlayerMoves(player)
	if debug {
		fmt.Println("minimax ---- player", player+1, "depth", depth, "maximizing", maximizing, "moves", moves)
	}

	// Check 2 leaf cases
	// Case 1: No moves available, in this case return score of cleaned board
	if len(moves) == 0 {
		cleanedUp := board.Copy()
		cleanedUp.Cleanup()
		if debug {
			fmt.Println("minimax -- no moves available")
		}
		return score(cleanedUp, scorePlayer), -1
	}
	// Case 2: Max depth reached, return score of board at current state
	if depth == maxDepth {
		if debug {
			fmt.Println("minimax -- max depth reached")
		}
		return score(board, scorePlayer), -1
	}

	// stores list of moves and their scores
	moveResults := make([]MoveResult, 0)

	// iterate over moves, calculate score for each move by advancing the tree
	for _, v := range moves {
		if debug {
			fmt.Println("minimax - making move player", player+1, "depth", depth, "maximizing", maximizing, "moves", moves, "move", v)
		}
		var tmpScore int
		tmpBoard := board.Copy()
		if tmpBoard.Move(player, v) {
			// move gives a mancala, move again without advancing depth
			tmpScore, _ = minimax(tmpBoard, player, scorePlayer, depth, maxDepth, maximizing)
		} else {
			// advance to the next players move
			tmpScore, _ = minimax(tmpBoard, (player+1)%2, scorePlayer, depth+1, maxDepth, !maximizing)
		}

		// save result
		moveResult := MoveResult{
			score: tmpScore,
			move:  v,
		}
		moveResults = append(moveResults, moveResult)
	}

	// picks and returns a random move from top moves
	pickedMove := pickRandomMove(moveResults, maximizing)
	return pickedMove.score, pickedMove.move
}

func pickRandomMove(moveResults []MoveResult, maximizing bool) MoveResult {
	// sorts moves by score,
	// order largest to smallest if maximizing
	// order smallest to largest if minimizing

	sort.SliceStable(moveResults[:], func(i, j int) bool {
		if maximizing {
			return moveResults[i].score > moveResults[j].score
		} else {
			return moveResults[i].score < moveResults[j].score
		}
	})

	// looks at the first value of moveResults & all subsequent moves of the same or very similar values
	// and picks a random one to determine which move will be made

	score := moveResults[0].score
	upper := 0
	for i, move := range moveResults {
		if move.score == score {
			upper = i
		}
	}

	return moveResults[rand.Intn(upper+1)]
}

// calculates score of the given board from the players perspective
func score(board Board, player int) int {
	return board.sides[player].mancala - board.sides[(player+1)%2].mancala
}
