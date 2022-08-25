package pkg

import "errors"

const (
	playerX      Player = "X"
	playerO      Player = "O"
	boardRows           = 3
	boardColumns        = 3
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

func (g *Game) Take(player Player, field Field) error {
	if field.X >= boardRows || field.Y >= boardColumns {
		return ErrOutOfBounds
	}

	if player != g.Turn {
		return ErrNotPlayerTurn
	}

	if g.board[field.X][field.Y] {
		return ErrFieldAlreadyTaken
	}

	g.Next()
	g.board[field.X][field.Y] = true

	return nil
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

func (g Game) PlayerOne() Player {
	return playerX
}

func (g Game) PlayerTwo() Player {
	return playerO
}
