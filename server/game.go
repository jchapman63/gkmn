package server

// THINKPOINT: check that there is not a memory problem here
type Game struct {
	Players          []*Player `json:"players"`
	AvailablePokemon []Pokemon `json:"pokemon"`
}

// all pokemon in one player's party have fainted
// TODO seriously update logic here
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
