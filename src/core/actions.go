package core

import (
	"fmt"
	"goat-lang/src/models"
)

func attack(attacker *models.Player, defender *models.Player) {
	if attacker.Stats.Mana <= 0 {
		fmt.Println("Tu n'as plus de mana pour attaquer!")
		return
	}

	damage := attacker.Stats.Damage - defender.Stats.Defense
	if damage < 0 {
		damage = 0
	}

	defender.Stats.Health -= damage
	attacker.Stats.Mana -= 5

	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", attacker.Name, defender.Name, damage)
	fmt.Printf("%s a maintenant %d HP\n", defender.Name, defender.Stats.Health)
}

func useSkill(attacker *models.Player, enemy *models.Player) {
	if len(attacker.Skills) == 0 {
		fmt.Println("Tu n'as aucune compétence disponible!")
		return
	}

	fmt.Println("\nCompétences disponibles:")
	for i, skill := range attacker.Skills {
		fmt.Printf("%d. %s (Coût: %d mana", i+1, skill.Name, skill.ManaCost)
		if skill.Damage > 0 {
			fmt.Printf(", Dégâts: %d", skill.Damage)
		}
		if skill.Healing > 0 {
			fmt.Printf(", Soin: %d", skill.Healing)
		}
		fmt.Println(")")
	}
	fmt.Println("0. Annuler")

	var choice int
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(attacker.Skills) {
		fmt.Println("Choix invalide!")
		return
	}

	selectedSkill := attacker.Skills[choice-1]

	if attacker.Stats.Mana < selectedSkill.ManaCost {
		fmt.Printf("Tu n'as pas assez de mana! (Requis: %d, Actuel: %d)\n", selectedSkill.ManaCost, attacker.Stats.Mana)
		return
	}

	attacker.Stats.Mana -= selectedSkill.ManaCost

	if selectedSkill.Type == "offensive" {
		damage := selectedSkill.Damage + (attacker.Stats.Intelligence / 2)
		enemy.Stats.Health -= damage
		fmt.Printf("%s utilise %s et inflige %d dégâts à %s!\n", attacker.Name, selectedSkill.Name, damage, enemy.Name)
		fmt.Printf("%s a maintenant %d HP\n", enemy.Name, enemy.Stats.Health)
	} else if selectedSkill.Type == "healing" {
		healing := selectedSkill.Healing
		attacker.Stats.Health += healing
		if attacker.Stats.Health > attacker.Stats.MaxHealth {
			attacker.Stats.Health = attacker.Stats.MaxHealth
		}
		fmt.Printf("%s utilise %s et récupère %d HP!\n", attacker.Name, selectedSkill.Name, healing)
		fmt.Printf("%s a maintenant %d HP\n", attacker.Name, attacker.Stats.Health)
	}
}

func flee(attacker *models.Player) bool {
	fmt.Printf("%s tente de fuir le combat...\n", attacker.Name)
	return false
}


