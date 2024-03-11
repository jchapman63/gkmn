package server

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	// Players connected to the game
	Players []*Player `json:"players"`

	// All pokemon
	AvailablePokemon []Pokemon `json:"pokemon"`

	// Action pokemon fighting in the game
	FightingPokemon []*Pokemon `json:"fightingPkmn"`

	// The player currently taking action in the game
	TurnTaker *uuid.UUID `json:"turnTaker"`
}

// all pokemon in one player's party have fainted
func (g *Game) IsGameOver() bool {
	for i := range g.Players {
		for j := range g.Players[i].Pokemon {
			if g.Players[i].Pokemon[j].Hp <= 0 {
				return true
			}
		}
	}
	return false
}

func (g *Game) AddPlayerToMatch(p *Player) {
	g.Players = append(g.Players, p)
}

// For players in game, switch who's turn it is to attack.
// If no player is set as TurnTaker, set turn taker to the
// first player in the array.
//
// TODO: Make this based off pokemon stats
func (g *Game) AlternateTurns() {
	if len(g.Players) != 2 {
		panic("Improper use of function `AlternateTurns`: Not enough players are in the game")
	}

	if g.TurnTaker == nil {
		fmt.Println("No turn set yet.  TurnTaker setting now")
		g.TurnTaker = &g.Players[0].ID
	} else {
		for i := range g.Players {
			current := g.TurnTaker
			if current != &g.Players[i].ID {
				g.TurnTaker = &g.Players[i].ID
			}
		}
	}

}

func (g *Game) logGameStatus() {
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
