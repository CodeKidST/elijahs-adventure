package game

import (
	"fmt"
	"strings"
)

func Get(param string, state *Game) {
	if state.CurrentRoom.Item == 0 {
		fmt.Printf("The room is empty\n")
	} else {

		item := state.items[state.CurrentRoom.Item]

		switch strings.ToLower(item.Name) {
		case "crafting bench",
			"couch",
			"chair",
			"fridge",
			"toilet":

			fmt.Printf("you can't pick that up\n")
			return
		}

		fmt.Printf("You picked up the %s\n", item.Name)
		state.CurrentRoom.Item = 0
		state.house[state.RoomName] = state.CurrentRoom

		if item.Name == "little froggy" {
			fmt.Printf("little froggy: You saved me!\n")
			fmt.Printf("little froggy: I got lost on my way to the pond can you help me get back!?\n")
			fmt.Printf("little froggy is now you're ally!\n")
			frog := Player{
				Name:    item.Name,
				Life:    50,
				Attack:  item.Attack,
				Defense: 1,
			}
			state.Player.Allies = append(state.Player.Allies, frog)
		} else {
			state.Player.Inventory = append(state.Player.Inventory, item)
			if item.Attack > 0 {
				state.Player.Attack += item.Attack
				fmt.Printf("%s attack went up by %d\n", state.Player.Name, item.Attack)
			}
		}

	}

	/*}  else if command == "" {
	 */
}
