package game

import (
	player "goat-lang/src/models/Player"
)

func NewGame(name, developer string, players []player.Player, tour int8, version string) *Game {
	return &Game{
		Name:        name,
		Developer:   developer,
		PlayersList: players,
		Tour:        tour,
		GameState:   PrepareGame,
		Version:     version,
	}
}

func (g *Game) WhoPlayed() *player.Player {
	if g.Tour == 1 {
		return &g.PlayersList[0]
	} else {
		return &g.PlayersList[1]
	}
}

func (g *Game) WhoEnemy() *player.Player {
	if g.Tour == 1 {
		return &g.PlayersList[1]
	} else {
		return &g.PlayersList[0]
	}
}

func (g *Game) GetPlayers() []player.Player {
	return g.PlayersList
}

func (g *Game) ChangeTour(newTour int8) {
	g.Tour = newTour
}

func (g *Game) ChangeGameState(newGameState GameState) {
	g.GameState = newGameState
}

func (g *Game) GetTour() int8 {
	return g.Tour
}

func (g *Game) GetGameState() GameState {
	return g.GameState
}

func (g *Game) IsPreparing() bool {
	return g.GameState == PrepareGame
}

func (g *Game) IsRunning() bool {
	return g.GameState == Running
}

func (g *Game) IsInRound() bool {
	return g.GameState == InRound
}

func (g *Game) IsFinished() bool {
	return g.GameState == EndGame
}