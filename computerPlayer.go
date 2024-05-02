package ttt

type computerPlayer struct {
	Symbol Symbol
	Id     int
}

func (c *computerPlayer) getSymbol() Symbol {
	return c.Symbol
}

func (c *computerPlayer) getNextMove() (int, int, error) {
	// To be implemented
	return 0, 0, nil
}

func (c *computerPlayer) getID() int {
	return c.Id
}
