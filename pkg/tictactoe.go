package pkg

import "errors"

const (
	playerX = "X"
	playerO = "O"
)

var ErrNotPlayerTurn = errors.New("not player turn")

type Field struct {
	X, Y int
}

type Game struct {
}

type Players [2]string

type Board [3][3]bool

func NewGame() Game {
	return Game{}
}

func (g *Game) Init() {

}

func (g *Game) Take(player string, field Field) error {
	return nil
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
