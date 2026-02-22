package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Engine struct {
	Board
	PlayerSide Side
}

func (e *Engine) SwitchSide() {
	switch e.PlayerSide {
	case White:
		e.PlayerSide = Black
	case Black:
		e.PlayerSide = White
	}
}

func (e *Engine) parsePosition(pos string) [2]int {
	var out [2]int
	col, _ := strconv.Atoi(pos[1:])
	row := []rune(pos[:1])

	out[0] = 8 - col
	out[1] = int(row[0]) - 'a'
	return out
}

func (e *Engine) MakeMove(startPos, endPos string) {
	sp := e.parsePosition(startPos)
	ep := e.parsePosition(endPos)

	sr := sp[0]
	sc := sp[1]
	er := ep[0]
	ec := ep[1]

	piece := e.Board[sr][sc]
	e.Board[sr][sc] = '.'
	e.Board[er][ec] = piece
}

func (e *Engine) PseudoRandomGame() {
	moves := map[int][]string{
		1:  {"e2-e4", "c7-c5"},
		2:  {"g1-f3", "d7-d6"},
		3:  {"d2-d4", "c5-d4"},
		4:  {"f3-d4", "g8-f6"},
		5:  {"b1-c3", "a7-a6"},
		6:  {"f1-g5", "e7-e6"},
		7:  {"f2-f4", "f8-e7"},
		8:  {"d1-f3", "b8-d7"},
		9:  {"e1-g1", "b7-b5"},
		10: {"a2-a3", "f8-b7"},
		11: {"g5-f6", "d7-f6"},
		12: {"g2-g4", "b5-b4"},
		13: {"c1-b1", "b4-c3"},
		14: {"b2-c3", "h7-h6"},
		15: {"g4-g5", "f6-e8"},
		16: {"f3-h5", "e8-f6"},
		17: {"g5-f6", "e7-f6"},
		18: {"h1-g1", "g7-g6"},
		19: {"h5-h4", "f6-f7"},
		20: {"e4-e5", "d6-e5"},
		21: {"f4-e5", "f7-e5"},
		22: {"d4-e6", "f7-e6"},
		23: {"d1-e6", "e8-f7"},
		24: {"e1-d7", "f7-g8"},
		25: {"d7-f7", "g8-f7"},
		26: {"e6-f7", "f7-f8"},
		27: {"f7-f8", "END-END"},
	}
	for k := range len(moves) {
		for _, v := range moves[k+1] {
			s := strings.Split(v, "-")
			sp := s[0]
			ep := s[1]
			if sp == "END" {
				break
			}
			time.Sleep(3 * time.Second)
			e.MakeMove(sp, ep)
			e.DrawBoard()
			fmt.Printf("Last move: %s\n", ep)
			time.Sleep(3 * time.Second)
		}
	}
}
