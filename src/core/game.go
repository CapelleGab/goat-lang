package core

import (
	"fmt"

	game "goat-lang/src/models/Game"
	player "goat-lang/src/models/Player"

	"github.com/google/uuid"
)

func PrepareGame() {
	p1 := player.NewPlayer(
		uuid.New(),
		"John Doe",
		player.PlayerStats{
			Mana:         50,
			MaxMana:      50,
			Health:       100,
			MaxHealth:    100,
			Damage:       15,
			Defense:      5,
			Intelligence: 100,
		},
		[]player.Skill{
			{
				Name:     "Eclair",
				ManaCost: 20,
				Damage:   10,
				Healing:  0,
				Type:     "Electric",
			},
		},
	)
	p2 := player.NewPlayer(
		uuid.New(),
		"Jane Doe",
		player.PlayerStats{
			Mana:         50,
			MaxMana:      50,
			Health:       100,
			MaxHealth:    100,
			Damage:       15,
			Defense:      5,
			Intelligence: 100,
		},
		[]player.Skill{
			{
				Name:     "Eclair",
				ManaCost: 20,
				Damage:   10,
				Healing:  0,
				Type:     "Electric",
			},
		},
	)
	gameInstance := game.NewGame("Goat Lang", "CapelleGab", []player.Player{*p1, *p2}, 1, "1.0.0")
	gameInstance.ChangeGameState(game.PrepareGame)

	fmt.Printf("Launching game %s...\n", gameInstance.Name)
	for _, player := range gameInstance.PlayersList {
		fmt.Printf("Player %s : %s \n", player.ID.String(), player.Name)
	}

	fmt.Println("Game is ready to start!")
	runGame(gameInstance)
}

func runGame(gameInstance *game.Game) {
	fmt.Println("Game is running!")
	// Change game state to running

	for {
		player1 := &gameInstance.GetPlayers()[0]
		player2 := &gameInstance.GetPlayers()[1]

		if player1.Stats.Health <= 0 {
			fmt.Printf("\n%s a été vaincu! %s gagne!\n", player1.Name, player2.Name)
			break
		}
		if player2.Stats.Health <= 0 {
			fmt.Printf("\n%s a été vaincu! %s gagne!\n", player2.Name, player1.Name)
			break
		}

		currentPlayer := gameInstance.WhoPlayed()
		if !startRound(gameInstance) {
			fmt.Printf("\n%s a fui le combat!\n", currentPlayer.Name)
			break
		}

		if gameInstance.GetTour() == 1 {
			fmt.Printf("\nFin du tour de %s\n", gameInstance.PlayersList[0].Name)
			fmt.Printf("Début du tour de %s\n", gameInstance.PlayersList[1].Name)
			gameInstance.ChangeTour(2)
		} else {
			fmt.Printf("\nFin du tour de %s\n", gameInstance.PlayersList[1].Name)
			fmt.Printf("Début du tour de %s\n", gameInstance.PlayersList[0].Name)
			gameInstance.ChangeTour(1)
		}
	}
}

func startRound(gameInstance *game.Game) bool {
	fmt.Println("\n--- Nouveau tour ---")
	attacker := gameInstance.WhoPlayed()
	enemy := gameInstance.WhoEnemy()

	fmt.Printf("Attaquant: %s (HP: %d/%d, Mana: %d/%d)\n",
		attacker.Name, attacker.Stats.Health, attacker.Stats.MaxHealth,
		attacker.Stats.Mana, attacker.Stats.MaxMana)
	fmt.Printf("Ennemi: %s (HP: %d/%d)\n",
		enemy.Name, enemy.Stats.Health, enemy.Stats.MaxHealth)

	fmt.Println("\nQue veux-tu faire ?")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser une compétence")
	fmt.Println("3. Fuir")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		player.Attack(attacker, enemy)
	case 2:
		player.UseSkill(attacker, enemy)
	case 3:
		return player.Flee(attacker)
	default:
		fmt.Println("Choix invalide")
	}

	return true
}
