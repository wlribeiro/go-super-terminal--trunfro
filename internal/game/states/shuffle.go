package states

import (
	"fmt"
	"math/rand"
	"super-trunfo/internal/cards"
	"super-trunfo/internal/game"
	"super-trunfo/utils"
	"time"
)

type ShuffleState struct {
}

func privateShuffle(cards []cards.Card) {
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func (s *ShuffleState) Execute(sm *game.StateMachine, ctx *game.Context) {
	fmt.Println("Embaralhando as cartas...")
	utils.AwaitRead()

	privateShuffle(ctx.Cards)

	sm.SetState(&DealCardsState{})
}
