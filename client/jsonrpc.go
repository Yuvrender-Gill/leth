package client

import (
	//"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
	"math/big"
)

type Response struct {
	Result string			`json:"result"`
}

type AccountResponse struct {
	Result []string			`json:"result"`
}

// this function gets the current block number by calling "eth_blockNumber"
func GetAccounts(url string) ([]string, error) {
	client := &http.Client{}
	var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	blockNumResp, err := client.Do(req)
	if err != nil {
       	return []string{}, err
	}
	defer blockNumResp.Body.Close()

	// print out response of eth_blockNumber
	//fmt.Println("response Status:", blockNumResp.Status)
	//fmt.Println("response Headers:", blockNumResp.Header)
	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)
	//fmt.Println("response Body:", string(blockNumBody))

	// parse json for result
	resp := new(AccountResponse)
	err = json.Unmarshal(blockNumBody, resp)
	if err != nil {
		return []string{}, err
	}

	return resp.Result, nil
}

// this function gets the current block number by calling "eth_blockNumber"
func GetBlockNumber(url string) (*big.Int, error) {
	client := &http.Client{}
	var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	blockNumResp, err := client.Do(req)
	if err != nil {
       	return nil, err
	}
	defer blockNumResp.Body.Close()

	// print out response of eth_blockNumber
	//fmt.Println("response Status:", blockNumResp.Status)
	//fmt.Println("response Headers:", blockNumResp.Header)
	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)
	//fmt.Println("response Body:", string(blockNumBody))

	// parse json for result
	resp := new(Response)
	err = json.Unmarshal(blockNumBody, resp)
	if err != nil {
		return nil, err
	}
	blockNumBytes, err := hex.DecodeString(resp.Result[2:])
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(blockNumBytes), nil
}