package core

import (
	"path/filepath"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func ContractFromPath(pathStr string) (*abi.ABI, error) {
	path, _ := filepath.Abs(pathStr)
	fmt.Println(path)
	return contractAbi(path)
}

// param: name of the contract (without extension)
func Contract(contract string) (*abi.ABI, error) {
	path, err := filepath.Abs("./build/" + contract + ".abi")
	if err != nil {
		return nil, err
	}
	contractabi, err := contractAbi(path)
	return contractabi, err
}

func contractAbi(pathStr string) (*abi.ABI, error) {
	file, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return nil, err
	}

	contractabi := new(abi.ABI)
	err = contractabi.UnmarshalJSON(file)
	if err != nil {
		return nil, err
	}
	return contractabi, nil	
}

func ContractNameFromPath(path string) string {
	return filepath.Base(path)
}