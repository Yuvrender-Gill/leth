package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"path"
	"path/filepath"
	"io/ioutil"

	"github.com/noot/leth/core"
	"github.com/noot/leth/new"
	"github.com/noot/leth/logger"

	"github.com/ethereum/go-ethereum/accounts/keystore"
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
	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)

	// subcommands
	// ./b
	if len(os.Args) > 1 {
		switch os.Args[1]{
			case "compile":
				compileCommand.Parse(os.Args[2:])
			case "deploy":
				deployCommand.Parse(os.Args[2:])
			default:
				// continue
		}
	} else {
		os.Exit(0)
	}

	if compileCommand.Parsed() {
		//contractArgs := compileCommand.Args()
		compile()
		os.Exit(0)	
	} 

	if deployCommand.Parsed() {
		deploy()
		os.Exit(0)
	}

	
	/*
	// initialize new contract object	
	contractName := new.ContractNameFromPath(contracts[0])
	exampleContract, err := new.Contract(contractName[0:len(contractName) - 4])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exampleContract)
	*/
}

func compile() ([]string) {
	contracts, err := core.Compile()
	if err != nil {
		log.Fatal(err, ": compilation error")
	} else {
		logger.Info("compilation completed. saving binaries in build/ directory.")
	}
	return contracts
}

func deploy() {
	contracts := []string{}
	buildexists, err := core.Exists("build/")
	if !buildexists {
		logger.Info("build/ directory not found. compiling contracts...")
		compile()
	}

	files, err := core.SearchDirectory("./build")
	if err != nil {
		log.Fatal(err)
	} else if len(files) < 2 {
		logger.Info("build/ directory empty. compiling contracts...")
		compile()
		files, err = core.SearchDirectory("./build")
	} else {
		for _, file := range files {
			if(path.Ext(file) == ".bin") {
				contracts = append(contracts, file)
			}
		}
	}

	names := []string{}

	for _, contract := range contracts {
		name := new.ContractNameFromPath(contract)
		names = append(names, name)
	}

	logger.Info(fmt.Sprintf("deploying %s", names))

	ks := newKeyStore("./keystore")
	ksaccounts := ks.Accounts()
	for i, account := range ksaccounts {
		fmt.Println("account", i, ":", account.Address.Hex())
	}

	config, err := readConfig()
	if err != nil {
		logger.Error("no config.json file found.")
		os.Exit(1)
	}
	fmt.Println(string(config))
}

func readConfig() ([]byte, error) {
	path, _ := filepath.Abs("./config.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}	
	return file, nil
}

func newKeyStore(path string) (*keystore.KeyStore) {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}