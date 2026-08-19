# tarot
Jeu du tarot

## Game architecture

```
franch-tarot/
├── go.mod
├── main.go                 # Ebiten app entry point
├── game/
│   ├── card.go            # Card definition & logic
│   ├── deck.go            # Deck management
│   ├── rules.go           # Game rules
│   └── state.go           # Game state machine
├── renderer/
│   ├── card_renderer.go   # Card rendering logic
│   └── ui.go              # UI components
└── assets/
    └── images/            # Card images, sprites, etc.
```

## TODOs

- Finish the architecture
- Do the card definition
- Make it so the differents modules can speak together
```
Deck Structure: I'd create a Deck type that holds a slice of Cards. It represents the full Tarot deck.

Methods I'd add:

    NewDeck() — Constructor that generates all 78 cards (4 suits × 14 cards + 22 Trump cards). Returns a fully initialized, unshuffled deck.

    Shuffle() — Randomizes the card order. You'll need this before dealing.

    Deal(n int) — Takes the top n cards from the deck and returns them. Modifies the deck by removing those cards.

    Remaining() — Returns how many cards are left in the deck. Useful for game logic.

    Reset() — Clears and regenerates a fresh deck. Handy for restarting games.

Optional but useful:

    String() — Prints the deck state for debugging
    GetCard(index int) — Retrieve a specific card by index (for testing)

```
