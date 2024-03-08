package server

import "github.com/google/uuid"

// the pokemon struct will need a move set. Moves will be a struct of their own.  I will start off just worrying about
// damage based moves to get me started.

// General pokemon type
// lowercase fields are not exposed, so these are uppercase
type Pokemon struct {
	Name  string       `json:"pokemon-name"`
	Hp    int          `json:"hp"`
	Moves []DamageMove `json:"moves"`
	ID    uuid.UUID    `json:"pkmn_id"`
}

func (p *Pokemon) LoseHealth(amount int) {
	p.Hp -= amount
}

// create a new pokemon for player based on request
// defaults to returning a pikachu
func NewMonster(name string) Pokemon {
	monsters := []Pokemon{Pika, Gibble, Bulbasaur, Whooper}

	for i := range monsters {
		if monsters[i].Name == name {
			return Pokemon{
				Name:  monsters[i].Name,
				Hp:    monsters[i].Hp,
				Moves: monsters[i].Moves,
				ID:    uuid.New(),
			}
		}
	}
	return Pokemon{
		Name:  Pika.Name,
		Hp:    Pika.Hp,
		Moves: Pika.Moves,
		ID:    uuid.New(),
	}
}

type DamageMove struct {
	Name  string
	Power int
}

// available pokemon
var Bulbasaur = Pokemon{
	Name: "bulbasaur",
	Hp:   100,
	Moves: []DamageMove{
		Tackle,
	},
}

var Pika = Pokemon{
	Name: "pikachu",
	Hp:   100,
	Moves: []DamageMove{
		Tackle,
	},
}

var Gibble = Pokemon{
	Name: "gibble",
	Hp:   100,
	Moves: []DamageMove{
		Tackle,
	},
}

var Whooper = Pokemon{
	Name: "whooper",
	Hp:   100,
	Moves: []DamageMove{
		Tackle,
	},
}

// available moves
var Tackle = DamageMove{
	Name:  "tackle",
	Power: 10,
}
