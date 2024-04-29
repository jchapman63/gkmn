package main

import (
	"flag"
	"fmt"
	"slices"

	"gkmn/client"
	"gkmn/server"
)

// / Dev Doc: starting out, I will just ask for simple input to get the server started.  Later, a more refined CLI would be nice.
func main() {
	fmt.Printf("Welcome to Gokemon!\n\n")

	flag.Parse()
	if slices.Contains(flag.Args(), "s") {
		fmt.Println("Starting Gokemon Server!")
		server.Server()
	} else {
		fmt.Println("Starting Gokemon Client!")
		client.ClientStart()
	}
}
