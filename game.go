package ttt

import "fmt"

type Game struct {
	board           *Board
	firstPlayer     iPlayer
	secondPlayer    iPlayer
	firstPlayerTurn bool
	moveIndex       int
	gameStatus      GameStatus
}

func InitGame(b *Board, p1, p2 iPlayer) *Game {
	game := &Game{
		board:           b,
		firstPlayer:     p1,
		secondPlayer:    p2,
		firstPlayerTurn: true,
		gameStatus:      GameInProgress,
	}
	return game
}

func (g *Game) Play(b *Board) error {
	var win bool
	var symbol Symbol
	var err error
	for {
		if g.firstPlayerTurn {
			firstPlayerCoordinates := g.firstPlayer.selectMove(b)
			win, symbol, err = g.board.markSymbol(firstPlayerCoordinates.x, firstPlayerCoordinates.y, g.firstPlayer.getSymbol())
			if err != nil {
				return err
			}
			g.firstPlayerTurn = false
			g.printMove(g.firstPlayer, firstPlayerCoordinates.x, firstPlayerCoordinates.y)
		} else {
			secondPlayerCoordinates := g.secondPlayer.selectMove(b)
			win, symbol, err = g.board.markSymbol(secondPlayerCoordinates.x, secondPlayerCoordinates.y, g.secondPlayer.getSymbol())
			if err != nil {
				return err
			}
			g.firstPlayerTurn = true
			g.printMove(g.secondPlayer, secondPlayerCoordinates.x, secondPlayerCoordinates.y)
		}
		g.moveIndex = g.moveIndex + 1
		g.setGameStatus(win, symbol)
		if g.gameStatus != GameInProgress {
			break
		}
	}
	return nil
}

func (g *Game) setGameStatus(win bool, symbol Symbol) {
	if win {
		if g.firstPlayer.getSymbol() == symbol {
			g.gameStatus = FirstPlayerWin
			return
		} else if g.secondPlayer.getSymbol() == symbol {
			g.gameStatus = SecondPlayerWin
			return
		}
	}
	if g.moveIndex == g.board.Dimension*g.board.Dimension {
		g.gameStatus = GameDraw
		return
	}
	g.gameStatus = GameInProgress

}

func (g *Game) printMove(player iPlayer, x, y int) {
	symbolString := ""
	symbol := player.getSymbol()
	if symbol == Cross {
		symbolString = "*"
	} else if symbol == Circle {
		symbolString = "o"
	}
	fmt.Printf("Player %d Move with Symbol %s at Position X: %d Y: %d\n", player.getID(), symbolString, x, y)
	g.board.PrintBoard()
	fmt.Println("")
}

func (g *Game) PrintResult() {
	switch g.gameStatus {
	case GameInProgress:
		fmt.Println("Game in Between")
	case GameDraw:
		fmt.Println("Game Drawn")
	case FirstPlayerWin:
		fmt.Println("First Player Win")
	case SecondPlayerWin:
		fmt.Println("Second Player Win")
	default:
		fmt.Println("Invalid Game Status")
	}
	g.board.PrintBoard()
}
