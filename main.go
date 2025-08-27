package main

import (
	"fmt"
	"log"
)

func main() {
	// Connect to the server
	client, err := NewRPCClient("ws://localhost:25585")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	player1 := Player{Name: "Test"}
	player2 := Player{Name: "Test2"}

	players := []any{
		player1,
		player2,
	}

	fmt.Println(addToAllowlist(*client, players))

	//fmt.Println(getRpcDiscover(*client))
}
