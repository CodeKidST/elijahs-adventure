package game

import (
	"fmt"
	"strings"
)

func (state *Game) Update() {

	fmt.Printf("\nYou are in the %s\n", state.CurrentRoom.Name)

	monster, ok := state.monsters[state.CurrentRoom.Monster]
	if ok {
		fmt.Printf("There is a %s\n", monster.Name)
		success := state.Battle(&monster)
		if !success {
			fmt.Printf("Game Over\n")
			return
		} else {
			state.CurrentRoom.Monster = ""
			state.house[state.RoomName] = state.CurrentRoom
			fmt.Printf("You were victorious!\n")
		}
	}

	// Handle user input
	input := state.Prompt("What do you do?")
	inputArray := strings.Split(input, " ")

	if len(inputArray) < 2 {
		if input == "exit" {
			state.Exit = true
			fmt.Printf("What a quitter!\n")
			return
		}

		fmt.Printf("That is a invalid command!\n")
	} else {
		// commands
		command := inputArray[0]
		param := inputArray[1]

		switch command {
		case "go":
			Go(param, state)
		case "look":
			Look(param, state)
		case "get":
			Get(param, state)
		case "check":
			Check(param, state)
		case "use":
			Use(param, state)
		case "fly":
			Fly(param, state)
		case "fight":
			Fight(param, state)
		case "eat":
			Eat(param, state)
			
		default:
			fmt.Printf("'%s' is an invalid command\n", input)
		}
	}
}
