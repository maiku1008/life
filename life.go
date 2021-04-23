package main

import (
	"flag"
	"life/game"
	"log"
)

func main() {
	initialCells := flag.Int("c", 70000, "The initial number of living cells")
	flag.Parse()

	if err := game.Start(*initialCells); err != nil {
		log.Fatal(err)
	}
}
