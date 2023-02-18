package game

import (
	"fmt"
	"math/rand"
	"strings"
)

func (state *Game) Battle(monster *Player) bool {
	fmt.Printf("\nYou are battling the %s!\n", monster.Name)

	for {
		// Players turn
		input := state.Prompt("What do you do?")
		inputArray := strings.Split(input, " ")

		command := inputArray[0]
		switch command {
		case "die":
			state.Player.Life = 0
		case "cheat":
			monster.Life = 0
		case "attack":
			damage := state.Player.Attack - monster.Defense
			if damage < 0 {
				damage = 0
			}

			monster.Life -= damage
			fmt.Printf("%s dealt %d damage to the %s\n", state.Player.Name, damage, monster.Name)
		case "magic":
			if state.Player.Mana < 15 {
				fmt.Printf("You don't have enough mana! ")
				continue
			}

			damage := state.Player.Magic + monster.Defense
			fmt.Printf("you have poisoned the %s\n", monster.Name)
			fmt.Printf("///////////////////////////\n")
			fmt.Printf("-------CRITICAL HIT!-------\n")
			fmt.Printf("///////////////////////////\n")

			monster.Life -= damage
			fmt.Printf("%s dealt %d damage to the %s\n", state.Player.Name, damage, monster.Name)
			state.Player.Mana -= 15
		default:
			fmt.Printf("You did nothing!\n")
		}

		if monster.Life <= 0 {
			fmt.Printf("You defeated the %s\n", monster.Name)
			fmt.Printf("\tGot %d gold!\n", monster.Gold)
			fmt.Printf("\tGot %d Mana!\n", monster.Mana)
			state.Player.Gold += monster.Gold
			state.Player.Mana += monster.Mana
			for foodKey, foodAmount := range monster.Food {
				fmt.Printf("\tGot %d %s!\n", foodAmount, state.food[foodKey])
				state.Player.Food[foodKey] += foodAmount
			}
			return true
		}

		// Allies turn
		for _, ally := range state.Player.Allies {
			damage := ally.Attack - monster.Defense
			if damage < 0 {
				damage = 0
			}

			monster.Life -= damage
			fmt.Printf("%s dealt %d damage to the %s\n", ally.Name, damage, monster.Name)
		}

		var target *Player
		if len(state.Player.Allies) == 0 {
			target = state.Player
		} else {
			chance := rand.Intn(10)
			if chance > 4 {
				target = state.Player
			} else {
				chance = rand.Intn(len(state.Player.Allies))
				target = &state.Player.Allies[chance]
			}
		}

		// See if the monster misses
		chance := rand.Intn(10)
		if chance < 3 {
			fmt.Printf("The %s missed.\n", monster.Name)
			continue
		}

		// Monsters turn
		damage := monster.Attack - target.Defense
		if damage < 0 {
			damage = 0
		}

		state.Player.Life -= damage
		fmt.Printf("The %s dealt %d to %s\n", monster.Name, damage, target.Name)

		if state.Player.Life <= 0 {
			fmt.Printf("=================\n")
			fmt.Printf("===you've-died===\n")
			fmt.Printf("=================\n")
			fmt.Printf("....GAME OVER....\n")
			return false
		}
	}
}
