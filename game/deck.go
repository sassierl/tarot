package game

type Deck []Card

func NewDeck() *Deck {
	deck := Deck{cards: []Card}
}
