# leth
golang tools for compiling, testing, and interacting with smart contracts

work in progress ~

# dependencies

solc 
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

# get 

`go get github.com/noot/leth`

`go get github.com/ethereum/go-ethereum`

# usage

`go build`

keystore setup: if you have an existing geth keystore, copy the keystore/ directory into this directory. 

to compile all contracts in contracts/ directory: `./leth compile`

to deploy contracts to network: `./leth deploy`

the default network to connect to is `default`. if you wish to connect to another network (as specified in config.json), use `./leth deploy --network network_name`

see config.json for example network configurations.
