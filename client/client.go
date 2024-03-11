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
	var isHost bool

	// attempt to build host
	if action == "host" {
		startServer()
		isHost = true
	}
	// initialize game
	var game server.Game
	err := UpdateGameData(&game)
	if err != nil {
		panic(err)
	}

	// player create interface
	playerName := NamePlayer()

	// Player joining sever
	player, err := JoinGame(playerName)
	if err != nil {
		fmt.Println("Player ", playerName, "Connection Failed: ", err)
		return
	}
	fmt.Println(player.Name, "Joined the server")

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
	_, err = AddPokemonToPlayer(player.Name, monster)
	if err != nil {
		fmt.Println("Player ", playerName, " Failed to add pokemon : ", err)
		return
	}

	// find player's opponent
	opponent := game.Players[1]
	if player.ID == opponent.ID {
		opponent = game.Players[0]
	}

	// wait for opponent to select pokemon
	for len(opponent.Pokemon) == 0 {
		UpdateGameData(&game)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Waiting for opponent selection...")
	}

	// set initial turn
	if isHost {
		_, err := ChangeTurns()
		if err != nil {
			panic(err)
		}
	}

	// ensure first turn is selected
	for game.TurnTaker == nil {
		UpdateGameData(&game)
	}

	// play the game
	isOver, err := IsGameOver()
	if err != nil {
		fmt.Println("Connection Failed: ", err)
	}
	for !isOver {
		// generate and get actions
		UpdateGameData(&game)
		fmt.Printf("TurnTaker: %s == player.ID: %s", game.TurnTaker, player.ID)
		fmt.Printf("Result of comparison: %t", *game.TurnTaker == player.ID)
		if *game.TurnTaker == player.ID {
			choice := AttackMenu()
			if choice != "quit" {
				// temporary
				pkmnToAttack := opponent.Pokemon[0].ID
				fmt.Println("pkmnToAttack", pkmnToAttack)
				_, err := AttackPkmn(pkmnToAttack, choice)
				if err != nil {
					panic(err)
				}

				id, err := ChangeTurns()
				if err != nil {
					panic(err)
				}
				fmt.Printf("result of turn taker from server: %s", id)
				UpdateGameData(&game)
			} else if choice == "quit" {
				return
			}
		} else {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Waiting for turn!")
		}

		isOver, err = IsGameOver()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
	}
}

func startServer() {
	fmt.Println("Building Server...")
	// build and run docker container
	exec.Command("cd", "../").Output()
	_, err := exec.Command("docker", "build", "-t", "gokemon-build", ".").Output()
	if err != nil {
		fmt.Println("image build failed")
		panic(err)
	}

	_, err = exec.Command("docker", "run", "-d", "-p", "8080:8080", "gokemon-build:latest").Output()
	if err != nil {
		fmt.Println("image run failed")
		panic(err)
	}
}
