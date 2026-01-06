package game

import (
	player "goat-lang/src/models/Player"
)

type GameState int8

const (
	PrepareGame GameState = iota
	Running
	InRound
	EndGame
)

type Game struct {
	Name        string
	Developer   string
	PlayersList []player.Player
	Tour        int8
	GameState   GameState
	Version     string
}
