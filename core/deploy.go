package core

import (
	"fmt"
	"encoding/hex"
	"context"
	"math/big"
	"path/filepath"
	"io/ioutil"

	"github.com/noot/leth/logger"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

type Config struct {
	Networks map[string]Network 	`json:"networks"`
}

type Network struct {
	Url string						`json:"url,omitempty"`
	From string						`json:"from,omitempty"`
	Keystore string					`json:"keystore,omitempty"`
	Password string					`json:"password,omitempty"`
	GasPrice int64  				`json:"gasPrice,omitempty"`
	Id string						`json:"id,omitempty"`
}

func Deploy(client *ethclient.Client, network Network, contracts []string, keys *keystore.KeyStore) error {
	for _, contract := range contracts {
		err := deploy(client, network, contract, keys)
		if err != nil {
			logger.FatalError(fmt.Sprintf("could not deploy contract: %s", err))
		}
	}

	return nil
}

func deploy(client *ethclient.Client, network Network, contract string, keys *keystore.KeyStore) error {
	from := new(accounts.Account)
	if network.From != "" {
		from.Address = common.HexToAddress(network.From[2:])
	} else {
		logger.FatalError("no from address specified in config.json")
	}

	data, err := getBytecode(contract)
	if err != nil {	
		logger.Error(fmt.Sprintf("could not get bytecode for contract %s", contract))
		return err
	} 

	nonce, _ := client.PendingNonceAt(context.Background(), from.Address)
	gasPrice :=  big.NewInt(network.GasPrice)
	id, _ := new(big.Int).SetString(network.Id, 10)

	tx := types.NewContractCreation(nonce, big.NewInt(0), uint64(4600000), gasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, network.Password, tx, id)
	if err != nil {
		logger.Error(fmt.Sprintf("could not sign tx: %s", err))
		return err
	}

	txHash := txSigned.Hash()
	logger.Info(fmt.Sprintf("attempting to send tx %s to from account %s to deploy contract %s", txHash.Hex(), from.Address.Hex(), contract))

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		logger.Error(fmt.Sprintf("could not send tx %s", txHash))
		return err
	}

	return nil
}

func getBytecode(contract string) ([]byte, error) {
	path, _ := filepath.Abs("./build/" + contract)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}	

	hexString := fmt.Sprintf("%s", file)
	//fmt.Println(hexString)

	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return hexBytes, nil
}