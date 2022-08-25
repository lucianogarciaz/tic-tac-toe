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

	t.Run(`Given a new game,
	when it is initialized,
	then we have two players in the game`, func(t *testing.T) {
		assert.Equal(t, 2, game.CountPlayers())
	})

	t.Run(`Given a new game,
	when it is initialized,
	then we have two players X and O`, func(t *testing.T) {
		assert.Equal(t, "X", game.PlayerOne())
		assert.Equal(t, "O", game.PlayerTwo())
	})

	t.Run(`Given a new game and a player O,
	When taking a field,
	Then it should return an ErrNotPlayerTurn error`, func(t *testing.T) {
		expectedError := pkg.ErrNotPlayerTurn
		field := pkg.Field{1, 2}
		err := game.Take(game.PlayerTwo(), field)
		assert.Equal(t, expectedError, err)
	})

	t.Run(`Given a new game and player X,
	When taking a field,
	Then it should return no error`, func(t *testing.T) {
		field := pkg.Field{1, 2}
		assert.NoError(t, game.Take(game.PlayerOne(), field))
	})

	t.Run(`Given a new game 
		When player X takes a field and player O takes a field
		Then it should return no error`, func(t *testing.T) {
		game := pkg.NewGame()
		game.Init()

		assert.NoError(t, game.Take(game.PlayerOne(), pkg.Field{0, 0}))
		assert.NoError(t, game.Take(game.PlayerTwo(), pkg.Field{0, 1}))
	})

	t.Run(`Given a game that was started
		When player 0 plays
		Then player X takes the next turn`, func(t *testing.T) {
		game := pkg.NewGame()
		game.Init()

		assert.NoError(t, game.Take(game.PlayerOne(), pkg.Field{0, 0}))
		assert.NoError(t, game.Take(game.PlayerTwo(), pkg.Field{0, 1}))
		assert.NoError(t, game.Take(game.PlayerOne(), pkg.Field{0, 2}))
	})

	t.Run(`Given a game that was started
		When player X takes a field
		and player O a field already taken
		Then an ErrFieldAlreadyTaken error should be returned`, func(t *testing.T) {
		game := pkg.NewGame()
		game.Init()

		assert.NoError(t, game.Take(game.PlayerOne(), pkg.Field{0, 0}))
		assert.ErrorIs(t, game.Take(game.PlayerTwo(), pkg.Field{0, 0}), pkg.ErrFieldAlreadyTaken)
	})

	t.Run(`Given a new game,
		When a player tries to take an out-of-bounds field,
		Then an ErrOutOfBounds error should be returned`, func(t *testing.T) {
		game := pkg.NewGame()
		game.Init()

		assert.ErrorIs(t, game.Take(game.PlayerOne(), pkg.Field{100, 100}), pkg.ErrOutOfBounds)
	})
}
