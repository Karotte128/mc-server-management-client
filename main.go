package main

import (
	"fmt"
	"log"

	"github.com/karotte128/mcsmplib"
)

func main() {
	// Connect to the server
	client, err := NewRPCClient("ws://localhost:25585")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	player1 := mcsmplib.Player{ID: "d8d5a923-7b20-43d8-883b-1150148d6955"}
	player2 := mcsmplib.Player{Name: "Test2"}

	players := []mcsmplib.Player{
		player1,
		player2,
	}

	request := mcsmplib.AllowlistAdd(players)

	response, err := client.Call(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", response.Result)
}
