package main

import (
	"fmt"
	"github.com/Amovement/auto-magical-animal/internal/game"
)

func main() {
	fmt.Println("Auto Magical Animal.")
	fmt.Println("Github: https://github.com/Amovement/auto-magical-animal")
	game.SetGameWindow("Auto Magical Animal")
	game.StartGame()
}
