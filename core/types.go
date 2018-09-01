package core

var DefaultNetwork = Network {
	Name: "default",
	Url: "http://localhost:8545",
	From: "0xADDRESS",
	Keystore: "./keystore",
	Password: "",
	Gas: 100000000,
	GasPrice: 4700000,
	Id: "",
}

var DefaultConfig = Config {
	Networks: map[string]Network{"default": DefaultNetwork},
}

type Config struct {
	Networks map[string]Network 	`json:"networks"`
}

type Network struct {
	Name string						`json:"name,omitempty"`
	Url string						`json:"url,omitempty"`
	From string						`json:"from,omitempty"`
	Keystore string					`json:"keystore,1omitempty"`
	Password string					`json:"password,omitempty"`
	Gas int64 						`json:"gas,omitempty"`
	GasPrice int64  				`json:"gasPrice,omitempty"`
	Id string						`json:"id,omitempty"`
}

type Transaction struct {
	From string 					`json:"from"`	
	To string 						`json:"to,omitempty"`
	Gas string 						`json:"gas,omitempty"`
	GasPrice string 				`json:"gasPrice,omitempty"`
	Value string 					`json:"value,omitempty"`
	Data string 					`json:"data,omitempty"`
}