package server

import "github.com/google/uuid"

type MonsterAdder struct {
	PlayerName  string `json:"player_name"`
	MonsterName string `json:"monster_name"`
}

func NewPlayer(name string) Player {
	return Player{
		Name:    name,
		Pokemon: []*Pokemon{},
		ID:      uuid.New(),
	}
}

// this will host the player struct and its data
type Player struct {
	Name    string     `json:"player-name"`
	Pokemon []*Pokemon `json:"player-pokemon"`
	ID      uuid.UUID
}

func (p *Player) AddPokemon(pkmn *Pokemon) {
	p.Pokemon = append(p.Pokemon, pkmn)
}
