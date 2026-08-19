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
	Jack
	Knight
	Queen
	King
)

type Card struct {
	Name   string
	Suit   Suit
	Rank   Rank
	Oudler bool
	Number int // Only for trump card
}

func (s Suit) String() string {
	switch s {
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	case Trump:
		return "Trump"
	default:
		return "Unknown"
	}
}

func (r Rank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Knight:
		return "Knight"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}

func AllSuits() []Suit {
	return []Suit{Diamonds, Hearts, Clubs, Spades, Trump}
}
