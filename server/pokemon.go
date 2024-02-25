package server

// the pokemon struct will need a move set. Moves will be a struct of their own.  I will start off just worrying about
// damage based moves to get me started.

// General pokemon type
// lowercase fields are not exposed, so these are uppercase
type Pokemon struct {
	Name  string       `json:"pokemon-name"`
	Hp    int          `json:"hp"`
	Moves []DamageMove `json:"moves"`
}

// Removes health from the pokemon based on attack's power
func (p *Pokemon) Attack(o *Pokemon, attack DamageMove) {
	o.Hp -= attack.Power
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
