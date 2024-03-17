package client

// NOTE: Ultimately, this file will need to be cleaned up, the functionality of server interactions should be from client.go and
// creating the CLI should only be done here.
import (
	"bufio"
	"fmt"
	"gkmn/server"
	"os"
	"strings"

	"github.com/nexidian/gocliselect"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// file design: function out all cli options to be called from a main loop in client.go
// Menus so far
// - MainMenu
// - AttackMenu
// - CreatePlayer
// - ChooseMonster
// - PrintGameState
func MainMenu() string {
	menu := gocliselect.NewMenu("How would you like to play?")

	menu.AddItem("Host", "host")
	menu.AddItem("Connect", "connect")

	choice := menu.Display()
	return choice
}

func NamePlayer() string {
	fmt.Print("Name your player: ")
	reader := bufio.NewReader(os.Stdin)
	// read up to and including delimiter
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

func ChooseMonster() string {
	pokemon, err := AvailableMonsters()
	if err != nil {
		fmt.Println("Failed to retrieve monsters from server")
	}

	actionMenu := gocliselect.NewMenu("Choose A Pokemon!")
	for i := range pokemon {
		titleMaker := cases.Title(language.English)
		key := titleMaker.String(pokemon[i])
		actionMenu.AddItem(key, pokemon[i])
	}

	actionChoice := actionMenu.Display()

	return actionChoice
}

// TODO: Add params of a pokemon's moves
func AttackMenu() string {
	actionMenu := gocliselect.NewMenu("Attack!")
	actionMenu.AddItem("Tackle", "tackle")
	actionMenu.AddItem("Quit", "quit")

	actionChoice := actionMenu.Display()

	return actionChoice
}

func PrintGameState(game *server.Game) {
	fmt.Println("--------------------------")
	fmt.Printf("Host: %s,\nHostMonster: %s,\nHP: %d;\n", game.Host.Name, game.Host.Pokemon[0].Name, game.Host.Pokemon[0].Hp)
	fmt.Println()
	fmt.Printf("Opponent: %s,\nOpponentMonster: %s,\nHP: %d;\n", game.Opponent.Name, game.Opponent.Pokemon[0].Name, game.Opponent.Pokemon[0].Hp)
	fmt.Println("--------------------------")
}
