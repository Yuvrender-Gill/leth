package core

import (
	"errors"
	"path/filepath"
	"io/ioutil"
	"encoding/json"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func NewConnection(network string) (*ethclient.Client, error) {
	file, err := ReadConfig()
	if err != nil {
		return nil, err
	}

	config, err := UnmarshalConfig(file)
	if err != nil {
		return nil, err
	}

	url := config.Networks[network].Url
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	
	return client, nil
}

func ContractAddress(contract string, network string) (common.Address, error) {
	var address common.Address

	path, _ := filepath.Abs("./deployed/" + network + ".json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return address, err
	}	

	var deployed map[string]string
	err = json.Unmarshal(file, &deployed)
	if err != nil {
		return address, err
	}

	addressString := deployed[contract]
	if addressString == "" {
		return address, errors.New("contract has not been deployed to network.")
	}

	address = common.HexToAddress(addressString)
	return address, nil
}