package server

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	// First player to connect is Host
	Host *Player `json:"host"`

	// Second player to connect is Opponent
	Opponent *Player `json:"opponent"`

	// All pokemon
	AvailablePokemon []Pokemon `json:"pokemon"`

	// Action pokemon fighting in the game
	FightingPokemon []*Pokemon `json:"fightingPkmn"`

	// The player currently taking action in the game
	TurnTaker uuid.UUID `json:"turnTaker"`
}

// all pokemon in one player's party have fainted
// TODO: allow for multiple pokemon per player
func (g *Game) IsGameOver() bool {
	for i := range g.Host.Pokemon {
		if g.Host.Pokemon[i].Hp <= 0 {
			return true
		}
	}

	for i := range g.Opponent.Pokemon {
		if g.Opponent.Pokemon[i].Hp <= 0 {
			return true
		}
	}

	return false
}

func (g *Game) AddPlayerToMatch(p *Player) {
	if g.Host != nil && g.Opponent != nil {
		panic("Game is full")
	}

	if g.Host == nil {
		g.Host = p
	} else {
		g.Opponent = p
	}
}

// TODO: Make this based off pokemon stats
func (g *Game) AlternateTurns() {
	if g.Host.ID == g.TurnTaker {
		g.TurnTaker = g.Opponent.ID
	} else {
		g.TurnTaker = g.Host.ID
	}
}

func (g *Game) logGameStatus() {
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
