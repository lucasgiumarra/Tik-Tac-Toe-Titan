package ttt

import (
	"fmt"
	"math/rand"
	"os"
)

type TwoLayerBot struct {
	Symbol Symbol
	Index  int
	Id     int
}

func (t *TwoLayerBot) getID() int {
	return t.Id
}

func (t *TwoLayerBot) getSymbol() Symbol {
	return t.Symbol
}

func (t *TwoLayerBot) selectMove(b *Board) Coordinate {
	var losingMove Coordinate
	var newboard Board
	var nb Board
	losingBool := false
	firstCandidates := b.GetLegalMoves()

	for i, val := range firstCandidates {
		// Check for winning move
		row := val.x
		col := val.y
		if i == 0 {
			newboard = b.CopyBoard()
		} else {
			newboard = newboard.CopyBoard()
		}
		fmt.Fprintf(os.Stderr, "\tVal x: %v\n", val.x)
		fmt.Fprintf(os.Stderr, "\tVal y: %v\n", val.y)
		newboard.MakeMove(row, col)
		newboard.markSymbol(row, col, t.Symbol)
		fmt.Fprintf(os.Stderr, "\tNewboard: \n")
		newboard.PrintTabBoard()
		winner, isWinner := newboard.WhoWinner()
		fmt.Fprintf(os.Stderr, "\tVal: %v\n", val)
		fmt.Fprintf(os.Stderr, "\twinner: %v\n\n", winner)
		if winner == t.getSymbol() && isWinner {
			return Coordinate{x: row, y: col}
		}

		secondCandidates := newboard.GetLegalMoves()
		for _, oppMove := range secondCandidates {
			// Check if opponent can win the game on their next turn
			x := oppMove.x
			y := oppMove.y
			if i == 0 {
				nb = b.CopyBoard()
			} else {
				nb = nb.CopyBoard()
			}

			nb.MakeMove(row, col)
			// fmt.Fprintf(os.Stderr, "\tOpponent Board: \n")
			nb.PrintTabBoard()
			win, isWin := nb.WhoWinner()
			// fmt.Fprintf(os.Stderr, "\toppVal: %v\n", oppMove)
			// fmt.Fprintf(os.Stderr, "\toppWinner: %v\n\n", win)
			if win == t.getSymbol() && isWin {
				losingMove = Coordinate{x: x, y: y}
				losingBool = true
				break
			}
		}
		// if losingBool {
		// 	break // Break out of the outer loop if we found a losing move
		// }
	}
	// If opponent has a winning move, block it
	if losingBool {
		fmt.Fprintf(os.Stderr, "\tlosingMove: %v\n", losingMove)
		return losingMove
	}

	// Otherwise return a random choice
	randInt := rand.Intn(len(firstCandidates))
	return firstCandidates[randInt]
}
