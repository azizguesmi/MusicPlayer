package main

import (
	"MuisicPlayer/ui"
	"os"
)

func main() {
	args := os.Args[1:]
	ui.Run(args)
}
