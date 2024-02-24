package server

// THINKPOINT: check that there is not a memory problem here
type Game struct {
	Players []*Player  `json:"players"`
	Pokemon []*Pokemon `json:"pokemon"`
}

// all pokemon in one player's party have fainted
func (g *Game) IsGameOver() bool {
	for i := range g.Pokemon {
		if g.Pokemon[i].Hp <= 0 {
			return true
		}
	}
	return false
}

func (g *Game) AddPlayerToMatch(p *Player) {
	g.Players = append(g.Players, p)
}
