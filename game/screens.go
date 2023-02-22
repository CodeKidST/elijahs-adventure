package game

import "fmt"

func TitleScreen() {
	fmt.Printf("=======THE=========\n")
	fmt.Printf("===Adventures of===\n")
	fmt.Printf("======FROGGY=======\n")
	fmt.Printf("===================\n")

}

func InstructionScreen() {
	fmt.Printf("------------------------------------------------------------------------------------\n")
	fmt.Printf("Welcome to The Adventures Of Froggy! Here are some instructions to help you understand the game!\n")
	fmt.Printf("When the screen shows the name of the room and asks, 'what do you do'\nyou can type go north, south, east, or west\n")
	fmt.Printf("To check your hydration type 'check hydration' and to check your mana type 'check mana' you need 15 mana to use a magic attack\n")
	fmt.Printf("and to check your level and xp type 'check level'\n")
	fmt.Printf("To check your HP type 'check life' and it should display your life. If it dosn't reset the game!\n")
	fmt.Printf("To check inventory type 'check inventory', if you would like to see what items are in the room\n type look around\n")
	fmt.Print("To pick up items type 'get (item name)'. To check your food type 'check food',\n and to eat type 'eat (item name)'\n")
	fmt.Printf("To use the crafting benches type 'use bench' and it will open a crafting menu\n")
	fmt.Printf("when you encounter a monster to attack type 'attack' and it will display the damage you have done to the monster\n")
	fmt.Printf("To use a magic attack type 'magic' while battling and it will do a special attack\n")
	fmt.Printf("you will have a defense against the monster lowering it's attack.\n the monsters attack will either hit you or miss\n")
	fmt.Printf("later in the game you meet an ally the object of this game is to return the ally to it's home!\n")
	fmt.Printf("once you do that you have to make it out of the dungeon with all upgrades to win the game!\n ENJOY!")
	fmt.Printf(" created by Elijah.A 2023\n")
	fmt.Printf("=======THE=======\n")
	fmt.Printf("==Adventures of==\n")
	fmt.Printf("======Froggy=====\n") 
}
