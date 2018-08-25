package create

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func Client(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}