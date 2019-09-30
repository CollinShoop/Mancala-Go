package board

import (
	"math/rand"
	"time"
)

var seeded = false

func GetRandomMove(side BoardSide) int {
	if !seeded {
		rand.Seed(time.Now().UTC().UnixNano())
		seeded = true
	}
	moves := side.GetMoves()
	return moves[rand.Intn(len(moves))]
}
