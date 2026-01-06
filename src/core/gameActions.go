package core

import "fmt"

func attack(attacker Player, defender Player) {
	// Implement attack logic here
	if !canAttack(attacker) {
		fmt.Printf("Tu n'as plus de mana")
		return;
	}
	
	
}

func canAttack(attacker Player) bool {
	if attacker.Stats.Mana <= 0 {
		return false
	}
	return true
}

func useMana(player Player, amount int) {
	// Implement mana usage logic here
	mana := getSpecificStatsForPlayer(player, "mana")
	
	newManaValue := mana - amount
	if newManaValue < 0 {
		fmt.Printf("Tu n'as plus de mana")
		return;
	}

	setSpecificStatsForPlayer(player, "mana", newManaValue)
}

func regenMana(player Player, amount int) {
	// Implement mana regeneration logic here (multithreading with goroutines (go func))
}

func useSkill(attacker Player, enemy Player) {
	if attacker.Stats.Mana < 10 {
		fmt.Println("Tu n'as pas assez de mana pour utiliser une compétence")
		return
	}
	
	skill := getSpecificStatsForPlayer(attacker, "skill")
	
	newManaValue := attacker.Stats.Mana - 10
	setSpecificStatsForPlayer(attacker, "mana", newManaValue)
	
	newHealthValue := enemy.Stats.Health - skill
	setSpecificStatsForPlayer(enemy, "health", newHealthValue)
}

func heal(attacker Player) {
	if attacker.Stats.Mana < 10 {
		fmt.Println("Tu n'as pas assez de mana pour te soigner")
		return
	}
	
	heal := getSpecificStatsForPlayer(attacker, "heal")
	
	newManaValue := attacker.Stats.Mana - 10
	setSpecificStatsForPlayer(attacker, "mana", newManaValue)
	
	newHealthValue := attacker.Stats.Health + heal
	setSpecificStatsForPlayer(attacker, "health", newHealthValue)
}


