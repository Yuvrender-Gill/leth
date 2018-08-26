package core

import (
	"os"
	"encoding/hex"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

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

func Exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func SearchDirectory(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}