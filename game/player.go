package game

type Player struct {
	Name  string
	Hand  []Card
	Score int
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}
