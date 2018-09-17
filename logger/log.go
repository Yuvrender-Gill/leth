package logger

import (
	"fmt"
	"os"
)

// todo: change these functions to allow multiple parameters
// and format them inside the function

func Info(out string) {
	fmt.Println("\x1b[92minfo:\x1b[0m", out)
}

func Warn(out string) {
	fmt.Println("\x1b[93mwarn:\x1b[0m", out)
}

func Error(out string) {
	fmt.Println("\x1b[91merror:\x1b[0m", out)
}

func FatalError(out string) {
	fmt.Println("\x1b[91merror:\x1b[0m", out)	
	os.Exit(1)
}

func CompilerWarn(out string) {
	fmt.Println("\x1b[90m", out, "\x1b[0m")
}

func CompilerError(out string) {
	fmt.Println("\x1b[91m", out, "\x1b[0m")
}