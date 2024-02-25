package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gkmn/server"
	"io"
	"net/http"
)

var baseUrl = "http://localhost:8080"

// TODO: Refactor to fit documentation comments
// modify a pointer to a game object
// Meant to be called once at server start
// Server will manipulate a game struct of its own and
// send it here
func UpdateGameData(g *server.Game) error {
	_, err := jsonResponseToGameStruct(g, baseUrl+"/state")
	if err != nil {
		return err
	}
	return nil
}

// g, the struct to unpack into
// endpoint, the full api url
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

func JoinGame(p *server.Player) (*http.Response, error) {
	newData, _ := json.Marshal(p)
	resp, err := http.Post(baseUrl+"/join", "application/json", bytes.NewBuffer(newData))
	if err != nil {
		panic(err)
	}
	return resp, nil
}

// TODO switch attacker to be using a UUID and target to be using a UUID
func AttackPkmn(attacker string, target string, move string) (*http.Response, error) {
	attackInfo := map[string]string{
		"attacker": attacker,
		"target":   target,
		"move":     move,
	}
	data, _ := json.Marshal(attackInfo)

	// TODO: finish API call
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

// will later have parameters for the pokemon attacking (the client's mon)
// I could either pass game as parameter, or I could get a new game object... new is more up to date with the server...
// there is more abstraction I can do between these functions...
// if I am always going to be getting a new game struct, then I can have one function that takes the url as a parameter ?
// I may actually want to pass other params too though, so I will abstract atleast the RESP and JSON to struct ?
func BasicAttack() (*server.Game, error) {
	// question, will passing game here update the current game object????
	// var game *server.Game
	// http.Get(baseUrl + "/damage")
	// game, err := GameData()
	// if err != nil {
	// 	return nil, err
	// }

	// return game, nil
	return nil, nil
}
