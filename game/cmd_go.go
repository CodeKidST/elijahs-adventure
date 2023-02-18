package game

import "fmt"

func Go(param string, state *Game) {
	if param == "north" {
		state.RoomName = state.CurrentRoom.North

	} else if param == "south" {
		state.RoomName = state.CurrentRoom.South

	} else if param == "east" {
		state.RoomName = state.CurrentRoom.East

	} else if param == "west" {
		state.RoomName = state.CurrentRoom.West
	} else {
		fmt.Printf("You can only go north, south, east or west...\n")
	}

	if state.RoomName == "exit" {
		state.Exit = true
	} else if state.RoomName == "wall" {
		fmt.Printf("There is a wall in the way!\n")
	} else {
		nextRoom, ok := state.house[state.RoomName]
		if !ok {
			fmt.Printf("Oh no an error has occurred!\n")
		} else {
			state.CurrentRoom = nextRoom
		}
	}
}
