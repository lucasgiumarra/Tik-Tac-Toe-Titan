package ttt

type iPlayer interface {
	getSymbol() Symbol
	selectMove(b *Board) Coordinate
	getID() int
}
