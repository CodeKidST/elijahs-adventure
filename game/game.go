package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	items       []Item
	monsters    map[string]Player
	food        map[string]string
	house       map[string]Room
	Player      *Player
	CurrentRoom Room
	RoomName    string
	Exit        bool
}

func NewGame(player *Player) *Game {

	f := GetFoodList()

	for foodKey := range f {
		player.Food[foodKey] = 0
	}

	return &Game{
		reader:   bufio.NewReader(os.Stdin),
		items:    GetItemList(),
		monsters: GetMonsterList(),
		food:     f,
		house:    GetMap(),
		Player:   player,
		Exit:     false,
	}
}

func (state *Game) Start() {
	TitleScreen()
	fmt.Printf("play game\n")
	fmt.Printf("read instructions\n")
	input := state.Prompt("Do you want to read the intructions or play the game?")
	if input == "read instructions" {
		InstructionScreen()
	}

	for {
		if state.Exit {
			break
		}

		state.Update()
	}
	fmt.Printf("You have left the dungeon!\n")
}

func (state *Game) SetRoom(name string) {
	if r, ok := state.house[name]; ok {
		state.CurrentRoom = r
		state.RoomName = name
	}
}

func (state *Game) Prompt(p string) string {
	fmt.Printf("%s ", p)
	input, err := state.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]
	return strings.ToLower(input)
}
