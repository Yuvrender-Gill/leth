package core

import (
	"os/exec"
)

func Compile(contract string) {
    app := "solc"

    arg0 := "--abi"
    arg1 := contract
    arg2 := "-o"
    arg3 := "build/"

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}