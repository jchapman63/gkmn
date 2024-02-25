package client

// TODO: Clean up docker container after game or if game ends early
import (
	"fmt"
	"gkmn/server"
	"os/exec"
	"time"
)

func ClientStart() {

	var action string = MainMenu()

	// attempt to build host
	if action == "host" {

		fmt.Println("Building Server...")
		// build and run docker container
		exec.Command("cd", "../").Output()
		_, err := exec.Command("docker", "build", "-t", "gokemon-build", ".").Output()
		if err != nil {
			panic(err)
		}

		_, err = exec.Command("docker", "run", "-d", "-p", "8080:8080", "gokemon-build:latest").Output()
		if err != nil {
			panic(err)
		}
	}
	// initialize game
	var game server.Game
	err := UpdateGameData(&game)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("here is game: ", game)
	}

	// player create interface
	playerName := CreatePlayer()
	var player = &server.Player{
		Name: playerName,
	}

	// Player joining sever
	resp, err := JoinGame(player)
	if err != nil {
		fmt.Println("Player ", playerName, "Connection Failed: ", err)
		return
	}
	fmt.Println(resp.StatusCode, " Player ", playerName, " Connected Server! ")

	// player queue
	if len(game.Players) != 2 {
		fmt.Println("Waiting for player 2")
	}
	for len(game.Players) != 2 {
		UpdateGameData(&game)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Checking for other players...")
	}

	// Player chooses pokemon to fight with
	monster := ChooseMonster()
	resp, err = AddPokemonToPlayer(player.Name, monster)
	if err != nil {
		fmt.Println("Player ", playerName, " Failed to add pokemon : ", err)
		return
	}
	fmt.Println(resp.StatusCode, " Player ", playerName, " added monster: ", monster)

	// play the game
	isOver, err := IsGameOver()
	if err != nil {
		fmt.Println("Connection Failed: ", err)
	}
	for !isOver {
		// generate and get actions
		choice := AttackMenu()

		// temporary print
		// fmt.Println("json data")
		// fmt.Println(game.Pokemon[0].Hp) // returns 100
		if choice != "quit" {
			// player's pokemon attacks oppenents pokemon
		} else if choice == "quit" {
			return
		}

		isOver, err = IsGameOver()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
	}
}
