package client

// TODO: Clean up docker container after game or if game ends early
import (
	"fmt"
	"gkmn/server"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

// initialize game variables
var player server.Player
var game server.Game

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
		StoreGameData()
		panic(err)
	}

	joinAndWait()
	playerChoosePokemon()
	setStartState()

	// set initial turn
	if isHost {
		_, err := ChangeTurns()
		if err != nil {
			StoreGameData()
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

	refreshCount := 0
	for !isOver {
		UpdateGameData(&game)
		if isHost && (game.TurnTaker.String() == game.Host.ID.String()) {
			PrintGameState(&game)
			// host plays
			played := attackEnemy(&isHost)
			if !played {
				// end game
				// TODO: clean up container
				return
			}
			ChangeTurns()
			UpdateGameData(&game)
			refreshCount = 0
		} else if !isHost && (game.TurnTaker.String() == game.Opponent.ID.String()) {
			PrintGameState(&game)
			// Opponent plays
			played := attackEnemy(&isHost)
			if !played {
				// end game
				// TODO: clean up container
				return
			}
			ChangeTurns()
			UpdateGameData(&game)
			refreshCount = 0
		} else {
			if refreshCount == 0 {
				fmt.Println("Waiting for turn!")
			}
			time.Sleep(1000 * time.Millisecond)
			refreshCount += 1
		}

		isOver, err = IsGameOver()
		if err != nil {
			StoreGameData()
			fmt.Println("Connection Failed: ", err)
		}
	}
	StoreGameData()
}

func setStartState() {
	fmt.Println("Waiting for all players to select a monster!")
	for len(game.Host.Pokemon) == 0 || len(game.Opponent.Pokemon) == 0 {
		time.Sleep(1000 * time.Millisecond)
		UpdateGameData(&game)
	}
}

// true if attacked, false is quit
func attackEnemy(isHost *bool) bool {
	choice := AttackMenu()

	var pkmnToAttack uuid.UUID
	if *isHost {
		pkmnToAttack = game.Opponent.Pokemon[0].ID
	} else {
		pkmnToAttack = game.Host.Pokemon[0].ID
	}

	if choice != "quit" {
		_, err := AttackPkmn(pkmnToAttack, choice)
		if err != nil {
			// TODO: The server is shut down at this point, there is no use in logging the game state.
			StoreGameData()
			panic(err)
		}
	} else {
		StoreGameData()
		return false
	}

	return true
}

func playerChoosePokemon() {
	// Player chooses pokemon to fight with
	monster := ChooseMonster()
	_, err := AddPokemonToPlayer(player.ID, monster)
	if err != nil {
		fmt.Println("Player ", player.Name, " Failed to add pokemon : ", err)
		return
	}
	UpdateGameData(&game)
}

func joinAndWait() {
	// create player name and join game
	playerName := NamePlayer()
	JoinGame(playerName, &player)
	fmt.Println(player.Name, "Joined the server")
	UpdateGameData(&game)

	// do not exit until both Host and Opponent are in game
	fmt.Println("Checking for other players...")
	for game.Opponent == nil || game.Host == nil {
		UpdateGameData(&game)
		time.Sleep(1000 * time.Millisecond)
	}
}

func startServer() {
	fmt.Println("Building Server...")
	// build and run docker container
	exec.Command("cd", "../").Output()

	_, err := exec.Command("docker", "compose", "up", "-d", "--force-recreate").Output()
	if err != nil {
		fmt.Println("Failed compose")
		panic(err)
	}
}
