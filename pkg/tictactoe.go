package pkg

type Game struct {
}

func NewGame() Game {
	return Game{}
}

func (g *Game) Init() {

}

func (g Game) State() string {
	return "open"
}
