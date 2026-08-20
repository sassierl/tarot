package game

type Phase int

const (
	PhaseDeal Phase = iota
	PhaseAuction
	PhasePlay
	PhaseScoring
)

type GameState struct {
	Phase       Phase
	CurrentTurn int
	Taker       *Player
	Bid         int
	Round       int
}

func (gs *GameState) NextPhase() {
	gs.Phase++
	if gs.Phase > PhaseScoring {
		gs.Phase = PhaseDeal
	}
}

func (gs *GameState) SetTaker(player *Player, bid int) {
	gs.Taker = player
	gs.Bid = bid
}
