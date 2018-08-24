package main

import (
	"fmt"

	"github.com/noot/leth/core"
)

func main() {
	//client := leth.Dial("http://localhost:8545")

	// var input string
	// fmt.Println("enter contract to compile (.sol file needs to be in contracts/)")
	// fmt.Scanln(&input)
	// path, _ := filepath.Abs("./contracts/" + input + ".sol")
	// fmt.Println("compiling file", path)
	err := core.Compile()
	if err != nil {
		fmt.Println(err, ": compilation error. try removing build/ directory.")
	} else {
		fmt.Println("compilation completed. saving binaries in build/ directory.")
	}
}