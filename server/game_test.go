package server

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestTakingTurns(t *testing.T) {
	bulb1 := NewMonster(Bulbasaur.Name)
	bulb2 := NewMonster(Bulbasaur.Name)
	var game = Game{
		Players: []*Player{
			{
				Name:    "Jordan",
				Pokemon: []*Pokemon{&bulb1},
				ID:      uuid.New(),
			},
			{
				Name:    "Jordan2",
				Pokemon: []*Pokemon{&bulb2},
				ID:      uuid.New(),
			},
		},
		FightingPokemon: []*Pokemon{
			&bulb1, &bulb2,
		},
	}
	fmt.Printf("%s: %s\n", game.Players[0].Name, game.Players[0].ID.String())
	fmt.Printf("%s: %s\n", game.Players[1].Name, game.Players[1].ID.String())

	// test that turn gets assigned to first player in game
	game.AlternateTurns()
	if game.TurnTaker != game.Players[0].ID {
		t.Errorf("Did not assign TurnTaker properly.  Expected: %s, Got: %s", game.Players[0].ID, game.TurnTaker)
	}

	// test that turn changes to a different player than it currently is
	currentTurn := game.TurnTaker
	game.AlternateTurns()
	if game.TurnTaker == currentTurn {
		t.Errorf("Did not update TurnTaker properly. Previous: %s, Current: %s", currentTurn, game.TurnTaker)
	}

	currentTurn = game.TurnTaker
	game.AlternateTurns()
	if game.TurnTaker == currentTurn {
		t.Errorf("Did not update TurnTaker properly. Previous: %s, Current: %s", currentTurn, game.TurnTaker)
	}
}
