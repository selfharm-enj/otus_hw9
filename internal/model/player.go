package model

type Player struct {
	Name string
	Elo  int
}

// Default bot player
var Bot = Player{
	Name: "Bot Player",
	Elo:  0,
}
