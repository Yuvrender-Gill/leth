package core

import (
	"fmt"
	"encoding/json"
	"context"
	"math/big"

	"github.com/noot/leth/logger"
	"github.com/noot/leth/jsonrpc"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

func DeployTestRPC(network Network, contracts []string) error {
	for _, contract := range contracts {
		logger.Info(fmt.Sprintf("deploying %s.sol to network %s", contract, network.Name))
		err := deployTestRPC(network, contract)
		if err != nil {
			logger.FatalError(fmt.Sprintf("could not deploy contracts: %s", err))
		}
	}
	return nil
}

func deployTestRPC(network Network, contract string) error {
	bytecode, err := getBytecode(contract)
	if err != nil {	
		logger.Error(fmt.Sprintf("could not get bytecode for contract %s", contract))
		return err
	} 

	tx := Transaction{}
	tx.Data = fmt.Sprintf("%x", bytecode)
	tx.From = network.From
	tx.GasPrice = fmt.Sprintf("%x", network.GasPrice)
	tx.Gas = fmt.Sprintf("%x", network.Gas)

	txBytes, err := json.Marshal(tx)
	if err != nil {
		logger.Error("could not create TestRPC transaction")
		return err
	}
	data := fmt.Sprintf("%s", txBytes)
	//fmt.Println(data)

	txHash, err := jsonrpc.SendTransaction(data, network.Url)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("contract creation at tx %s", txHash))
	return nil
}

func Deploy(client *ethclient.Client, network Network, contracts []string, keys *keystore.KeyStore) error {
	for _, contract := range contracts {
		logger.Info(fmt.Sprintf("deploying %s.sol to network %s", contract, network.Name))
		err := deploy(client, network, contract, keys)
		if err != nil {
			logger.Error(fmt.Sprintf("could not deploy contracts: %s", err))
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
	logger.Info(fmt.Sprintf("attempting to send tx %s to from account %s to deploy contract %s.sol", txHash.Hex(), from.Address.Hex(), contract))

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		logger.Error(fmt.Sprintf("could not send tx %s", txHash.Hex()))
		return err
	}

	//logger.Info(fmt.Sprintf("contract creation at tx %s", txHash.Hex()))

	waitOnPending(client, txHash)

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		logger.Error(fmt.Sprintf("could not deploy contract %s.sol", contract))
		return nil
	}

	contractAddr := receipt.ContractAddress
	logger.Info(fmt.Sprintf("contract deployed at address %s", contractAddr.Hex()))
	logger.Info(fmt.Sprintf("gas used to deploy contract %s.sol: %s", contract, receipt.GasUsed))
	return nil
}

func waitOnPending(client *ethclient.Client, txHash common.Hash) (*types.Transaction) {
	for {
		tx, pending, _ := client.TransactionByHash(context.Background(), txHash)
		if !pending { 
			return tx 
		}
	}
}