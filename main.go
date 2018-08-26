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
	"github.com/noot/leth/jsonrpc"
	"github.com/noot/leth/logger"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	// flags
	help := flag.Bool("help", false, "print out command-line options")

	// compile subcommand and flags
	compileCommand := flag.NewFlagSet("compile", flag.ExitOnError)
	bindFlag := compileCommand.Bool("bind", true, "specify whether to create bindings for contracts while compiling")

	// deploy subcommand and flags
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
		compile(*bindFlag)
		os.Exit(0)	
	} 

	if deployCommand.Parsed() {
		deploy(*network)
		os.Exit(0)
	}
}

func bind() (error) {
	//fmt.Println(contracts)
	err := create.Bindings()
	if err != nil {
		logger.FatalError(fmt.Sprintf("could not create bindings: %s", err))
	} else {
		logger.Info("generation of bindings completed. saving bindings in bind/ directory.")
	}
	return nil
} 

func compile(bindFlag bool) ([]string) {
	contracts, err := core.Compile()
	if err != nil {
		logger.FatalError(fmt.Sprintf("compilation error: %s", err))
	} else {
		logger.Info("compilation completed. saving binaries in build/ directory.")
	}
	if bindFlag {
		err = bind()
		if err != nil {
			logger.FatalError(fmt.Sprintf("could not generate bindings: %s", err))
		}
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
		compile(false) // don't need to generate bindings for deployment
	}

	files, err := core.SearchDirectory("./build")
	if err != nil {
		log.Fatal(err)
	} else if len(files) < 2 {
		logger.Info("build/ directory empty. compiling contracts...")
		compile(false)
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

	// read config file
	file, err := readConfig()
	if err != nil {
		logger.FatalError("no config.json file found.")
		os.Exit(1)
	}

	config, err := unmarshalConfig(file)
	if err != nil {
		logger.FatalError(fmt.Sprintf("could not read config.json: %s", err))
	}

	ntwk := config.Networks[network]

	// dial client for network
	//ntwk := new(core.Network)
	ethclient, err := create.Client(ntwk.Url)
	if err != nil {
		logger.FatalError("cannot dial client; likely incorrect url in config.json")
	}

	logger.Info(fmt.Sprintf("deploying %s to network %s", names, network))

	if ntwk.Keystore == "" {
		accounts, err := jsonrpc.GetAccounts(ntwk.Url)
		if err != nil {
			logger.FatalError(fmt.Sprintf("unable to get accounts from client url: %s", err))
		}
		//logger.Info(fmt.Sprintf("accounts: %s", accounts))
		printAccounts(accounts)

		if ntwk.From == "" {
			ntwk.From = accounts[0]
		}

		err = core.DeployTestRPC(ntwk, names)
		if err != nil {
			logger.FatalError("could not deploy contracts.")
		}
	} else {
		ks := newKeyStore(ntwk.Keystore)
		ksaccounts := ks.Accounts()
		printKeystoreAccounts(ksaccounts)
		err = core.Deploy(ethclient, ntwk, names, ks)
		if err != nil {
			logger.FatalError("could not deploy contracts.")
		}
	}

	blockNum, err := jsonrpc.GetBlockNumber(ntwk.Url)
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}
	logger.Info(fmt.Sprintf("block number: %s", blockNum))
}

func readConfig() ([]byte, error) {
	path, _ := filepath.Abs("./config.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}	
	return file, nil
}

func unmarshalConfig(file []byte) (*core.Config, error) {
	conf := new(core.Config)
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

func printAccounts(accounts []string) {
	for i, account := range accounts {
		logger.Info(fmt.Sprintf("account %d: %s", i, account))
	}
}

func printKeystoreAccounts(accounts []accounts.Account) {
	for i, account := range accounts {
		logger.Info(fmt.Sprintf("account %d: %s", i, account.Address.Hex()))
	}
}