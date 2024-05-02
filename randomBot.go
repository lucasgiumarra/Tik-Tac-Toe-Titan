package ttt

import (
	"math/rand"
)

type RandomBot struct {
	Symbol Symbol
	Index  int
	Id     int
}

func (r *RandomBot) selectMove(b *Board) Coordinate {
	candidates := b.GetLegalMoves()
	randIndex := rand.Intn(len(candidates))
	return candidates[randIndex]
}

func (r *RandomBot) getID() int {
	return r.Id
}

func (r *RandomBot) getSymbol() Symbol {
	return r.Symbol
}
