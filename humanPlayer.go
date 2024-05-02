package ttt

import "fmt"

var (
	MovesPlayer1 = [4][2]int{{1, 1}, {2, 0}, {2, 2}, {2, 1}}
	MovesPlayer2 = [4][2]int{{1, 2}, {0, 2}, {0, 0}, {0, 0}}
)

type HumanPlayer struct {
	Symbol Symbol
	Index  int
	Id     int
}

func (h *HumanPlayer) getSymbol() Symbol {
	return h.Symbol
}

func (h *HumanPlayer) getNextMove() (int, int, error) {
	if h.Symbol == Cross {
		h.Index = h.Index + 1
		return MovesPlayer1[h.Index-1][0], MovesPlayer1[h.Index-1][1], nil
	} else if h.Symbol == Circle {
		h.Index = h.Index + 1
		return MovesPlayer2[h.Index-1][0], MovesPlayer2[h.Index-1][1], nil
	}

	return 0, 0, fmt.Errorf("Invalid Symbol")
}

func (h *HumanPlayer) getID() int {
	return h.Id
}
