package main

import (
	"flag"

	"github.com/selfharm-enj/otus_hw9/internal/model"
)

func main() {
	gameMode := flag.String("m", "human", "Game mode: `pc` or `human`")
	flag.Parse()
	game := model.NewGame(true)
	switch *gameMode {
	case "human":
		game.ReadStdin()
	case "auto":
		game.PseudoRandomGame()
	}
}
