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

	// Make a simple request
	response, err := client.Call("rpc.discover", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", response.Result)
}
