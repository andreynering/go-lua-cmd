package main

import (
	"fmt"
	"os"

	"github.com/Shopify/go-lua"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error: not argument given")
		os.Exit(1)
	}

	file := args[0]

	l := lua.NewState()
	lua.OpenLibraries(l)

	if err := lua.DoFile(l, file); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
