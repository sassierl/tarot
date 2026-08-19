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
