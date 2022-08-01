package pkg

type Game struct {
}

type Board [3][3]bool

func NewGame() Game {
	return Game{}
}

func (g *Game) Init() {

}

func (g Game) State() string {
	return "open"
}

func (g Game) Board() Board {
	return Board{}
}

func (g Game) CountPlayers() int {
	return 0
}
