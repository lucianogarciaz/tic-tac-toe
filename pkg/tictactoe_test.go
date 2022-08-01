package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	t.Run(`Given a new game,
	when it is initialized,
	then the state changes to open`, func(t *testing.T) {
		game := pkg.NewGame()
		game.Init()

		assert.Equal(t, "open", game.State())
	})
}
