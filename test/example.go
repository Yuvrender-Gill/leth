package test

import (
	"fmt"
	"log"
	//"math/big"

	"github.com/ChainSafeSystems/leth/bindings"
	"github.com/ChainSafeSystems/leth/core"

	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestExample() {
	conn, err := core.NewConnection("ganache")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}	

	// //auth, err := bind.NewTransactor(strings.NewReader(password), "my awesome super secret password")

	// // Create an IPC based RPC connection to a remote node
	// // conn, err := ethclient.Dial("http://localhost:8545")
	// // if err != nil {
	// // 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// // }

	//address := common.HexToAddress("0xecf168f325e745f196df6a21e4779ddf338e373a")
	// Instantiate the contract and display its name
	address, err := core.ContractAddress("Example", "ganache")
	if err != nil {
		log.Fatal(err)
	}

	ex, err := bindings.NewExample(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	owner, err := ex.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())

	// ex.ExampleFilterer.FromBlock = big.NewInt(0)
	// ex.ExampleFilterer.FilterLogs(context.Background, query)
}