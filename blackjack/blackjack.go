package blackjack

import (
	"math/rand"
	"time"

	"github.com/autovelop/playthos/render"
)

type CardDefinition struct {
	Score int
	IsAce bool
}

var cardDefinitions map[string]*CardDefinition

func init() {
	cardDefinitions = map[string]*CardDefinition{
		"SA":  &CardDefinition{Score: 11, IsAce: true},
		"S2":  &CardDefinition{Score: 2},
		"S3":  &CardDefinition{Score: 3},
		"S4":  &CardDefinition{Score: 4},
		"S5":  &CardDefinition{Score: 5},
		"S6":  &CardDefinition{Score: 6},
		"S7":  &CardDefinition{Score: 7},
		"S8":  &CardDefinition{Score: 8},
		"S9":  &CardDefinition{Score: 9},
		"S10": &CardDefinition{Score: 10},
		"SJ":  &CardDefinition{Score: 10},
		"SQ":  &CardDefinition{Score: 10},
		"SK":  &CardDefinition{Score: 10},
		"CA":  &CardDefinition{Score: 11, IsAce: true},
		"C2":  &CardDefinition{Score: 2},
		"C3":  &CardDefinition{Score: 3},
		"C4":  &CardDefinition{Score: 4},
		"C5":  &CardDefinition{Score: 5},
		"C6":  &CardDefinition{Score: 6},
		"C7":  &CardDefinition{Score: 7},
		"C8":  &CardDefinition{Score: 8},
		"C9":  &CardDefinition{Score: 9},
		"C10": &CardDefinition{Score: 10},
		"CJ":  &CardDefinition{Score: 10},
		"CQ":  &CardDefinition{Score: 10},
		"CK":  &CardDefinition{Score: 10},
		"DA":  &CardDefinition{Score: 11, IsAce: true},
		"D2":  &CardDefinition{Score: 2},
		"D3":  &CardDefinition{Score: 3},
		"D4":  &CardDefinition{Score: 4},
		"D5":  &CardDefinition{Score: 5},
		"D6":  &CardDefinition{Score: 6},
		"D7":  &CardDefinition{Score: 7},
		"D8":  &CardDefinition{Score: 8},
		"D9":  &CardDefinition{Score: 9},
		"D10": &CardDefinition{Score: 10},
		"DJ":  &CardDefinition{Score: 10},
		"DQ":  &CardDefinition{Score: 10},
		"DK":  &CardDefinition{Score: 10},
		"HA":  &CardDefinition{Score: 11, IsAce: true},
		"H2":  &CardDefinition{Score: 2},
		"H3":  &CardDefinition{Score: 3},
		"H4":  &CardDefinition{Score: 4},
		"H5":  &CardDefinition{Score: 5},
		"H6":  &CardDefinition{Score: 6},
		"H7":  &CardDefinition{Score: 7},
		"H8":  &CardDefinition{Score: 8},
		"H9":  &CardDefinition{Score: 9},
		"H10": &CardDefinition{Score: 10},
		"HJ":  &CardDefinition{Score: 10},
		"HQ":  &CardDefinition{Score: 10},
		"HK":  &CardDefinition{Score: 10},
	}
}

type BlackJack struct {
	busted   bool
	player   []string
	dealer   []string
	deck     []string
	cardBack *render.Texture
}

func New() *BlackJack {
	bj := &BlackJack{}
	bj.player = make([]string, 0)
	bj.dealer = make([]string, 0)
	bj.deck = make([]string, 52, 52)

	i := 0
	for key, _ := range cardDefinitions {
		bj.deck[i] = key
		i++
	}

	return bj
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
	if len(b.deck) <= 0 {
		panic("somehow the deck of cards ran out playing Black Jack...")
	}

	drawnCard := b.deck[0]
	b.deck = b.deck[1:]

	b.dealer = append(b.dealer, drawnCard)
}

func (b *BlackJack) HitPlayer() {
	if len(b.deck) <= 0 {
		panic("somehow the deck of cards ran out playing Black Jack...")
	}

	drawnCard := b.deck[0]
	b.deck = b.deck[1:]

	b.player = append(b.player, drawnCard)
}

func (b *BlackJack) calculateScores() {
	// aceCount := 0
	// handScore := 0
	// for _, card := range b.hand {
	// 	if card.IsAce {
	// 		aceCount++
	// 	}
	// 	handScore += card.Score
	// }

	// if handScore > 21 {
	// 	b.Bust()
	// } else if handScore == 21 {
	// 	b.Stand()
	// }
}

func (b *BlackJack) Deal() {
	b.HitPlayer()
	b.HitDealer()
	b.HitPlayer()
	b.HitDealer()
}

func (b *BlackJack) Stand() {
}

func (b *BlackJack) Bust() {
}

func (b *BlackJack) Hands() ([]string, []string) {
	return b.player, b.dealer
}
