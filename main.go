package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"path"
	"encoding/json"
	"io/ioutil"

	"github.com/noot/leth/core"
	"github.com/noot/leth/create"
	"github.com/noot/leth/jsonrpc"
	"github.com/noot/leth/logger"
	"github.com/noot/leth/test"

	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	// flags
	help := flag.Bool("help", false, "print out command-line options")

	// init subcommand
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)

	// bind subcommand 
	bindCommand := flag.NewFlagSet("bind", flag.ExitOnError)

	// compile subcommand and flags
	compileCommand := flag.NewFlagSet("compile", flag.ExitOnError)
	bindFlag := compileCommand.Bool("bind", true, "specify whether to create bindings for contracts while compiling")

	// deploy subcommand and flags
	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
	network := deployCommand.String("network", "default", "specify network to connect to (configured in config.json)")

	// test subcommand
	testCommand := flag.NewFlagSet("test", flag.ExitOnError)

	flag.Parse() 
	if *help {
		fmt.Println("\t\x1b[93mleth help\x1b[0m")
		fmt.Println("\tleth bind: create go bindings for all contracts in contracts/ directory and save in bind/")
		fmt.Println("\tleth compile: compile all contracts in contracts/ directory and save results in build/. `compile` will automatically execute `bind`; to compile with out binding, use --bind=false")
		fmt.Println("\tleth deploy: deploy all contracts in contracts/ directory and save results of deployment in deployed/. specify network name with `--network NETWORK_NAME`. if no network is provided, leth will connect to the default network as specified in config.json")
		fmt.Println("\tleth test: run tests in test/ directory")
		os.Exit(0)
	} 

	// subcommands
	if len(os.Args) > 1 {
		switch os.Args[1]{
			case "init":
				initCommand.Parse(os.Args[2:])
			case "bind":
				bindCommand.Parse(os.Args[2:])
			case "compile":
				compileCommand.Parse(os.Args[2:])
			case "deploy":
				deployCommand.Parse(os.Args[2:])
			case "test":
				testCommand.Parse(os.Args[2:])
			default:
				// continue
		}
	} else {
		os.Exit(0)
	}

	if initCommand.Parsed() {
		lethInit()
		os.Exit(0)
	}

	if bindCommand.Parsed() {
		bind()
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

	if testCommand.Parsed() {
		testrun()
		os.Exit(0)
	}
}

func lethInit() {
	files, err := core.SearchDirectory("./")
	if len(files) > 1 {
		logger.FatalError("cannot init in non-empty directory.")
	}
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}

	os.Mkdir("./contracts", os.ModePerm)
	os.Mkdir("./keystore", os.ModePerm)
	os.Mkdir("./test", os.ModePerm)

	jsonStr, err := json.MarshalIndent(core.DefaultConfig, "", "\t")
	if err != nil {
		logger.Error(fmt.Sprintf("%s", err))
	}

	ioutil.WriteFile("./config.json", jsonStr, os.ModePerm)
}

func bind() {
	//fmt.Println(contracts)
	err := create.Bindings()
	if err != nil {
		logger.FatalError(fmt.Sprintf("could not create bindings: %s", err))
	} else {
		logger.Info("generation of bindings completed. saving bindings in bind/ directory.")
	}
} 

func compile(bindFlag bool) ([]string) {
	contracts, err := core.Compile()
	if err != nil {
		logger.FatalError(fmt.Sprintf("compilation error: %s", err))
	} else {
		logger.Info("compilation completed. saving binaries in build/ directory.")
	}
	if bindFlag {
		bind()
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

	names := core.GetContractNames(contracts)

	// read config file
	file, err := core.ReadConfig()
	if err != nil {
		logger.FatalError("no config.json file found.")
		os.Exit(1)
	}

	config, err := core.UnmarshalConfig(file)
	if err != nil {
		logger.FatalError(fmt.Sprintf("could not read config.json: %s", err))
	}

	ntwk := config.Networks[network]
	ntwk.Name = network

	// dial client for network
	//ntwk := new(core.Network)
	client, err := create.Client(ntwk.Url)
	if err != nil {
		logger.FatalError("cannot dial client; likely incorrect url in config.json")
	}

	//logger.Info(fmt.Sprintf("deploying %s to network %s", names, network))

	if ntwk.Name == "testrpc" || ntwk.Name == "ganache" || ntwk.Name == "ganache-cli" {
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
		err = core.Deploy(client, ntwk, names, ks)
		if err != nil {
			logger.FatalError("could not deploy contracts.")
		}
	}

	// blockNum, err := jsonrpc.GetBlockNumber(ntwk.Url)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("%s", err))
	// }
	// logger.Info(fmt.Sprintf("block number: %s", blockNum))
}

func testrun() {
	test.TestExample()
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