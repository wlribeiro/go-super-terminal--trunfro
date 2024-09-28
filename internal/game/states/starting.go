package states

import (
	"fmt"
	"super-trunfo/internal/cards"
	"super-trunfo/internal/game"
	"super-trunfo/internal/player"
	"super-trunfo/utils"
)

type StartingState struct{}

func (s *StartingState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("Iniciando o jogo...")
	utils.AwaitRead()

	deck := cards.Gods

	ctx.Player = player.Player{
		Name:  "Player",
		Deck:  []cards.Card{},
		Score: 0,
	}

	ctx.Computer = player.Player{
		Name:  "Computer",
		Deck:  []cards.Card{},
		Score: 0,
	}

	ctx.Cards = deck

	sm.SetState(&ShuffleState{})
}
