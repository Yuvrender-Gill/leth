package core

import (
	"fmt"
	"os/exec"
	"os"
)

func Compile(contract string) error { 
	buildexists, err := exists("build/")
	fmt.Println(buildexists)
	if !buildexists {
		os.Mkdir("./build", os.ModePerm)
	}

	// generate bytecode
    app := "solc"
    arg0 := "--bin"
    arg1 := contract
    arg2 := "-o"
    arg3 := "build/"


    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err := cmd.Output()

    if err != nil {
        return err
    }

    print(string(stdout))

    // generate abi
    app = "solc"
    arg0 = "--abi"
    arg1 = contract
    arg2 = "-o"
    arg3 = "build/"

    cmd = exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err = cmd.Output()

    if err != nil {
        return err
    }

    print(string(stdout))	
    return nil
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}