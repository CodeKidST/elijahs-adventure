package main

import (
	"eli1/game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	player := game.Player{
		Name:    "Elijah",
		Life:    100,
		Mana:    5,
		Attack:  5,
		Magic:   5,
		Defense: 2,
		Hydration: 100,
		Food:    make(map[string]int),
	}

	g := game.NewGame(&player)
	g.SetRoom("sunroom")
	g.Start()

}
