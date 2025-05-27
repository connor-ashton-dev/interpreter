package main

import (
	"fmt"
	"os"

	"github.com/connor-ashton-dev/crafting_interpreters/lox"
)

func main() {
	args := os.Args[1:]
	l := lox.New()

	if len(args) > 1 {
		fmt.Println("Usage: go run main.go [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		l.RunFile(args[0])
	} else {
		l.RunPrompt()
	}
}
