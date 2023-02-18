package game

import "fmt"

func Fight(param string, state *Game) {
	monster, ok := state.monsters[param]
	if !ok {
		fmt.Printf("That monster doesn't exist\n")
	} else {
		success := state.Battle(&monster)
		if !success {
			fmt.Printf("Dungeon failed")
			return
		}
	}
}
