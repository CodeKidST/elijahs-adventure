package game

type Player struct {
	Name      string
	Life      int
	Mana        int
	Gold      int
	Magic     int
	Attack    int
	Defense   int
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
