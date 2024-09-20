package logparser

type Game struct {
	TotalKills int
	Players    map[string]Player
}

func (g *Game) PlayerNames() []string { //retorna lista de players
	names := make([]string, 0, len(g.Players))
	for _, player := range g.Players {
		names = append(names, player.Name, " - ")
	}
	return names
}

func (g *Game) Kills() map[string]int { //retorna um map de players e kills
	kills := make(map[string]int)
	for _, player := range g.Players {
		kills[player.Name] = player.Kills
	}
	return kills
}

type Player struct {
	Name  string
	Kills int
}

type KillData struct {
	Killer string
	Victim string
	Weapon string
}
