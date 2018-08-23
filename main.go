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
	core.Compile(input)
}