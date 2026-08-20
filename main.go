package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sassierl/tarot/game"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	deck := game.NewDeck()
	fmt.Printf("Deck has %d cards\n", len(deck.cards))

	// // Test Shuffle
	// deck.Shuffle()
	// fmt.Println("Deck shuffled!")

	// // Test Deal
	// hand := deck.Deal(5)
	// fmt.Printf("Dealt %d cards\n", len(hand))
}

// old main for ebiten
// func main() {
// 	ebiten.SetWindowSize(640, 480)
// 	ebiten.SetWindowTitle("Hello, World!")
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}
// }
