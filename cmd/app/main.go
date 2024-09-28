package main

import (
	"fmt"
	"super-trunfo/internal/cards"
	"super-trunfo/internal/game"
	"super-trunfo/internal/game/states"
	"super-trunfo/internal/player"
)

func menu() {
	fmt.Println("1 - Jogar")
	fmt.Println("2 - Sair")
}

func main() {
	sm := game.NewStateMachine()

	deck := cards.Gods

	ctx := &game.Context{
		Player:   player.Player{},
		Computer: player.Player{},
		Cards:    deck,
	}

	sm.SetState(&states.StartingState{})
	sm.Process(ctx)
}
