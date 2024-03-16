package server

import (
	"testing"

	"github.com/google/uuid"
)

func TestTakingTurns(t *testing.T) {
	bulb1 := NewMonster(Bulbasaur.Name)
	bulb2 := NewMonster(Bulbasaur.Name)
	host := Player{Name: "Jordan", Pokemon: []*Pokemon{&bulb1}, ID: uuid.New()}
	opponent := Player{Name: "FooBar", Pokemon: []*Pokemon{&bulb1}, ID: uuid.New()}
	var game = Game{
		FightingPokemon: []*Pokemon{
			&bulb1, &bulb2,
		},
		Host:     &host,
		Opponent: &opponent,
	}

	// test that turn gets assigned to first player in game
	game.AlternateTurns()
	if game.TurnTaker != game.Host.ID {
		t.Errorf("Did not assign TurnTaker properly.  Expected: %s, Got: %s", game.Host.ID, game.TurnTaker)
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
