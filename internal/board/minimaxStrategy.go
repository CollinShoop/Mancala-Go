package board

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var debug = false
var MAX_DEPTH = 7

type MoveResult struct {
	move, score int
}

// calculates best move using minimax algorithm
func (b *Board) GetMiniMaxMove(player int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	// find max score after applying all moves
	fmt.Println("Calculating move...")
	score, move := minimax(b.Copy(), player, 0, MAX_DEPTH)
	fmt.Println("Minimax picked max move", move, "score", score)
	return move
}

// calculates score of the board using minimax algorithm
func minimax(board Board, player, depth, maxDepth int) (int, int) {
	moves := board.GetPlayerMoves(player)
	if debug {
		fmt.Println("minimax ---- player", player+1, "depth", depth, "moves", moves)
	}

	// Check 2 leaf cases
	// Case 1: No moves available, in this case return score of cleaned board
	if len(moves) == 0 {
		cleanedUp := board.Copy()
		cleanedUp.Cleanup()
		if debug {
			fmt.Println("minimax -- no moves available")
		}
		return score(cleanedUp, player), -1
	}
	// Case 2: Max depth reached, return score of board at current state
	if depth == maxDepth {
		if debug {
			fmt.Println("minimax -- max depth reached")
		}
		return score(board, player), -1
	}

	// stores list of moves and their scores
	moveResults := make([]MoveResult, 0)

	// iterate over moves, calculate score for each move by advancing the tree
	for _, v := range moves {
		if debug {
			fmt.Println("minimax - making move player", player+1, "depth", depth, "moves", moves, "move", v)
		}
		var tmpScore int
		tmpBoard := board.Copy()
		if tmpBoard.Move(player, v) {
			// move gives a mancala, move again without advancing depth
			tmpScore, _ = minimax(tmpBoard, player, depth, maxDepth)
		} else {
			// advance to the next players move
			tmpScore, _ = minimax(tmpBoard, (player+1)%2, depth+1, maxDepth)
		}

		// save result
		moveResult := MoveResult{
			score: tmpScore,
			move:  v,
		}
		moveResults = append(moveResults, moveResult)
	}

	// picks and returns a random move from top moves
	var pickedMove MoveResult
	if depth == 0 {
		pickedMove = pickRandomMove(moveResults)
	} else {
		pickedMove = moveResults[0]
	}

	return pickedMove.score, pickedMove.move
}

func pickRandomMove(moveResults []MoveResult) MoveResult {
	// sorts moves by score,
	// order largest to smallest if maximizing
	sort.SliceStable(moveResults[:], func(i, j int) bool {
		return moveResults[i].score > moveResults[j].score
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

// calculates best move using minimax algorithm
func (b *Board) GetMiniMaxMove_mt(player int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	// find max score after applying all moves
	fmt.Println("Calculating move...")
	score, move := minimax_mt(b.Copy(), player, 0, MAX_DEPTH)
	fmt.Println("Minimax picked max move", move, "score", score)
	return move
}

// calculates score of the board using minimax algorithm
func minimax_mt(board Board, player, depth, maxDepth int) (int, int) {
	moves := board.GetPlayerMoves(player)
	if debug {
		fmt.Println("minimax ---- player", player+1, "depth", depth, "moves", moves)
	}

	// Check 2 leaf cases
	// Case 1: No moves available, in this case return score of cleaned board
	if len(moves) == 0 {
		cleanedUp := board.Copy()
		cleanedUp.Cleanup()
		if debug {
			fmt.Println("minimax -- no moves available")
		}
		return score(cleanedUp, player), -1
	}
	// Case 2: Max depth reached, return score of board at current state
	if depth == maxDepth {
		if debug {
			fmt.Println("minimax -- max depth reached")
		}
		return score(board, player), -1
	}

	// stores list of moves and their scores
	moveResults := make([]MoveResult, 0)

	// iterate over moves, calculate score for each move by advancing the tree
	var wg sync.WaitGroup
	var mrm sync.Mutex
	for _, v := range moves {
		if debug {
			fmt.Println("minimax - making move player", player+1, "depth", depth, "moves", moves, "move", v)
		}

		tmpBoard := board.Copy()
		tmpMove := v

		wg.Add(1)
		go func() {
			defer wg.Done()
			var tmpScore int
			if tmpBoard.Move(player, tmpMove) {
				// move gives a mancala, move again without advancing depth
				if depth < 1 {
					tmpScore, _ = minimax_mt(tmpBoard, player, depth, maxDepth)
				} else {
					tmpScore, _ = minimax(tmpBoard, player, depth, maxDepth)
				}
			} else {
				// advance to the next players move
				if depth < 1 {
					tmpScore, _ = minimax_mt(tmpBoard, (player+1)%2, depth+1, maxDepth)
				} else {
					tmpScore, _ = minimax(tmpBoard, (player+1)%2, depth+1, maxDepth)
				}
			}

			// save result
			moveResult := MoveResult{
				score: tmpScore,
				move:  tmpMove,
			}
			mrm.Lock()
			moveResults = append(moveResults, moveResult)
			mrm.Unlock()
		}()
	}
	wg.Wait()

	// picks and returns a random move from top moves
	var pickedMove MoveResult
	if depth == 0 {
		pickedMove = pickRandomMove(moveResults)
	} else {
		pickedMove = moveResults[0]
	}

	return pickedMove.score, pickedMove.move
}

// calculates score of the given board from the players perspective
func score(board Board, player int) int {
	return board.sides[player].mancala - board.sides[(player+1)%2].mancala
}
