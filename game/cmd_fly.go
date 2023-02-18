package game

import "fmt"

func Fly(param string, state *Game) {
	nextRoom, ok := state.house[param]
	if !ok {
		fmt.Printf("an error has occurred!\n")
	} else {
		state.RoomName = param
		state.CurrentRoom = nextRoom
	}
}
