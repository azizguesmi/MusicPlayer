package main

import (
	"MuisicPlayer/ui"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	err := ui.Run(args)
	if err != nil {
		fmt.Println(err)
	}
}
