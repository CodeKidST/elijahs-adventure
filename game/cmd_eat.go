package game

import "fmt"

func Eat(param string, state *Game) {
	foodAmount, ok := state.Player.Food[param]
	if !ok || foodAmount == 0 {
		fmt.Printf("I cannot eat that\n")
	} else {
		fmt.Printf("You ate %s for %d HP\n", state.food[param], 5)
		state.Player.Life += 5
		if state.Player.Life > 100 {
			state.Player.Life = 100
		}
		state.Player.Food[param]--
	}
}
