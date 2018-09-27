package test

import (
	"fmt"
	"log"
	//"io/ioutil"
	//"strings"
	//"math/big"

	"github.com/ChainSafeSystems/leth/bindings"
	"github.com/ChainSafeSystems/leth/core"

	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func Test() {
	conn, err := core.NewConnection("default")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}	

	// alternatively, create an IPC based RPC connection to a remote node
	// conn, err := ethclient.Dial("http://localhost:8545")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }

	// instantiate the contract and display its name
	address, err := core.ContractAddress("Token", "default")
	if err != nil {
		log.Fatal(err)
	}

	// alternatively, manually specify the contract address
	// address := common.HexToAddress("0xecf168f325e745f196df6a21e4779ddf338e373a")

	ex, err := bindings.NewExample(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	owner, err := ex.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())

	// create a new transactor to write to the contract 
	// file, err := ioutil.ReadFile("./keystore/UTC--2018-07-31T19-43-24.789345360Z--e8b7b81f281a947840de4b23f40442b3843c5f49")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// auth, err := bind.NewTransactor(strings.NewReader(string(file)), "password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(auth)

	//exRaw := new(bindings.ExampleRaw)
	// create a tx to call the fallback function
	//tx, err := ex.Transfer(auth, big.NewInt(1))

	// ex.ExampleFilterer.FromBlock = big.NewInt(0)
	// ex.ExampleFilterer.FilterLogs(context.Background, query)
}