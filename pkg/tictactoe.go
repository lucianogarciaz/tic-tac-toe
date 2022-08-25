package pkg

import "errors"

const (
	playerX      = "X"
	playerO      = "O"
	boardRows    = 3
	boardColumns = 3
)

var (
	ErrNotPlayerTurn     = errors.New("not player turn")
	ErrFieldAlreadyTaken = errors.New("field already taken")
	ErrOutOfBounds       = errors.New("field out of bounds")
)

type Field struct {
	X, Y int
}

type Game struct {
	Turn  Player
	board Board
}

type Turn string

type Players [2]Player
type Player string

type Board [boardRows][boardColumns]bool

func NewGame() Game {
	return Game{
		Turn:  playerX,
		board: Board{},
	}
}

func (g *Game) Init() {

}

func (g *Game) Take(player string, field Field) error {
	if field.X >= boardRows || field.Y >= boardColumns {
		return ErrOutOfBounds
	}

	if player == string(g.Turn) && g.board[field.X][field.Y] {
		return ErrFieldAlreadyTaken
	}

	if player == string(g.Turn) {
		g.Next()
		g.board[field.X][field.Y] = true
		return nil
	}

	return ErrNotPlayerTurn
}

func (g *Game) Next() {
	if g.Turn == playerX {
		g.Turn = playerO

		return
	}

	g.Turn = playerX
}

func (g Game) State() string {
	return "open"
}

func (g Game) Board() Board {
	return Board{}
}

func (g Game) CountPlayers() int {
	return len(Players{})
}

func (g Game) PlayerOne() string {
	return playerX
}

func (g Game) PlayerTwo() string {
	return playerO
}
