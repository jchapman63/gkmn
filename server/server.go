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
	game := Game{
		AvailablePokemon: []Pokemon{
			Pika,
			Bulbasaur,
			Gibble,
			Whooper,
		},
	}

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var player Player

		err := decoder.Decode(&player)
		if err != nil {
			panic(err)
		}

		game.AddPlayerToMatch(&player)
	})

	// a simple attack as a demo
	// TODO: FINISH handling decoded JSON for one player's specific mon to attack another player's specific mon with a specific move
	http.HandleFunc("/damage", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var results any
		err := decoder.Decode(results)
		if err != nil {
			panic(err)
		}
		fmt.Println(results)
		// pika.Attack(bulbasaur, tackle)
	})

	http.HandleFunc("/addPokemonToPlayer", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var adder MonsterAdder
		err := decoder.Decode(&adder)
		if err != nil {
			panic(err)
		}

		playerName := adder.PlayerName
		monsterName := adder.MonsterName
		monster := NewMonster(monsterName)
		for i := range game.Players {
			if game.Players[i].Name == playerName {
				game.Players[i].Pokemon = append(game.Players[i].Pokemon, &monster)
			}
		}
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
	})

	// return digestable game state
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	})

	// calls Game's method to check if game over
	http.HandleFunc("/isOver", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game.IsGameOver())
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
