package core

import (
	"fmt"
	"goat-lang/src/models"
)

func getPlayers() []models.Player {
	return []models.Player{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			Stats: models.PlayerStats{
				Mana:         50,
				MaxMana:      100,
				Health:       100,
				MaxHealth:    100,
				Damage:       15,
				Defense:      5,
				Intelligence: 8,
			},
			Skills: []models.Skill{
				{Name: "Boule de feu", ManaCost: 15, Damage: 20, Healing: 0, Type: "offensive"},
				{Name: "Éclair", ManaCost: 20, Damage: 30, Healing: 0, Type: "offensive"},
				{Name: "Soin mineur", ManaCost: 10, Damage: 0, Healing: 25, Type: "healing"},
			},
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane@example.com",
			Stats: models.PlayerStats{
				Mana:         50,
				MaxMana:      100,
				Health:       100,
				MaxHealth:    100,
				Damage:       15,
				Defense:      5,
				Intelligence: 8,
			},
			Skills: []models.Skill{
				{Name: "Lame de glace", ManaCost: 15, Damage: 22, Healing: 0, Type: "offensive"},
				{Name: "Tempête", ManaCost: 25, Damage: 35, Healing: 0, Type: "offensive"},
				{Name: "Régénération", ManaCost: 12, Damage: 0, Healing: 30, Type: "healing"},
			},
		},
	}
}

func PrepareGame() {
	game := models.Game{
		Name:      "Goat Lang",
		Developer: "CapelleGab",
		Players:   getPlayers(),
		Version:   "1.0.0",
		Tour:      1,
	}

	fmt.Printf("Launching game %s...\n", game.Name)
	for _, player := range game.Players {
		fmt.Printf("Player %d: %s \n", player.ID, player.Name)
	}

	fmt.Println("Game is ready to start!")
	runGame(game)
}

func runGame(game models.Game) {
	fmt.Println("Game is running!")

	for {
		player1 := &game.Players[0]
		player2 := &game.Players[1]

		if player1.Stats.Health <= 0 {
			fmt.Printf("\n%s a été vaincu! %s gagne!\n", player1.Name, player2.Name)
			break
		}
		if player2.Stats.Health <= 0 {
			fmt.Printf("\n%s a été vaincu! %s gagne!\n", player2.Name, player1.Name)
			break
		}

		currentPlayer := whoPlayed(game.Tour, &game)
		if !startRound(&game) {
			fmt.Printf("\n%s a fui le combat!\n", currentPlayer.Name)
			break
		}

		if game.Tour == 1 {
			fmt.Printf("\nFin du tour de %s\n", game.Players[0].Name)
			fmt.Printf("Début du tour de %s\n", game.Players[1].Name)
			game.Tour = 2
		} else {
			fmt.Printf("\nFin du tour de %s\n", game.Players[1].Name)
			fmt.Printf("Début du tour de %s\n", game.Players[0].Name)
			game.Tour = 1
		}
	}
}

func whoPlayed(tour int, game *models.Game) *models.Player {
	if tour == 1 {
		return &game.Players[0]
	} else {
		return &game.Players[1]
	}
}

func whoEnemy(tour int, game *models.Game) *models.Player {
	if tour == 1 {
		return &game.Players[1]
	} else {
		return &game.Players[0]
	}
}

func startRound(game *models.Game) bool {
	fmt.Println("\n--- Nouveau tour ---")
	attacker := whoPlayed(game.Tour, game)
	enemy := whoEnemy(game.Tour, game)

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
		attack(attacker, enemy)
	case 2:
		useSkill(attacker, enemy)
	case 3:
		return flee(attacker)
	default:
		fmt.Println("Choix invalide")
	}

	return true
}
