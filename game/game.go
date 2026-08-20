package game

type Game struct {
	Players []Player
	Deck    *Deck
	State   *GameState
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

// PlayRound executes one complete round
// func (g *Game) PlayRound() error {
// 	for {
// 		switch g.State.Phase {
// 		case PhaseDeal:
// 			if err := g.ExecuteDeal(); err != nil {
// 				return err
// 			}
// 		case PhaseAuction:
// 			if err := g.ExecuteAuction(); err != nil {
// 				return err
// 			}
// 		case PhasePlay:
// 			if err := g.ExecutePlay(); err != nil {
// 				return err
// 			}
// 		case PhaseScoring:
// 			if err := g.ExecuteScoring(); err != nil {
// 				return err
// 			}
// 			return nil
// 		default:
// 			return fmt.Errorf("unknown phase: %v", g.State.Phase)
// 		}
// 	}
// }

func (g *Game) IsGameOver() bool {
	for _, p := range g.Players {
		if p.Score >= 500 { // Example threshold TODO: get the real number
			return true
		}
	}
	return false
}

func (g *Game) GetWinner() Player {
	winner := g.Players[0]
	for _, p := range g.Players {
		if p.Score > winner.Score {
			winner = p
		}
	}
	return winner
}
