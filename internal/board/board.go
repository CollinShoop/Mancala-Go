package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	sides [2]*BoardSide
}

// Apply move, returns whether or not player got a mancala where player can make another move
func (b *Board) Move(player, n int) bool {
	var side = b.sides[player]

	// make move
	numStones, endIndex := side.StartMove(n)

	// check for case where player move ends on mancala
	if numStones == 0 {
		if endIndex == 6 {
			// mancala!
			return true
		} else {
			b.CheckAndCapture(player, endIndex)
		}
	}
	// TODO check for case where move ended on same-side tile that was previously empty
	// TODO in this case, everything in that tile and corresponding tile on the opponents side
	// TOD are moved to mancala

	// if there are remaining stones, continue distributing
	for numStones > 0 {
		// select the other side of the board
		player = (player + 1) % 2
		side = b.sides[player]
		numStones, _ = side.Move(0, numStones)
	}
	return false // no move-again
}

func (b *Board) CheckAndCapture(player, tile int) {
	// for capture to be valid, the tile must have size 1
	ps, po := b.sides[player], b.sides[(player+1)%2]
	otile := 5 - tile
	if ps.tile[tile] == 1 && po.tile[otile] > 0 {
		//captureTotal := ps.tile[tile] + po.tile[otile]
		//color.Green("Capturing opponents tile %v, +%v", otile, captureTotal)
		// move piece from ps[n] to ps.mancala
		ps.mancala += ps.tile[tile]
		ps.tile[tile] = 0

		// move piece from po[opposite to n] to ps.mancala
		ps.mancala += po.tile[otile]
		po.tile[otile] = 0
	}
}

// Get a list of valid moves that the given player can make. Player should be a value of 0 or 1 to
// indicate Player 1 and Player 2, respectively.
func (b *Board) GetPlayerMoves(player int) []int {
	if player < 0 || player > 2 {
		panic("Player should be 0 or 1")
	}
	return b.sides[player].GetMoves()
}

// Moves any remaining pieces to each players Mancala, ending the game
func (b *Board) Cleanup() {
	b.sides[0].Cleanup()
	b.sides[1].Cleanup()
}

func (b Board) GetWinner() int {
	mi, mv := -1, 0
	for i, v := range b.sides {
		if v.mancala > mv {
			mi, mv = i, v.mancala
		}
	}
	return mi
}

func (b Board) IsGameOver() bool {
	sum := 0
	for _, side := range b.sides {
		for _, n := range side.tile {
			sum += n
		}
	}
	return sum == 0
}

func (b Board) String() string {
	var boardf = "--------------------------------------------------------------\n" +
		"  Player 2           6     5     4     3     2     1                   \n" +
		"            ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐          \n" +
		"            │     │  ** │  ** │  ** │  ** │  ** │  ** │     │          \n" +
		"            │  $$ ├─────┼─────┼─────┼─────┼─────┼─────┤  $$ │          \n" +
		"            │     │  && │  && │  && │  && │  && │  && │     │          \n" +
		"            └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘          \n" +
		"                     1     2     3     4     5     6           Player 1  "

	// Write P1 values
	formatV := func(n int) string {
		ns := strconv.Itoa(n)
		if n < 10 {
			return ns + " "
		}
		return ns
	}
	for _, v := range b.sides[0].tile {
		boardf = strings.Replace(boardf, "&&", formatV(v), 1)
	}
	// Write P2 values backwards
	for i := range b.sides[1].tile {
		boardf = strings.Replace(boardf, "**", formatV(b.sides[1].tile[len(b.sides[1].tile)-i-1]), 1)
	}
	// write P2 mancala value
	boardf = strings.Replace(boardf, "$$", formatV(b.sides[1].mancala), 1)

	// write P1 mancala value
	boardf = strings.Replace(boardf, "$$", formatV(b.sides[0].mancala), 1)

	return fmt.Sprintf(boardf)
}

func (b *Board) SelectRandomMove(player int) int {
	// TODO add check to make sure moves are available and that move is selected
	return GetRandomMove(*b.sides[player])
}

func (b *Board) Copy() Board {
	board := Board{
		sides: [2]*BoardSide{},
	}
	board.sides[0] = &BoardSide{
		mancala: b.sides[0].mancala,
	}
	board.sides[1] = &BoardSide{
		mancala: b.sides[1].mancala,
	}
	var side0cpy, side1cpy [6]int
	for i, v := range b.sides[0].tile {
		side0cpy[i] = v
	}
	for i, v := range b.sides[1].tile {
		side1cpy[i] = v
	}
	board.sides[0].tile = side0cpy
	board.sides[1].tile = side1cpy
	return board
}

// Creates a new starting board with 2 players. Each player has 4 pieces in each tile and 0 pieces in the	 Mancala
func GetNewBoard() Board {
	board := Board{}
	board.sides[0] = &BoardSide{
		mancala: 0,
		tile:    [6]int{4, 4, 4, 4, 4, 4},
	}
	board.sides[1] = &BoardSide{
		mancala: 0,
		tile:    [6]int{4, 4, 4, 4, 4, 4},
	}
	return board
	// Output: A new starting board
}
