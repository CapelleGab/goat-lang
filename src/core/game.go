package core

import (
	"fmt"
)

func getPlayers() []Player {
	return []Player{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			Stats: PlayerStats{
				Mana:         50,
				MaxMana:      100,
				Health:       100,
				MaxHealth:    100,
				Damage:       10,
				Defense:      5,
				Intelligence: 8,
			},
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane@example.com",
			Stats: PlayerStats{
				Mana:         50,
				MaxMana:      100,
				Health:       100,
				MaxHealth:    100,
				Damage:       10,
				Defense:      5,
				Intelligence: 8,
			},
		},
	}
}

func PrepareGame() {
	game := Game{
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

func runGame(game Game) {
	fmt.Println("Game is running!")

	fmt.Println(getSpecificStatsForPlayer(getPlayer(1), "health"))
	var isRunning bool = true
	for isRunning {

		startRound(game.Tour)

		if game.Tour == 1 {
			fmt.Println("Fin du tour de", getPlayer(1).Name)
			fmt.Println("Début du tour de", getPlayer(2).Name)
			game.Tour = 2
		} else if game.Tour == 2 {
			fmt.Println("Fin du tour de", getPlayer(2).Name)
			fmt.Println("Début du tour de", getPlayer(1).Name)
			game.Tour = 1
		}
	}
}

func whoPlayed(tour int) Player {
	if tour == 1 {
		return getPlayer(0)
	} else {
		return getPlayer(1)
	}
}

func whoEnemy(tour int) Player {
	if tour == 1 {
		return getPlayer(1)
	} else {
		return getPlayer(0)
	}
}

func startRound(tour int) {
	fmt.Println("Starting round...")
	attacker := whoPlayed(tour)
	enemy := whoEnemy(tour)

	fmt.Printf("Attacker: %s , ", attacker.Name)
	fmt.Printf("Enemy: %s\n", enemy.Name)
	
	fmt.Println("Que veux tu faire ?")
	
	if attacker.Stats.Mana >= 10 {
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser une compétence")
		fmt.Println("3. Se soigner")
		fmt.Println("4. Fuir")
	} else {
		fmt.Println("1. Attaquer")
		fmt.Println("2. Se soigner")
		fmt.Println("3. Fuir")
	}
	
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		attack(attacker, enemy)
	case 2:
		useSkill(attacker, enemy)
	case 3:
		heal(attacker)
	case 4:
		flee(attacker)
	default:
		fmt.Println("Choix invalide")
	}

}
