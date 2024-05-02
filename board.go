package ttt

import "fmt"

type Coordinate struct {
	x int
	y int
}

type Board struct {
	Square    [][]Symbol
	Dimension int
	Moves     []Coordinate
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.Dimension; i++ {
		for j := 0; j < b.Dimension; j++ {
			if b.Square[i][j] == Dot {
				fmt.Print(".")
			} else if b.Square[i][j] == Cross {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println("")
	}
}

func (b *Board) PrintTabBoard() {
	for i := 0; i < b.Dimension; i++ {
		fmt.Print("\t")
		for j := 0; j < b.Dimension; j++ {
			if b.Square[i][j] == Dot {
				fmt.Print(".")
			} else if b.Square[i][j] == Cross {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println("")
	}
}

func (b *Board) markSymbol(i, j int, symbol Symbol) (bool, Symbol, error) {
	if i > b.Dimension || j > b.Dimension {
		return false, Dot, fmt.Errorf("index input is greater than Dimension")
	}
	if b.Square[i][j] != Dot {
		return false, Dot, fmt.Errorf("input Square alrady marked already")
	}
	if symbol != Cross && symbol != Circle {
		return false, Dot, fmt.Errorf("incorrect Symbol")
	}
	b.Square[i][j] = symbol
	win := b.checkWin(i, j, symbol)

	return win, symbol, nil
}

func (b *Board) checkWin(i, j int, symbol Symbol) bool {
	// Check Row
	rowMatch := true
	for k := 0; k < b.Dimension; k++ {
		if b.Square[i][k] != symbol {
			rowMatch = false
		}
	}

	if rowMatch {
		return rowMatch
	}

	// Check Row
	columnMatch := true
	for k := 0; k < b.Dimension; k++ {
		if b.Square[k][j] != symbol {
			columnMatch = false
		}
	}

	if columnMatch {
		return columnMatch
	}

	// Check diagonal
	diagonalMatch := false
	if i == j {
		diagonalMatch = true
		for k := 0; k < b.Dimension; k++ {
			if b.Square[k][k] != symbol {
				diagonalMatch = false
			}
		}
	}

	return diagonalMatch
}

func (b *Board) isSpaceEmpty(row, col int) bool {
	return b.Square[row][col] == Dot
}

func (b *Board) MakeMove(row, col int) {
	if b.isSpaceEmpty(row, col) {
		b.markSymbol(row, col, b.Square[row][col])
		b.AddMove(row, col)
	} else {
		fmt.Println("Error: Attempting to make a move to a space that is occupied")
	}
}

// Example function to add a move to the board
func (b *Board) AddMove(row, col int) {
	move := Coordinate{x: row, y: col}
	b.Moves = append(b.Moves, move)
}

func (b *Board) LastMove() (Coordinate, error) {
	if len(b.Moves) == 0 {
		return Coordinate{}, fmt.Errorf("no moves have been made")
	}
	return b.Moves[len(b.Moves)-1], nil
}

func (b *Board) GetLegalMoves() []Coordinate {
	choices := []Coordinate{}

	for i := 0; i < b.Dimension; i++ {
		for j := 0; j < b.Dimension; j++ {
			if b.isSpaceEmpty(i, j) {
				choices = append(choices, Coordinate{x: i, y: j})
			}
		}
	}
	return choices
}

func (b *Board) CopyBoard() Board {
	// Create a new Board instance with the same Dimension
	newBoard := Board{
		Dimension: b.Dimension,
	}

	// Copy the Square
	newBoard.Square = make([][]Symbol, b.Dimension)
	for i := range b.Square {
		newBoard.Square[i] = make([]Symbol, len(b.Square[i]))
		copy(newBoard.Square[i], b.Square[i])
	}

	// Copy the Moves
	newBoard.Moves = make([]Coordinate, len(b.Moves))
	copy(newBoard.Moves, b.Moves)

	return newBoard
}
