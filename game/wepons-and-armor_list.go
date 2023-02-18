package game

type Item struct {
	Name   string
	Attack int
}

func GetItemList() []Item {
	return []Item{
		{Name: "Nothing"},
		{Name: "Chair"},
		{
			Name: "couch",
		},
		{
			Name: "fridge",
		},
		{
			Name: "toilet",
		},
		{
			Name: "spider webs",
		},
		{
			Name: "chest full of gold",
		},
		{
			Name: "piece of wood",
		},
		{
			Name:   "Rusty sword",
			Attack: 2,
		},
		{
			Name: "Iron ores",
		},
		{
			Name: "bottle",
		},
		{
			Name: "Rare Gem",
		},
		{
			Name: "bottle of stardust",
		},
		{
			Name: "Cell key",
		},
		{
			Name:   "Rusty axe",
			Attack: 3,
		},
		{
			Name:   "little froggy",
			Attack: 3,
		},
		{
			Name: "pile of bones",
		},
		{
			Name:   "Iron sword",
			Attack: 5,
		},
		{
			Name: "broken pipe",
		},
		{
			Name:   "star sword",
			Attack: 9,
		},
		{Name: "crafting bench"},
		{Name: "baseball bat",
			Attack: 3,
		},
		{Name: "wand",
	    Attack: 3,},
	}
}

func GetItem(itemList []Item, name string) (Item, bool) {
	for _, i := range itemList {
		if i.Name == name {
			return i, true
		}
	}

	return Item{}, false
}
