package player

import "github.com/google/uuid"

type PlayerStats struct {
	Mana         int
	MaxMana      int
	Health       int
	MaxHealth    int
	Damage       int
	Defense      int
	Intelligence int
}

type Skill struct {
	Name     string
	ManaCost int
	Damage   int
	Healing  int
	Type     string
}

type Player struct {
	ID          uuid.UUID
	Name        string
	Stats       PlayerStats
	Skills      []Skill
}