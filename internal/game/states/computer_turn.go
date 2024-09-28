package states

import (
	"fmt"
	"super-trunfo/internal/cards"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
)

type ComputerTurnState struct{}

func (s *ComputerTurnState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("É a vez do computador.")

	ctx.Computer.ChosenCard = ctx.Computer.Deck[0]

	if ctx.Player.ChosenAttribute != "" {
		ctx.Computer.ChosenAttribute = ctx.Player.ChosenAttribute
		sm.SetState(&BattlePhase{})
		return
	}

	value := 0
	selectedValue := 0
	var selectedAttribute string

	for _, attributeValue := range cards.AttributeMap {
		value, _ = utils.GetCardAttributeValue(ctx.Computer.ChosenCard, attributeValue)
		if value > selectedValue {
			selectedValue = value
			selectedAttribute = attributeValue
		}
	}

	ctx.Computer.ChosenAttribute = selectedAttribute

	sm.SetState(&PlayerTurnState{})
}
