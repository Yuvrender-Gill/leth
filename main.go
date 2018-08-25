package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"path"
	"path/filepath"
	"encoding/json"
	"io/ioutil"

	"github.com/noot/leth/core"
	"github.com/noot/leth/create"
	"github.com/noot/leth/logger"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Config struct {
	Networks map[string]Network 	`json:"networks"`
}

type Network struct {
	Url string						`json:"url"`
	From string						`json:"from"`
	Keystore string					`json:"keystore"`
	Password string					`json:"password"`
}

func main() {
	//client := leth.Dial("http://localhost:8545")

	help := flag.Bool("help", false, "print out command-line options")

	compileCommand := flag.NewFlagSet("compile", flag.ExitOnError)
	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
	network := deployCommand.String("network", "default", "specify network to connect to (configured in config.json)")

	flag.Parse() 
	if *help {
		fmt.Println("\t\x1b[93mleth help\x1b[0m")
		fmt.Println("\tleth compile: compile all contracts in contracts/ directory")
		os.Exit(0)
	} 

	// subcommands
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
		deploy(*network)
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

// set up deployment to network
// compile, read config, dial network, set up accounts
func deploy(network string) {
	// compilation of contracts, if needed
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
		name := create.ContractNameFromPath(contract)
		names = append(names, name)
	}

	logger.Info(fmt.Sprintf("deploying %s to network %s", names, network))

	// note: we want to dial the client and read accounts from there
	// we can use this instead if they wish to specify a keystore for a network
	ks := newKeyStore("./keystore")
	ksaccounts := ks.Accounts()
	for i, account := range ksaccounts {
		fmt.Println("account", i, ":", account.Address.Hex())
	}

	// read config file
	file, err := readConfig()
	if err != nil {
		logger.Error("no config.json file found.")
		os.Exit(1)
	}

	config, err := unmarshalConfig(file)
	if err != nil {
		logger.Error(fmt.Sprintf("could not read config.json: %s", err))
	}

	// dial client for network
	client, err := create.Client(config.Networks[network].Url)
	if err != nil {
		logger.Error("cannot dial client; likely incorrect url in config.json")
	}
	fmt.Println(client)
}


func readConfig() ([]byte, error) {
	path, _ := filepath.Abs("./config.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}	
	return file, nil
}

func unmarshalConfig(file []byte) (*Config, error) {
	conf := new(Config)
	err := json.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func newKeyStore(path string) (*keystore.KeyStore) {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}