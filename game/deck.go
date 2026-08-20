package game

import "strconv"

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
			// Case Trump
			for i := range 22 {
				oudler := false
				if i == 0 || i == 1 || i == 21 {
					oudler := true
				}
				card := Card{
					Suit:   suit,
					Rank:   0,
					Oudler: oudler,
					Number: i,
				}

				card.Name = card.Suit.String() + " number " + strconv.Itoa(card.Number)
				deck.cards = append(deck.cards, card)
			}
		}
	}
	return deck
}
