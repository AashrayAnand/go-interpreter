package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, Welcome to Aashray's go interpreter!\n", user.Username)
	fmt.Printf("Type any legal commands, you may have to read through" +
		"my code to guess the language semantics >:)\n")
	repl.Start(os.Stdin, os.Stdout)
}
