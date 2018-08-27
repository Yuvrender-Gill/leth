# leth
golang tools for compiling, deploying, testing, and interacting with smart contracts

work in progress ~

currently compiles + deploys. need to get a nice flow for testing down.

# dependencies

solc
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

go-ethereum

`go get github.com/ethereum/go-ethereum`

abigen
```
cd $GOPATH/src/github.com/ethereum/go-ethereum
godep go install ./cmd/abigen
```

# get 

`go get github.com/noot/leth`

# usage

```
cd $GOPATH/src/github.com/noot/leth
go build
go install
```

keystore setup: if you have an existing geth keystore, copy the keystore/ directory into this directory.

`leth init` to initialize setup

`leth compile` to compile all contracts in contracts/ directory

`leth deploy` to deploy contracts to network

the default network to connect to is `default`. if you wish to connect to another network (as specified in config.json), use `leth deploy --network network_name`

see config.json for example network configurations.
