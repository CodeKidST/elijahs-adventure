package game

type Room struct {
	Name    string
	Item    int
	Monster string
	North   string
	South   string
	East    string
	West    string
}

func GetMap() map[string]Room {
	return map[string]Room{
		// The House
		"sunroom": {
			Name:  "sunroom",
			Item:  1,
			North: "exit",
			South: "living_room",
			East:  "wall",
			West:  "wall",
		},
		"living_room": {
			Name:  "living room",
			Item:  2,
			North: "sunroom",
			South: "kitchen",
			East:  "wall",
			West:  "wall",
		},
		"kitchen": {
			Name:  "kitchen",
			Item:  3,
			North: "living_room",
			South: "wall",
			East:  "bathroom",
			West:  "wall",
		},
		"bathroom": {
			Name:  "bathroom",
			Item:  4,
			North: "basement",
			South: "wall",
			East:  "wall",
			West:  "kitchen",
		}, // The Dungeon
		"basement": {
			Name:    "basement",
			Item:    5,
			Monster: "one-eyed alien",
			North:   "dark_room",
			South:   "bathroom",
			East:    "secret_room",
			West:    "wall",
		},
		"secret_room": {
			Name:  "secret_room",
			Item:  6,
			North: "wall",
			South: "wall",
			East:  "wall",
			West:  "basement",
		},
		"dark_room": {
			Name:    "dark_room",
			Item:    20,
			Monster: "slime",
			North:   "wall",
			South:   "basement",
			East:    "Library",
			West:    "wall",
		},

		"Library": {
			Name:    "Library",
			Item:    8,
			Monster: "cotton-candy monster",
			North:   "tunnel",
			South:   "wall",
			East:    "abandoned-mine",
			West:    "dark_room",
		},

		"abandoned-mine": {
			Name:  "abandoned-mine",
			Item:  9,
			North: "wall",
			South: "abandoned-lab",
			East:  "wall",
			West:  "Library",
		},

		"abandoned-lab": {
			Name:    "abandoned-lab",
			Item:    10,
			Monster: "Mutated scientist",
			North:   "abandoned-mine",
			South:   "pool-room",
			East:    "atomic room",
			West:    "wall",
		},

		"atomic room": {
			Name:    "atomic room",
			Item:    11,
			Monster: "Atomic Bomb",
			North:   "wall",
			South:   "wall",
			East:    "wall",
			West:    "abandoned-lab",
		},

		"tunnel": {
			Name:  "tunnel",
			Item:  11,
			North: "Room-stars",
			South: "Library",
			East:  "wall",
			West:  "wall",
		},

		"Room-stars": {
			Name:    "Room of stars",
			Item:    12,
			Monster: "Dragon of doom",
			North:   "wall",
			South:   "tunnel",
			East:    "wall",
			West:    "jail-cells",
		},

		"jail-cells": {
			Name:    "jail cells",
			Item:    13,
			Monster: "Evil broccoli",
			North:   "wall",
			South:   "evil-burgers",
			East:    "Room-stars",
			West:    "wall",
		},

		"evil-burgers": {
			Name:    "room of evil burgers",
			Item:    14,
			Monster: "Burger of death",
			North:   "jail-cells",
			South:   "wall",
			East:    "jail-cell",
			West:    "frog-room",
		},

		"frog-room": {
			Name:    "froggy room",
			Item:    15,
			Monster: "Froggy",
			North:   "challenge-room",
			South:   "wall",
			East:    "evil-burgers",
			West:    "skeleton-graveyard",
		},

		"skeleton-graveyard": {
			Name:    "skeleton graveyard",
			Item:    16,
			Monster: "Blue Devil",
			North:   "wall",
			South:   "wall",
			East:    "frog-room",
			West:    "wall",
		},

		"challenge-room": {
			Name:    "room of challenge",
			Item:    17,
			Monster: "witch",
			North:   "wall",
			South:   "frog-room",
			East:    "boss-room",
			West:    "stair-case",
		},

		"stair-case": {
			Name:    "stair case",
			Item:    18,
			Monster: "slime",
			North:   "wall",
			South:   "crafting-room",
			East:    "challenge-room",
			West:    "atl",
		},

		"atl": {
			Name:  "abandoned tech lab",
			Item:  21,
			North: "spy tower",
			South: "sorcerers-tower",
			East:  "challenge-room",
			West:  "wall",
		},

		"sorcerers-tower": {
			Name:    "sorcerers tower",
			Item:    22,
			Monster: "sorcerer",
			North:   "atl",
			South:   "dcave",
			East:    "trophy room",
			West:    "wall",
		},
	}

}
