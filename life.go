package main

import (
	"life/game"
	"log"
)

func main() {
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
