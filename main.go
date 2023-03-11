package main

import (
	"os"

	"oulanbator/game"
)

func main() {
	input := os.Stdin
	output := os.Stdout

	g := game.NewGame(input, output)

	g.Start()
	g.Play()
}
