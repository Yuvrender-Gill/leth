package migrations

import (
	"github.com/ChainSafeSystems/leth/core"
)

func Migrate() {
	err := core.Migrate("Example.sol")
}