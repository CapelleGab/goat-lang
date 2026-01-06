package core

type Player struct {
	ID    int
	Name  string
	Email string
	Stats PlayerStats
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

type Game struct {
	Name      string
	Developer string
	Players   []Player
	Tour      int
	Version   string
}

func getPlayer(id int) Player {
	for _, player := range getPlayers() {
		if player.ID == id {
			return player
		}
	}
	return Player{}
}

func getPlayerStats(player Player) PlayerStats {
	return player.Stats
}

func getSpecificStatsForPlayer(player Player, statName string) int {
	stats := getPlayerStats(player)
	switch statName {
	case "mana":
		return stats.Mana
	case "max_mana":
		return stats.MaxMana
	case "health":
		return stats.Health
	case "max_health":
		return stats.MaxHealth
	case "damage":
		return stats.Damage
	case "defense":
		return stats.Defense
	case "intelligence":
		return stats.Intelligence
	default:
		return 0
	}
}

func setSpecificStatsForPlayer(player Player, statName string, value int) {
	stats := getPlayerStats(player)
	switch statName {
	case "mana":
		stats.Mana = value
	case "max_mana":
		stats.MaxMana = value
	case "health":
		stats.Health = value
	case "max_health":
		stats.MaxHealth = value
	case "damage":
		stats.Damage = value
	case "defense":
		stats.Defense = value
	case "intelligence":
		stats.Intelligence = value
	default:
		return
	}
	player.Stats = stats
}