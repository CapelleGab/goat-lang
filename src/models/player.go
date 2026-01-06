package models

type Player struct {
	ID     int
	Name   string
	Email  string
	Stats  PlayerStats
	Skills []Skill
}

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
