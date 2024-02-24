package server

type MonsterAdder struct {
	PlayerName  string `json:"player_name"`
	MonsterName string `json:"monster_name"`
}

// this will host the player struct and its data
type Player struct {
	Name string `json:"player-name"`
	// pointer for addressing ?
	Pokemon []*Pokemon `json:"player-pokemon"`

	Pokedex []*Pokemon `json:"player-pokedex"`
}

func (p *Player) AddPokemon(pkmn *Pokemon) {
	p.Pokemon = append(p.Pokemon, pkmn)
}
