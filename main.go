package main

import "ttt"

func main() {
	board := &ttt.Board{
		Square:    [][]ttt.Symbol{{ttt.Dot, ttt.Dot, ttt.Dot}, {ttt.Dot, ttt.Dot, ttt.Dot}, {ttt.Dot, ttt.Dot, ttt.Dot}},
		Dimension: 3,
	}

	player1 := &ttt.RandomBot{
		Symbol: ttt.Cross,
		Id:     1,
	}

	player2 := &ttt.InvinciBot{
		Symbol: ttt.Circle,
		Id:     2,
	}

	game := ttt.InitGame(board, player1, player2)
	game.Play(board)
	game.PrintResult()
}
