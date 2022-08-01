package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lucianogarciaz/tic-tac-toe/pkg"
)

func TestGame(t *testing.T) {
	game := pkg.NewGame()
	game.Init()

	t.Run(`Given a new game,
	when it is initialized,
	then the state changes to open`, func(t *testing.T) {
		assert.Equal(t, "open", game.State())
	})

	t.Run(`Given a new game,
	when it is initialized,
	then all the cells in the board are empty`, func(t *testing.T) {
		assert.Empty(t, game.Board())
	})
}
