package pkg

import "errors"

const (
	playerX = "X"
	playerO = "O"
)

var (
	ErrNotPlayerTurn     = errors.New("not player turn")
	ErrFieldAlreadyTaken = errors.New("field already taken")
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

type Board [3][3]bool

func NewGame() Game {
	return Game{
		Turn:  playerX,
		board: Board{},
	}
}

func (g *Game) Init() {

}

func (g *Game) Take(player string, field Field) error {
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
