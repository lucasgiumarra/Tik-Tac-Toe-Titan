package ttt

type GameStatus uint8

const (
	GameInProgress GameStatus = iota
	GameDraw
	FirstPlayerWin
	SecondPlayerWin
)
