package game

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{
		Cards: []Card{},
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
				deck.Cards = append(deck.Cards, card)
			}
		} else {
			// Case Trump
			for i := range 22 {
				oudler := i == 0 || i == 1 || i == 21 // These are the value where we set the oudler
				card := Card{
					Suit:   suit,
					Rank:   0,
					Oudler: oudler,
					Number: i,
				}

				card.Name = card.Suit.String() + " number " + strconv.Itoa(card.Number)
				deck.Cards = append(deck.Cards, card)
			}
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Deal(n int) []Card {
	deckLength := len(d.Cards)
	hand := slices.Clone(d.Cards[deckLength-n:])
	slices.Delete(d.Cards, deckLength-n, deckLength)
	return hand
}

func (d *Deck) PrintDeck() {
	for i := range len(d.Cards) {
		fmt.Println(d.Cards[i])
	}
}
