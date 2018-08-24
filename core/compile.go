package core

import (
	"fmt"
	"path"
	"path/filepath"
	"os/exec"
	"os"
	"strings"
)

func Compile() ([]string, error) {
	buildexists, err := exists("build/")
	if buildexists {
		os.RemoveAll("./build")
	}
	os.Mkdir("./build", os.ModePerm)

	dir, _ := filepath.Abs("contracts/")
	files := []string{}

	files, err = searchDirectory(dir, files)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
		return nil, err
	}

	contracts := []string{}

	//fmt.Println(files)
	for _, file := range files {
		if(path.Ext(file) == ".sol") {
			err = compile(file)
			if err != nil {
				return nil, err
			}
			contracts = append(contracts, file)
		}
	}

	return contracts, nil
}

func searchDirectory(dir string, files []string) ([]string, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		//fmt.Printf("visited file: %q\n", path)
		return nil
	})

	if err != nil {
		//fmt.Printf("error walking the path %q: %v\n", dir, err)
		return nil, err
	}

	//fmt.Println(files)
	return files, nil
}

func compile(contract string) error { 
	fmt.Println("compiling", contract)

	// generate bytecode
    app := "solc"
    arg0 := "--bin"
    arg1 := contract
    arg2 := "-o"
    arg3 := "build/"

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err := cmd.CombinedOutput()

    out := string(stdout)
    if strings.Contains(out, "Warning") {
	    fmt.Println("\x1b[93m", out, "\x1b[0m")
	} else if strings.Contains(out, "Error") {
	    fmt.Println("\x1b[91m", out, "\x1b[0m")
	} else {
		fmt.Println(out)
	}

    if err != nil {
        return err
    }

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

    //print(string(stdout))	
    return nil
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}