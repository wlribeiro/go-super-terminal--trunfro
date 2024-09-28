package states

import (
	"fmt"
	"super-trunfo/internal/cards"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
)

type PlayerTurnState struct{}

func (s *PlayerTurnState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("É a vez do jogador.")
	utils.AwaitRead()
	utils.ClearScreen()

	ctx.Player.Score += 1

	card := ctx.Player.Deck[0]

	if ctx.Computer.ChosenAttribute != "" {
		ctx.Computer.ChosenCard.MountCardInterface()
		fmt.Println("Atributo escolhido pelo computador: ", ctx.Computer.ChosenAttribute)
		fmt.Println("Sua carta:")
		card.MountCardInterface()
		ctx.Player.ChosenCard = card
		ctx.Player.ChosenAttribute = ctx.Computer.ChosenAttribute
		utils.AwaitRead()

		sm.SetState(&BattlePhase{})
		return
	}

	card.MountCardInterface()
	// input to choose the attribute to battle
	fmt.Println("Escolhendo o atributo para batalha...")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Valor inválido")
		sm.SetState(&PlayerTurnState{})
		return
	} else if input < 0 || input > 2 {
		fmt.Println("Opção não encontrada")
		sm.SetState(&PlayerTurnState{})
		return
	}

	ctx.Player.ChosenCard = card
	ctx.Player.ChosenAttribute = cards.AttributeMap[cards.Attribute(input)]

	sm.SetState(&ComputerTurnState{})

}
