package core

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
	Result string					`json:"result"`
}

type ReceiptResponse struct {
	Result Receipt 					`json:"result"`
}

type AccountResponse struct {
	Result []string					`json:"result"`
}

type Receipt struct {
	 TransactionHash string 		`json:"transactionHash"`
     TransactionIndex string 		`json:"transactionIndex"`
     BlockNumber string 			`json:"blockNumber"`
     BlockHash string 				`json:"blockHash"`
     CumulativeGasUsed string 		`json:"cumulativeGasUsed"`
     GasUsed string 				`json:"gasUsed"`
     ContractAddress string 		`json:"contractAddress"`
     Logs []string 					`json:"logs"`
     LogsBloom string 				`json:"logsBloom"`
     Status string 					`json:"status"`
}

// this function gets the accounts in the client by calling "eth_accounts"
func GetAccounts(url string) ([]string, error) {
	client := &http.Client{}
	var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	accountResp, err := client.Do(req)
	if err != nil {
       	return []string{}, err
	}
	defer accountResp.Body.Close()

	accountBody, _ := ioutil.ReadAll(accountResp.Body)

	resp := new(AccountResponse)
	err = json.Unmarshal(accountBody, resp)
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

	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)

	resp := new(Response)
	err = json.Unmarshal(blockNumBody, resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Result) == 3 {
		num, _ := new(big.Int).SetString(resp.Result[2:], 16)
		return num, nil
	}

	blockNumBytes, err := hex.DecodeString(resp.Result[2:])
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(blockNumBytes), nil
}

func GetTransactionReceipt(txHash string, url string) (Receipt, error) {
	client := &http.Client{}
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":1}`
    jsonBytes := []byte(jsonStr)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { 
    	return Receipt{}, err 
    }
    req.Header.Set("Content-Type", "application/json")
    txResp, err := client.Do(req)
	
	receiptBody, _ := ioutil.ReadAll(txResp.Body)

	resp := new(ReceiptResponse)
	err = json.Unmarshal(receiptBody, resp)
	if err != nil {
		return Receipt{}, err
	}

	return resp.Result, nil
}

/*
params: [{
  "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
  "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567", //not needed for new contract
  "gas": "0x76c0", // 30400
  "gasPrice": "0x9184e72a000", // 10000000000000
  "value": "0x9184e72a", // 2441406250
  "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
}]
*/
func SendTransaction(txData string, url string) (string, error) {
	client := &http.Client{}
    jsonStr := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[` + txData + `],"id":1}`
    jsonBytes := []byte(jsonStr)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return "", err }
    req.Header.Set("Content-Type", "application/json")
    txResp, err := client.Do(req)
	
	txHashBody, _ := ioutil.ReadAll(txResp.Body)

	resp := new(Response)
	err = json.Unmarshal(txHashBody, resp)
	if err != nil {
		return "", err
	}

	return resp.Result, nil
}