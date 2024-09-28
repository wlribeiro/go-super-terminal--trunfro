package states

import (
	"fmt"
	"super-trunfo/internal/game"
)

type PlayerWonState struct{}

func (s *PlayerWonState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("O jogador ganhou!")

	card, fail := ctx.Computer.DrawCard()
	fmt.Println("quantidade de cartas", len(ctx.Computer.Deck))
	if fail != nil {
		sm.SetState(&GameOverState{})
		return
	}

	ctx.Computer.ResetChoice()

	ctx.Player.AddCard(card)
	card, fail = ctx.Player.DrawCard()
	if fail != nil {
		sm.SetState(&GameOverState{})
		return
	}
	ctx.Player.AddCard(ctx.Player.ChosenCard)
	ctx.Player.ResetChoice()

	if !ctx.Computer.HaveCards() {
		sm.SetState(&GameOverState{})
		return
	}
	sm.SetState(&PlayerTurnState{})

}
