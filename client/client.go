package client

import (
	"fmt"
	"gkmn/server"
)

func ClientStart() {

	var action string = MainMenu()

	if action == "host" {
		// later will build a docker container
		// server.Server()
	} else if action == "connect" {
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

		// Player chooses pokemon to fight with
		monster := ChooseMonster()
		resp, err = AddPokemonToPlayer(player.Name, monster)
		if err != nil {
			fmt.Println("Player ", playerName, " Failed to add pokemon : ", err)
			return
		}
		fmt.Println(resp.StatusCode, " Player ", playerName, " added monster: ", monster)

		game, err := GameData()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
		// a "while" loop that goes until the game is over happens here.
		isOver, err := IsGameOver()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
		}
		for isOver != true {
			// generate and get actions
			choice := AttackMenu()

			// temporary print
			fmt.Println("json data")
			fmt.Println(game.Pokemon[0].Hp) // returns 100
			if choice == "tackle" {
				// call attack, it returns a game state -> which is the struct of interest
				game, err := BasicAttack()
				if err != nil {
					fmt.Println("failed attack called: ", err)
					return
				}
				fmt.Println("json data after attack")
				fmt.Println(game.Pokemon[1].Hp)
			} else if choice == "quit" {
				return
			}

			isOver, err = IsGameOver()
			if err != nil {
				fmt.Println("Connection Failed: ", err)
			}
		}
	}
}
