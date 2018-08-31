package create

import (
	"github.com/noot/leth/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Client(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Connection(network string) (*ethclient.Client, error) {
	file, err := core.ReadConfig()
	if err != nil {
		return nil, err
	}

	config, err := core.UnmarshalConfig(file)
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