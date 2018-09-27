package core

import (
	"math/big"
	"fmt"
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func WatchAllEvents(conn *ethclient.Client, contract common.Address, fromBlock *big.Int, logsChan chan []types.Log) error {
	filter := new(ethereum.FilterQuery)
	filter.FromBlock = fromBlock			


	contractArr := make([]common.Address, 1)
	contractArr = append(contractArr, contract)
	filter.Addresses = contractArr

	logs, err := conn.FilterLogs(context.Background(), *filter)
	if err != nil {
		fmt.Println(err)
	}

	if len(logs) != 0 {
		logsChan <- logs
	}

	return nil
}