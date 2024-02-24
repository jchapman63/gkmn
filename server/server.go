package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Server() {
	gameServer := &http.Server{
		Addr: ":8080",
	}
	// initialize a simple game environment
	tackle := DamageMove{
		Name:  "tackle",
		Power: 10,
	}
	pika := &Pokemon{
		Name: "pikachu",
		Hp:   100,
		Moves: []DamageMove{
			tackle,
		},
	}
	bulbasaur := &Pokemon{
		Name: "bulbasaur",
		Hp:   100,
		Moves: []DamageMove{
			tackle,
		},
	}

	game := Game{
		Pokemon: []*Pokemon{
			pika,
			bulbasaur,
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

	http.HandleFunc("/addPokemonToPlayer", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var adder MonsterAdder
		err := decoder.Decode(&adder)
		if err != nil {
			panic(err)
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

	// a simple attack as a demo
	// TODO: Handle arguments for which pokemon attacks and which pokemon gets attacked
	http.HandleFunc("/damage", func(w http.ResponseWriter, r *http.Request) {
		pika.Attack(bulbasaur, tackle)
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
	// why nil here? But, this will serve the app
	log.Fatal(gameServer.ListenAndServe())

	if game.IsGameOver() {
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()
		gameServer.Shutdown(shutdownCtx)
	}
}

func localAttackTest(pika *Pokemon, bulb *Pokemon) {
	move := pika.Moves[0]
	pika.Attack(bulb, move)
	fmt.Print("bubla health: ", bulb.Hp)
}
