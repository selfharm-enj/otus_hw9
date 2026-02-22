package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	Player1 Player
	Player2 Player
	Tui     bool
	Engine
}

func NewGame(tui bool) *Game {
	g := Game{}
	g.Player1 = Player{}
	g.Player2 = Player{}
	b := NewBoard()
	g.Board = b
	g.PlayerSide = White
	if tui {
		g.Tui = tui
		g.DrawBoard()
	}
	return &g
}

func (g *Game) ReadStdin() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("White Peace: ")
		sp, _ := reader.ReadString('\n')
		sp = strings.TrimSpace(sp)
		fmt.Println("White Target: ")
		ep, _ := reader.ReadString('\n')
		ep = strings.TrimSpace(ep)
		g.MakeMove(sp, ep)
		if g.Tui {
			g.DrawBoard()
		}
	}
}
