package models

type Game struct {
	Name      string
	Developer string
	Players   []Player
	Tour      int
	Version   string
}