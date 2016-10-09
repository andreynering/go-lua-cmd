package main

import (
	"fmt"
	"os"

	"github.com/yuin/gopher-lua"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Error: no argument given")
		os.Exit(1)
	}

	fileName := args[0]

	l := lua.NewState()
	defer l.Close()

	if err := l.DoFile(fileName); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
