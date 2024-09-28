package states

import (
	"fmt"
	"super-trunfo/internal/game"
)

type DrawState struct{}

func (s *DrawState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("Empate!")
	sm.SetState(&GameOverState{}) // Transição para o estado de jogo terminado
}
