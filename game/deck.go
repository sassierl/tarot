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
				card := Card{
					Suit:   suit,
					Rank:   Rank(i + 1),
					Oudler: false,
				}

				card.Name = card.Rank.String() + " of " + card.Suit.String()
				deck.cards = append(deck.cards, card)
			}
		} else {
			// TODO: finish this
		}
	}
	return deck
}
