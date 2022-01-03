package main

import (
	"fmt"
	"github.com/metelang/mete-repl/repl"
	"os"
	"os/user"
)

func main() {
	osUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Mete Lang REPL\n", osUser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
