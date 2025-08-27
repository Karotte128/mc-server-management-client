package main

import (
	"fmt"
	"log"
)

func getRpcDiscover(client RPCClient) string {
	response, err := client.Call("rpc.discover", nil)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", response.Result)
}

func addToAllowlist(client RPCClient, players []any) string {

	response, err := client.Call("minecraft:allowlist/add", players)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", response.Result)
}
