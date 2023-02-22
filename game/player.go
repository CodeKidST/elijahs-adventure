package game

import "math"

type Player struct {
	Name      string
	Life      int
	Mana      int
	Gold      int
	Magic     int
	Attack    int
	Defense   int
	Hydration int
	Level     int
	XP        int
	Inventory []Item
	Food      map[string]int
	Allies    []Player
}

func (p *Player) CountInventory() map[string]int {
	m := make(map[string]int)

	for _, i := range p.Inventory {
		m[i.Name]++
	}

	return m
}

func (p *Player) AddXP(xp int) bool {
	p.XP += xp
	nl := int(math.Floor(.2 * math.Sqrt(float64(p.XP))))
	if p.Level < nl {
		p.Level = nl
		return true
	}
	return false
}
