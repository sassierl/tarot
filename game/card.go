// game/card.go
package game

type Suit int
type Rank int

const (
	Diamonds Suit = iota
	Hearts
	Clubs
	Spades
	Trump // The fool will be numbered 0 in the Trump type
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Page
	Knight
	Queen
	King
)

type Card struct {
	Name   string
	Suit   Suit
	Rank   Rank
	Oudler bool
}
