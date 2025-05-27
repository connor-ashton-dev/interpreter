package lox

import (
	"bufio"
	"fmt"
	"os"

	"github.com/connor-ashton-dev/crafting_interpreters/scanner"
)

type Lox struct {
	Source   string
	HadError bool
}

// The run function is the entry point for the scanner.
// It takes the source code as a string and returns a list of tokens.
func (l *Lox) run() {
	s := scanner.New(l.Source, l.errorHandler)
	tokens := s.ScanTokens()
	for _, tok := range tokens {
		fmt.Println(tok)
	}
}

func New() *Lox {
	return &Lox{
		Source: "",
	}
}

// RunFile reads the file at the given path and runs the scanner on the contents
func (l *Lox) RunFile(path string) {
	bytes, error := os.ReadFile(path)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return
	}

	l.Source = string(bytes)
	l.run()
}

func (l *Lox) RunPrompt() {
	fmt.Print("> ")

	var line string
	scanner := bufio.NewReader(os.Stdin)

	line, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	l.Source = line
	l.run()
}

func (l *Lox) errorHandler(line int, message string) {
	fmt.Printf("[line %d] Error: %s\n", line, message)
	l.HadError = true
}
