package states

import (
	"fmt"
	"os"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
)

type GameOverState struct{}

func (s *GameOverState) Execute(sm *game.StateMachine, ctx *game.Context) {
	utils.ClearScreen()
	fmt.Println("O jogo terminou.")

	fmt.Println("Pontuação:")
	fmt.Println(ctx.Player.Name, ": ", ctx.Player.Score)
	fmt.Println(ctx.Computer.Name, ": ", ctx.Computer.Score)

	os.Exit(0)
}
