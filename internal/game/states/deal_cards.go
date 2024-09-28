package states

import (
	"fmt"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
)

type DealCardsState struct{}

func (s *DealCardsState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("Distribuindo cartas...")
	utils.AwaitRead()
	utils.ClearScreen()

	size := len(ctx.Cards)

	ctx.Player.Deck = ctx.Cards[:size]
	ctx.Computer.Deck = ctx.Cards[size/2:]

	sm.SetState(&PlayerTurnState{})
}
