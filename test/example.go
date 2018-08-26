package test

import (
	"fmt"
	"log"

	"github.com/noot/leth/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func TestExample() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	ex, err := bind.NewExample(common.HexToAddress("0x70ea7bcc6bba08ae16cc51f0520b8746740560ce"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	owner, err := ex.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())
}