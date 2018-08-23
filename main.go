package main

import (
	"fmt"

	"github.com/noot/leth/core"
)

func main() {
	fmt.Println("welcome to leth")

	//client := leth.Dial("http://localhost:8545")

	var input string
	fmt.Println("enter contract to compile:")
	fmt.Scanln(&input)
	err := core.Compile(input)
	if err != nil {
		fmt.Println(err, ": compilation error. try removing build/ directory.")
	}
}