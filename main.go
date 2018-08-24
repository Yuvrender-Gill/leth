package main

import (
	"fmt"
	"log"

	"github.com/noot/leth/core"
	"github.com/noot/leth/new"
)

func main() {
	//client := leth.Dial("http://localhost:8545")

	// var input string
	// fmt.Println("enter contract to compile (.sol file needs to be in contracts/)")
	// fmt.Scanln(&input)
	// path, _ := filepath.Abs("./contracts/" + input + ".sol")
	// fmt.Println("compiling file", path)
	contracts, err := core.Compile()
	if err != nil {
		log.Fatal(err, ": compilation error")
	} else {
		fmt.Println("compilation completed. saving binaries in build/ directory.")
	}

	//fmt.Println(contracts)
	contractName := new.ContractNameFromPath(contracts[0])
	//fmt.Println(contractName)
	exampleContract, err := new.Contract(contractName[0:len(contractName) - 4])
	if err != nil {
		fmt.Println(err)
	}
	if false {
		fmt.Println(exampleContract)
	}
}