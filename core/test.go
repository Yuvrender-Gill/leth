package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
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