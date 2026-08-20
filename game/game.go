package game

type Game struct {
	Players []Player
	Deck    Deck
	State   GameState
}

func NewGame(players []Player) *Game {
	return &Game{
		Players: players,
		Deck:    NewDeck(),
		State: &GameState{
			Phase:       PhaseDeal,
			CurrentTurn: 0,
			Taker:       nil,
			Bid:         0,
			Round:       1,
		},
	}
}
