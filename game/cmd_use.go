package game

import "fmt"

func Use(param string, state *Game) {
	if param == "bench" {
		fmt.Printf("you are now using the crafting bench\n")
		fmt.Printf("you can craft\n")
		fmt.Printf("star sword: items needed\n x1 bottle of stardust:\n x1 Iron ores: ")

		itemCounts := state.Player.CountInventory()

		if itemCounts["bottle of stardust"] >= 1 && itemCounts["Iron ores"] >= 1 {
			if i, ok := GetItem(state.items, "star sword"); ok {
				state.Player.Inventory = append(state.Player.Inventory, i)
				state.Player.Attack += i.Attack
			}

		} else {
			fmt.Printf("You don't have enough materials!\n")
		}
	}
}
