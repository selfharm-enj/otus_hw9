package model

import (
	"fmt"
	"strings"
)

type Side int

const (
	White Side = iota
	Black
)

type Board [][]rune

func NewBoard() Board {
	board := [][]rune{
		{'♜', '♞', '♝', '♛', '♚', '♝', '♞', '♜'},
		{'♟', '♟', '♟', '♟', '♟', '♟', '♟', '♟'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'♙', '♙', '♙', '♙', '♙', '♙', '♙', '♙'},
		{'♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'},
	}
	return board
}

func (b *Board) DrawBoard() {
	out := strings.Builder{}
	out.WriteString("Player 2\n")
	for i, r := range *b {
		s := fmt.Sprintf("%d ", 8-i)
		out.WriteString(s)
		for _, v := range r {
			s := fmt.Sprintf("%s ", string(v))
			out.WriteString(s)
		}
		out.WriteString("\n")
	}
	out.WriteString("  a b c d e f g h\n")
	out.WriteString(" Player 1\n")
	fmt.Println(out.String())
}
