package game

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{
		cards: []Card{},
	}

	for _, suit := range AllSuits() {
		if suit != Trump {
			for i := 0; i < 14; i++ {
				deck.cards[i] = Card{Suit: suit; Rank: i + 1, Oudler: false}
				deck.cards[i].Name =
			}
		}
	}
	return deck
}
