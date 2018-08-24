package main

import (
	"fmt"
	"log"
	"flag"
	"os"

	"github.com/noot/leth/core"
	"github.com/noot/leth/new"
)

func main() {
	//client := leth.Dial("http://localhost:8545")

	help := flag.Bool("help", false, "print out command-line options")

	flag.Parse() 
	if *help {
		fmt.Println("\t\x1b[93mleth help\x1b[0m")
		fmt.Println("\tleth compile: compile all contracts in contracts/ directory")
		os.Exit(0)
	}

	compileCommand := flag.NewFlagSet("compile", flag.ExitOnError)

	// subcommands
	// ./b
	if len(os.Args) > 1 {
		switch os.Args[1]{
			case "compile":
				compileCommand.Parse(os.Args[2:])
			default:
		}
	} else {
		os.Exit(0)
	}

	contracts := []string{}
	var err error
	if compileCommand.Parsed() {
		//contractArgs := compileCommand.Args()
		contracts, err = core.Compile()
		if err != nil {
			log.Fatal(err, ": compilation error")
		} else {
			fmt.Println("compilation completed. saving binaries in build/ directory.")
		}
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