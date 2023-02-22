package game

import "fmt"

func Check(param string, state *Game) {

	switch param {
	case "inventory":
		checkInventory(state)
	case "food":
		checkFood(state)
	case "life":
		checkLife(state)
	case "gold":
		checkGold(state)
	case "allies":
		checkAllies(state)
	case "mana":
		checkMana(state)
	case "hydration":
		checkHydration(state)
	case "status":
		checkStatus(state)
    
	default:
		fmt.Printf("That's not a valid command!\n")
	}
}

func checkInventory(state *Game) {
	if len(state.Player.Inventory) == 0 {
		fmt.Printf("your Inventory is empty.\n")
	} else {
		fmt.Printf("You are carrying.\n")
		for index, item := range state.Player.Inventory {
			fmt.Printf("\t%d: %s\n", index+1, item.Name)
		}
	}
}

func checkFood(state *Game) {
	fmt.Printf("You are carrying:\n")
	for foodKey, foodAmount := range state.Player.Food {
		fmt.Printf("\t%s X %d\n", state.food[foodKey], foodAmount)
	}
}

func checkLife(state *Game) {
	fmt.Printf("%d hp\n", state.Player.Life)
}

func checkGold(state *Game) {
	fmt.Printf("You have %d gold!\n", state.Player.Gold)
}

func checkAllies(state *Game) {
	for _, ally := range state.Player.Allies {
		fmt.Printf("%s has %d life\n", ally.Name, ally.Life)

	}
}

func checkMana(state *Game) {
	fmt.Printf("You have %d mana!\n", state.Player.Mana)
}

func checkHydration(state *Game) {
	fmt.Printf("you are %d hydrated", state.Player.Hydration)
	if state.Player.Hydration <= 0 {
		state.Player.Life = 0
		fmt.Printf("GAME OVER\n You have died from thirst")
	}
}

func checkStatus(state *Game) {
	fmt.Printf("Player Status\n")
	fmt.Printf("-------------\n")
	fmt.Printf("\tHP: %d\n", state.Player.Life)
	fmt.Printf("\tAtk: %d\n", state.Player.Attack)
	fmt.Printf("\tDef: %d\n", state.Player.Defense)
	fmt.Printf("\tMana: %d\n", state.Player.Mana)
	fmt.Printf("\nHydration: %d\n", state.Player.Hydration)

}
