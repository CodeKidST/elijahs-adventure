package game

func GetMonsterList() map[string]Player {
	return map[string]Player{
		"slime": {
			Name:    "Slime",
			Life:    10,
			XP:      5,
			Gold:    1,
			Attack:  3,
			Defense: 1,
			Mana:    3,
			Food: map[string]int{
				"plum":  1,
				"apple": 1,
			},
		},

		"zombie": {
			Name:    "zombie",
			Life:    12,
			Gold:    3,
			Attack:  3,
			Defense: 1,
			Mana:    5,
			Food: map[string]int{
				"apple": 1,
			},
		},

		"witch": {
			Name:    "witch",
			Life:    13,
			Gold:    5,
			Attack:  4,
			Defense: 2,
			Mana:    3,
			Food: map[string]int{
				"banana": 1,
			},
		},

		"one-eyed alien": {
			Name:    "one-eyed alien",
			Life:    11,
			Gold:    7,
			Attack:  3,
			Defense: 2,
			Mana:    3,
			Food: map[string]int{
				"alien-guts": 2,
			},
		},

		"cotton-candy monster": {
			Name:    "cotton-candy monster",
			Life:    4,
			Gold:    2,
			Attack:  3,
			Defense: 1,
			Mana:    2,
			Food: map[string]int{
				"cotton-candy": 1,
			},
		},

		"Mutated scientist": {
			Name:    "BOSS\nMutated scientist",
			Life:    25,
			Gold:    16,
			Attack:  7,
			Defense: 2,
			Mana:    16,
			Food: map[string]int{
				"lemon": 1,
			},
		},

		"Atomic Bomb": {
			Name:    "Atomic Bomb",
			Life:    1000,
			Gold:    100,
			Attack:  500,
			Defense: 0,
		},

		"Dragon of doom": {
			Name:    "Dragon of doom",
			Life:    13,
			Gold:    11,
			Attack:  5,
			Defense: 2,
			Mana:    9,
			Food: map[string]int{
				"sweet-fruit": 1,
			},
		},

		"Evil broccoli": {
			Name:    "BOSS\nEvil broccoli",
			Life:    21,
			Gold:    15,
			Attack:  6,
			Defense: 2,
			Mana:    4,
			Food: map[string]int{
				"salad": 2,
			},
		},

		"Burger of death": {
			Name:    "Burger of death",
			Life:    13,
			Gold:    4,
			Attack:  4,
			Defense: 1,
			Mana:    6,
			Food: map[string]int{
				"rainbow-hamburger": 1,
			},
		},

		"Froggy": {
			Name:    "Froggy",
			Life:    17,
			Gold:    6,
			Attack:  5,
			Defense: 2,
			Mana:    2,
			Food: map[string]int{
				"froggy-meat": 1,
				"frog-legs":   1,
			},
		},

		"Blue Devil": {
			Name:    "Blue Devil",
			Life:    14,
			Gold:    8,
			Attack:  3,
			Defense: 1,
			Mana:    5,
			Food: map[string]int{
				"hot-dog": 1,
			},
		},

		"clown": {
			Name:    "clown",
			Life:    19,
			Gold:    5,
			Attack:  5,
			Defense: 1,
			Mana:    2,
			Food: map[string]int{
				"magic-corn": 2,
			},
		},

		"vicious dog": {
			Name:    "vicious dog",
			Life:    19,
			Gold:    3,
			Attack:  6,
			Defense: 2,
			Mana:    7,
			Food: map[string]int{
				"magic-corn": 2,
			},
		},

		"sorcerer": {
			Name:    "sorcerer",
			Life:    21,
			Gold:    7,
			Attack:  8,
			Defense: 3,
			Mana:    20,
			Food: map[string]int{
				"protien-orb": 1,
			},
		},

		// DO NOT ADD MONSTERS AFTER THIS LINE
	}
}
