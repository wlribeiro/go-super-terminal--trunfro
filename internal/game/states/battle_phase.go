package states

import (
	"fmt"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
)

type BattlePhase struct{}

func (s *BattlePhase) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("Fase de batalha")

	playerPoint, _ := utils.GetCardAttributeValue(ctx.Player.ChosenCard, ctx.Player.ChosenAttribute)
	computerPoint, _ := utils.GetCardAttributeValue(ctx.Computer.ChosenCard, ctx.Computer.ChosenAttribute)

	// clear interface
	utils.ClearScreen()

	fmt.Println("Player: ", ctx.Player.Name)
	fmt.Println("Carta:")
	ctx.Player.ChosenCard.MountCardInterface()
	fmt.Println("Atributo: ", ctx.Player.ChosenAttribute)
	fmt.Println()
	fmt.Println("Computer: ", ctx.Computer.Name)
	fmt.Println("Carta:")
	ctx.Computer.ChosenCard.MountCardInterface()
	fmt.Println("Atributo: ", ctx.Computer.ChosenAttribute)

	// await for input to continue

	utils.AwaitRead()
	utils.ClearScreen()

	fmt.Println("Player: ", ctx.Player.ChosenCard.Name)
	fmt.Println("Attribute: ", ctx.Player.ChosenAttribute)
	fmt.Println("Valor: ", playerPoint)

	fmt.Println("X")

	fmt.Println("Player: ", ctx.Computer.ChosenCard.Name)
	fmt.Println("Attribute: ", ctx.Computer.ChosenAttribute)
	fmt.Println("Valor: ", computerPoint)

	if playerPoint > computerPoint {
		sm.SetState(&PlayerWonState{})
		return
	}

	if computerPoint > playerPoint {
		sm.SetState(&ComputerWonState{})
		return
	}

	sm.SetState(&DrawState{})
}
