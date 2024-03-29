package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var gameServer *http.Server = &http.Server{
	Addr: ":8080",
}

func Server() {
	game := &Game{
		AvailablePokemon: []Pokemon{
			Pika,
			Bulbasaur,
			Gibble,
			Whooper,
		},
	}

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var playerName string

		err := decoder.Decode(&playerName)
		if err != nil {
			panic(err)
		}

		player := NewPlayer(playerName)

		game.AddPlayerToMatch(&player)

		// make player grabbable from server
		json.NewEncoder(w).Encode(player)

		game.logGameStatus()
	})

	// a simple attack as a demo
	// TODO: FINISH handling decoded JSON for one player's specific mon to attack another player's specific mon with a specific move
	http.HandleFunc("/damage", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var results map[string]any
		err := decoder.Decode(&results)
		if err != nil {
			panic(err)
		}

		// unpack json
		moveName := results["move"]

		var move DamageMove
		// TODO: fix to handle any move
		if moveName == "tackle" {
			move = Tackle
		}
		target := results["target"]

		for i := range game.FightingPokemon {
			if game.FightingPokemon[i].ID.String() == target {
				fmt.Println("target: ", target)
				game.FightingPokemon[i].LoseHealth(move.Power)
			}
		}

		game.logGameStatus()
	})

	// TODO: Update to be done by player ID and not name
	http.HandleFunc("/addPokemonToPlayer", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var adder MonsterAdder
		err := decoder.Decode(&adder)
		if err != nil {
			panic(err)
		}

		playerID := adder.PlayerID
		monsterName := adder.MonsterName
		monster := NewMonster(monsterName)

		if playerID == game.Host.ID {
			game.Host.Pokemon = append(game.Host.Pokemon, &monster)
		} else {
			game.Opponent.Pokemon = append(game.Opponent.Pokemon, &monster)
		}

		game.FightingPokemon = append(game.FightingPokemon, &monster)

		game.logGameStatus()
	})

	// allow players to choose an available monster
	http.HandleFunc("/getMonsters", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		pokemon := []string{
			Bulbasaur.Name,
			Gibble.Name,
			Pika.Name,
			Whooper.Name,
		}
		json.NewEncoder(w).Encode(pokemon)

		game.logGameStatus()
	})

	// update and return new player id who's turn it is
	http.HandleFunc("/changeTurns", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		game.AlternateTurns()
		game.logGameStatus()
		json.NewEncoder(w).Encode(game.TurnTaker)
	})

	// return digestable game state, might not ever use
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
		game.logGameStatus()
	})

	// calls Game's method to check if game over
	http.HandleFunc("/isOver", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game.IsGameOver())
	})

	http.HandleFunc("/storeGame", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Trying to Store")
		StoreGameState(game)
	})

	http.HandleFunc("/testDatabase", func(w http.ResponseWriter, r *http.Request) {
		results := TestConnection()
		fmt.Println(results)
	})

	fmt.Println("Server is listening localhost:8080")

	// serve the application
	log.Fatal(gameServer.ListenAndServe())
	if game.IsGameOver() {
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()
		gameServer.Shutdown(shutdownCtx)
	}
}
