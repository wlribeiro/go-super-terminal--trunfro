package player

import (
	"fmt"
	"super-trunfo/internal/cards"
)

type Player struct {
	Name            string
	Deck            []cards.Card
	Score           int
	ChosenCard      cards.Card
	ChosenAttribute string
}

func (p *Player) ResetChoice() {
	p.ChosenCard = cards.Card{}
	p.ChosenAttribute = ""
}

func (p *Player) HaveCards() bool {
	return len(p.Deck) > 0
}

func (p *Player) DrawCard() (cards.Card, error) {
	if len(p.Deck) == 0 {
		return cards.Card{}, fmt.Errorf("no cards to draw")
	}

	card := p.Deck[0]
	p.Deck = p.Deck[1:]
	return card, nil
}

func (p *Player) AddCard(card cards.Card) {
	p.Deck = append(p.Deck, card)
}
