package test

import (
	"fmt"
	"log"

	"github.com/noot/leth/example"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestExample() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	token, err := NewExample(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}
	// name, err := token.Name(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token name: %v", err)
	// }
	// fmt.Println("Token name:", name)
}