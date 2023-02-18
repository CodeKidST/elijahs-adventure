package game

import "fmt"

func Look(param string, state *Game) {
	fmt.Printf("You look around and see a %s\n", state.items[state.CurrentRoom.Item].Name)
}
