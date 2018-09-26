package migrations

import (
	"fmt"

	"github.com/ChainSafeSystems/leth/core"
)

func Migrate() {
	err := core.Migrate("testnet", "Token")
	if err != nil {
		fmt.Println("could not deploy Token.sol to testnet")
	}

	// err = core.Migrate("testnet", "ExampleLib")
	// if err != nil {
	// 	fmt.Println("could not deploy ExampleLib.sol to testnet")
	// }
}