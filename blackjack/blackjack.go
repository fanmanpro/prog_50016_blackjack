package blackjack

import (
	"math/rand"
	"time"

	"github.com/autovelop/playthos/render"
)

type GameState int

const (
	NewGame GameState = 0
	Turn    GameState = 1
	Stand   GameState = 2
	Won     GameState = 3
	Bust    GameState = 4
	Lost    GameState = 5
)

type BlackJack struct {
	busted   bool
	player   []string
	dealer   []string
	deck     []string
	cardBack *render.Texture

	gameState GameState
}

func New() *BlackJack {
	bj := &BlackJack{}
	bj.Reset()
	bj.gameState = NewGame
	return bj
}

func (b *BlackJack) Reset() {
	b.player = make([]string, 0)
	b.dealer = make([]string, 0)
	b.deck = make([]string, 52, 52)

	i := 0
	for key, _ := range cardDefinitions {
		b.deck[i] = key
		i++
	}
}

func (b *BlackJack) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(b.deck))
	perm := r.Perm(len(b.deck))
	for i, randIndex := range perm {
		ret[i] = b.deck[randIndex]
	}
	b.deck = ret
}

func (b *BlackJack) HitDealer() {
	b.dealDealer()
	if b.gameState == Stand {
		b.dealerAI()
	}
}

func (b *BlackJack) dealerAI() {
	score := b.calcScore(b.dealer)
	if score < 17 {
		b.HitDealer()
	} else {
		if score > 21 {
			b.updateGameState(Won)
		} else {
			if score > b.calcScore(b.player) {
				b.updateGameState(Lost)
			} else {
				b.updateGameState(Won)
			}
		}
	}
}

func (b *BlackJack) HitPlayer() {
	b.dealPlayer()

	score := b.calcScore(b.player)
	if score > 21 {
		b.Bust()
	} else if score == 21 {
		b.Stand()
	}
}

func (b *BlackJack) Deal() {
	b.HitPlayer()
	b.HitDealer()
	b.HitPlayer()
	b.HitDealer()
	b.updateGameState(Turn)
}

func (b *BlackJack) Stand() {
	b.updateGameState(Stand)
	b.dealerAI()
}

func (b *BlackJack) Bust() {
	b.updateGameState(Bust)
}

func (b *BlackJack) Deck() []string {
	return b.deck
}

func (b *BlackJack) Hands() ([]string, []string) {
	return b.player, b.dealer
}

func (b *BlackJack) GameState() GameState {
	return b.gameState
}

func (b *BlackJack) updateGameState(gameState GameState) {
	b.gameState = gameState
}

func (b *BlackJack) calcScore(hand []string) int {
	aceCount := 0
	score := 0
	for _, card := range hand {
		if cardDefinitions[card].IsAce {
			aceCount++
		}
		score += cardDefinitions[card].Score
	}
	return score
}

func (b *BlackJack) drawCard() string {
	if len(b.deck) <= 0 {
		panic("somehow the deck of cards ran out playing Black Jack...")
	}

	card := b.deck[0]
	b.deck = b.deck[1:]

	return card
}

func (b *BlackJack) dealDealer() {
	card := b.drawCard()
	b.dealer = append(b.dealer, card)
}

func (b *BlackJack) dealPlayer() {
	card := b.drawCard()
	b.player = append(b.player, card)
}
