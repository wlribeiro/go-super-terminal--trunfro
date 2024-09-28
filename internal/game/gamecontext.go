package game

import (
	"super-trunfo/internal/cards"
	"super-trunfo/internal/player"
)

type Context struct {
	Player   player.Player
	Computer player.Player
	Cards    []cards.Card // Cartas não distribuídas
}
