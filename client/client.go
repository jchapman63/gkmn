package client

// TODO: Clean up docker container after game or if game ends early
import (
	"fmt"
	"gkmn/server"
	"os/exec"
	"time"
)

// initialize game variables
var player server.Player
var game server.Game
var opponent server.Player

func ClientStart() {

	var action string = MainMenu()
	var isHost bool

	// attempt to build host
	if action == "host" {
		startServer()
		isHost = true
	}

	err := UpdateGameData(&game)
	if err != nil {
		panic(err)
	}

	joinAndWait()
	playerChoosePokemon()
	findOpponent()

	// set initial turn
	if isHost {
		_, err := ChangeTurns()
		if err != nil {
			panic(err)
		}
	}

	// ensure first turn is selected
	for game.TurnTaker.String() == "" {
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
		updatePlayer()
		if game.TurnTaker.String() == player.ID.String() {
			// TODO: see how this is incomplete
			fmt.Println(player)
			// fmt.Println("--------------------------")
			// fmt.Printf("%s\n%s: %d", player.Name, player.Pokemon[0].Name, player.Pokemon[0].Hp)
			// fmt.Println("--------------------------")
			choice := AttackMenu()
			if choice != "quit" {
				// temporary
				pkmnToAttack := opponent.Pokemon[0].ID
				fmt.Println("pkmnToAttack", pkmnToAttack)
				_, err := AttackPkmn(pkmnToAttack, choice)
				if err != nil {
					panic(err)
				}

				_, err = ChangeTurns()
				if err != nil {
					panic(err)
				}
				UpdateGameData(&game)
			} else if choice == "quit" {
				return
			}
		} else {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("Waiting for turn!")
		}

		isOver, err = IsGameOver()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
	}
}

func playerChoosePokemon() {
	// Player chooses pokemon to fight with
	monster := ChooseMonster()
	_, err := AddPokemonToPlayer(player.ID, monster)
	if err != nil {
		fmt.Println("Player ", player.Name, " Failed to add pokemon : ", err)
		return
	}
	updatePlayer()
}

func joinAndWait() {
	// player create interface
	playerName := NamePlayer()

	// Player joining sever
	JoinGame(playerName, &player)
	fmt.Println(player.Name, "Joined the server")

	// player queue
	if len(game.Players) != 2 {
		fmt.Println("Waiting for player 2")
	}
	for len(game.Players) != 2 {
		UpdateGameData(&game)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Checking for other players...")
	}
}

func findOpponent() {
	// find player's opponent
	opponent = *game.Players[1]
	if player.ID == opponent.ID {
		opponent = *game.Players[0]
	}

	// wait for opponent to select pokemon
	for len(opponent.Pokemon) == 0 {
		UpdateGameData(&game)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Waiting for opponent selection...")
	}
}

func updatePlayer() {
	for _, connectedPlayer := range game.Players {
		if connectedPlayer.ID == player.ID {
			player = *connectedPlayer
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
