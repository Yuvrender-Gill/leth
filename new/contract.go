package new

import (
	"path/filepath"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func ContractFromPath(pathStr string) (*abi.ABI, error) {
	path, _ := filepath.Abs(pathStr)
	return contract(path)
}

// param: name of the contract (without extension)
func Contract(contract string) (*abi.ABI, error) {
	path, _ := filepath.Abs('../build/' + contract + '.abi')
	return contract(path)
}

func contract(pathStr string) (*abi.ABI, error) {
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