package core 

import (
	"path/filepath"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func Contract(pathStr string) (*abi.ABI, error) {
	path, _ := filepath.Abs(pathStr)
	file, err := ioutil.ReadFile(path)
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