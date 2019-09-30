package board

type BoardSide struct {
	mancala int
	tile    [6]int
}

func (b *BoardSide) Cleanup() {
	for i := 0; i < 6; i++ {
		b.mancala += b.tile[i]
		b.tile[i] = 0
	}
}

func (b *BoardSide) GetMoves() []int {
	moves := make([]int, 0)
	for i := 0; i < 6; i++ {
		if b.IsValidMove(i) {
			moves = append(moves, i)
		}
	}
	return moves
}

func (b *BoardSide) IsValidMove(n int) bool {
	if n < 0 || n > 5 {
		return false
	}
	return b.tile[n] > 0
}

func (b *BoardSide) StartMove(n int) (int, int) {
	var stones = b.tile[n]
	b.tile[n] = 0
	return b.Move(n+1, stones)
}

/*
	Starts a new move at tile by number from 0 to 5

	Returns the number of remaining stones after applying move to this side of the board,
	also returns the index that the last move was on, 0 to 6. values 0 to 5 indicate the move ended
	on a tile with 6 indicating that the move ended on the mancala.
*/
func (b *BoardSide) Move(tile, numStones int) (int, int) {
	var stop int

	// case where there are no stones, should never happen as this is not a valid move.
	if numStones == 0 {
		return 0, tile
	}
	for cellI := tile; cellI < 6 && numStones > 0; cellI++ {
		b.tile[cellI]++
		numStones--
		stop = cellI
	}
	if numStones > 0 {
		b.mancala++
		numStones--
		stop = 6
	}
	return numStones, stop
}
