package ttt

import (
// "fmt"
// "os"
)

type Choice struct {
	Move  Coordinate
	Value int
	Depth int
}

type InvinciBot struct {
	Symbol Symbol
	Index  int
	Id     int
}

func (i *InvinciBot) getID() int {
	return i.Id
}

func (i *InvinciBot) getSymbol() Symbol {
	return i.Symbol
}

func (i *InvinciBot) Minimax(b *Board, is_max bool, currentPlayer Symbol, depth int) Choice {
	// Base Cases
	winner, isWinner := b.WhoWinner()
	// fmt.Fprintf(os.Stderr, "Winner: %v\n", winner)
	// fmt.Fprintf(os.Stderr, "isWinner: %v\n", isWinner)
	// fmt.Fprintf(os.Stderr, "currentPlayer: %v\n", currentPlayer)
	if winner == currentPlayer {
		return Choice{Value: 10 - depth}
	} else if winner == currentPlayer.Other() {
		return Choice{Value: -depth + 10}
	} else if winner == Dot && isWinner {
		return Choice{Value: 0}
	}

	candidate_choices := []Choice{}
	candidates := b.GetLegalMoves()

	for _, val := range candidates {
		row := val.x
		col := val.y
		newboard := b.CopyBoard()
		newboard.MakeMove(row, col)
		newboard.markSymbol(row, col, currentPlayer)
		result := i.Minimax(&newboard, !is_max, currentPlayer.Other(), depth+1)
		result.Move = val
		candidate_choices = append(candidate_choices, result)
	}

	// Select best move for maximizing player
	if is_max {
		max_choice := Choice{Value: -100}
		for _, choice := range candidate_choices {
			if choice.Value > max_choice.Value {
				max_choice = choice
			}
		}
		return max_choice
	}

	// Select best move for minimizing player
	min_choice := Choice{Value: 100}
	for _, choice := range candidate_choices {
		if choice.Value < min_choice.Value {
			min_choice = choice
		}
	}
	return min_choice
}

//	func (i *InvinciBot) Minimax(b *Board, is_max bool, currentPlayer Symbol, depth int) Choice {
//		// If board has a winner or is a tie return with appropriate values
//		_, isWinner := b.WhoWinner()
//		if currentPlayer == Cross && isWinner {
//			lastVal, err := b.LastMove()
//			if err == nil {
//				return Choice{Move: lastVal, Value: 10 - depth, Depth: depth}
//			} else {
//				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
//			}
//		} else if currentPlayer == Circle && isWinner {
//			lastVal, err := b.LastMove()
//			if err == nil {
//				return Choice{Move: lastVal, Value: -depth + 10, Depth: depth}
//			} else {
//				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
//			}
//		}
//		if len(b.Moves) == 9 {
//			lastVal, err := b.LastMove()
//			if err == nil {
//				return Choice{Move: lastVal, Value: 0, Depth: depth}
//			} else {
//				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
//			}
//		}
//
//		candidate_choices := []Choice{}
//		candidates := b.GetLegalMoves()
//		// fmt.Fprintf(os.Stderr, "Candidates: %v\n", candidates)
//		for _, val := range candidates {
//			row := val.x
//			col := val.y
//			newboard := b.CopyBoard()
//			newboard.MakeMove(row, col)
//			result := i.Minimax(&newboard, !is_max, currentPlayer.Other(), depth+1)
//			lastVal, err := newboard.LastMove()
//			result.Move = lastVal
//			if err == nil {
//				candidate_choices = append(candidate_choices, result)
//			}
//		}
//		fmt.Fprintf(os.Stderr, "Candidates choices: %v\n\n", candidate_choices)
//		max_choice := Choice{}
//		max_value := -100
//		min_choice := Choice{}
//		min_value := 100
//
//		// determine which board combinations result in the best move for a paticular agent
//		for i := range candidate_choices {
//			choice := candidate_choices[i]
//			if is_max && choice.Value > max_value {
//				max_choice = Choice{Move: choice.Move, Value: choice.Value, Depth: choice.Depth}
//				max_value = choice.Value
//			} else if !is_max && choice.Value < min_value {
//				min_choice = Choice{Move: choice.Move, Value: choice.Value, Depth: choice.Depth}
//				min_value = choice.Value
//			}
//		}
//
//		fmt.Fprintf(os.Stderr, "min choices: %v\n", min_choice)
//		fmt.Fprintf(os.Stderr, "max choices: %v\n", max_choice)
//		// fmt.Fprintf(os.Stderr, "max value: %v\n", max_value)
//		// fmt.Fprintf(os.Stderr, "max value: %v\n\n", max_va
//
//		// Pick whichever move is the best for the patiuclar agent
//		if is_max {
//			return max_choice
//		}
//
//		return min_choice
//	}
func (b *Board) WhoWinner() (Symbol, bool) {
	// Check rows
	for row := 0; row < 3; row++ {
		if !b.isSpaceEmpty(row, 0) && b.Square[row][0] == b.Square[row][1] && b.Square[row][0] == b.Square[row][2] {
			return b.Square[row][0], true
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		if !b.isSpaceEmpty(0, col) && b.Square[0][col] == b.Square[1][col] && b.Square[0][col] == b.Square[2][col] {
			return b.Square[0][col], true
		}
	}

	// Check diagonals
	if !b.isSpaceEmpty(0, 0) && b.Square[0][0] == b.Square[1][1] && b.Square[0][0] == b.Square[2][2] {
		return b.Square[0][0], true
	}
	if !b.isSpaceEmpty(0, 2) && b.Square[0][2] == b.Square[1][1] && b.Square[0][2] == b.Square[2][0] {
		return b.Square[0][2], true
	}

	// Check if the board is full (tie)
	if b.isFull() {
		return Dot, true
	}

	// No winner yet
	return Dot, false
}

func (b *Board) isFull() bool {
	for row := 0; row < b.Dimension; row++ {
		for col := 0; col < b.Dimension; col++ {
			if b.Square[row][col] == Dot {
				// If any cell is empty, the board is not full
				return false
			}
		}
	}
	// If no cell is empty, the board is full
	return true
}

// func (b *Board) WhoWinner() (symbol Symbol, isWinner bool) {
// 	winner := false
// 	for i := 0; i < b.Dimension; i++ {
// 		for j := 0; j < b.Dimension; j++ {
// 			winner = b.checkWin(i, j, b.Square[i][j])
// 			if winner {
// 				fmt.Fprintf(os.Stderr, "A winning opprotunity is on the board\n")
// 				return b.Square[i][j], true
// 			}
// 		}
// 	}
// 	return Dot, false
// }

func (i *InvinciBot) selectMove(b *Board) Coordinate {
	choice := i.Minimax(b, true, i.Symbol, 0)
	return choice.Move
}
