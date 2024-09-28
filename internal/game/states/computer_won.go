package states

import (
	"fmt"
	"super-trunfo/internal/game"
)

type ComputerWonState struct{}

func (s *ComputerWonState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("O computador ganhou!")
	ctx.Computer.Score += 1

	card, fail := ctx.Player.DrawCard()
	if fail != nil {
		sm.SetState(&GameOverState{})
		return
	}

	ctx.Player.ResetChoice()

	ctx.Computer.AddCard(card)
	card, fail = ctx.Computer.DrawCard()
	if fail != nil {
		sm.SetState(&GameOverState{})
		return
	}
	ctx.Computer.AddCard(ctx.Player.ChosenCard)
	ctx.Computer.ResetChoice()

	if !ctx.Player.HaveCards() {
		sm.SetState(&GameOverState{})
		return
	}
	sm.SetState(&ComputerTurnState{})
}
