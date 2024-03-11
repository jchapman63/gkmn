package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gkmn/server"
	"io"
	"net/http"

	"github.com/google/uuid"
)

var baseUrl = "http://localhost:8080"

func UpdateGameData(g *server.Game) error {
	_, err := jsonResponseToGameStruct(g, baseUrl+"/state")
	if err != nil {
		panic(err)
	}
	return nil
}

// g, the struct to unpack into
// endpoint, the full api url
// should fully update g
func jsonResponseToGameStruct(g *server.Game, endpoint string) (*server.Game, error) {
	respJSON, err := http.Get(endpoint)
	if err != nil {
		fmt.Print("Data Request Failed: ", err)
		return nil, err
	}
	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		fmt.Println("Error reading json: ", err)
		return nil, err
	}
	if err := json.Unmarshal(bodyJSON, &g); err != nil {
		return nil, err
	}
	return g, nil
}

func AvailableMonsters() ([]string, error) {
	var monsters []string
	respJSON, err := http.Get(baseUrl + "/getMonsters")
	if err != nil {
		fmt.Print("Data Request Failed: ", err)
		return nil, err
	}

	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		fmt.Println("Error reading json: ", err)
		return nil, err
	}

	if err := json.Unmarshal(bodyJSON, &monsters); err != nil {
		return nil, err
	}

	return monsters, nil
}

func IsGameOver() (bool, error) {
	respJSON, err := http.Get(baseUrl + "/isOver")
	if err != nil {
		return false, err
	}

	var isOver bool
	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(bodyJSON, &isOver); err != nil {
		return false, err
	}

	return isOver, nil
}

func ChangeTurns() (uuid.UUID, error) {
	respJSON, err := http.Get(baseUrl + "/changeTurns")
	if err != nil {
		panic(err)
	}

	var turnID uuid.UUID
	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		panic(nil)
	}
	if err := json.Unmarshal(bodyJSON, &turnID); err != nil {
		panic(nil)
	}

	return turnID, nil
}

func JoinGame(p string) (server.Player, error) {
	newData, _ := json.Marshal(p)
	resp, err := http.Post(baseUrl+"/join", "application/json", bytes.NewBuffer(newData))
	if err != nil {
		panic(err)
	}

	var player server.Player
	bodyJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bodyJSON, &player); err != nil {
		panic(err)
	}

	return player, nil
}

func AttackPkmn(target uuid.UUID, move string) (*http.Response, error) {
	attackInfo := map[string]any{
		"target": target,
		"move":   move,
	}
	data, _ := json.Marshal(attackInfo)
	resp, err := http.Post(baseUrl+"/damage", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	return resp, nil
}

func AddPokemonToPlayer(playerName string, pkmnName string) (*http.Response, error) {

	data := server.MonsterAdder{
		PlayerName:  playerName,
		MonsterName: pkmnName,
	}

	newData, _ := json.Marshal(data)
	resp, err := http.Post(baseUrl+"/addPokemonToPlayer", "application/json", bytes.NewBuffer(newData))
	if err != nil {
		panic(err)
	}
	return resp, nil
}
