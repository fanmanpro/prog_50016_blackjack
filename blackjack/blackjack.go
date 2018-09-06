package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

var cardDefinitions map[string]*Card

func init() {
	cardDefinitions = map[string]*Card{
		"AS": &Card{Points: 11, IsAce: true},
		"2S": &Card{Points: 2},
		"3S": &Card{Points: 3},
		"4S": &Card{Points: 4},
		"5S": &Card{Points: 5},
	}
}

// MAKE A CARD COMPONENT
type Card struct {
	Points int
	IsAce  bool
}

type BlackJack struct {
	total  int
	busted bool
	hand   map[string]*Card
	deck   []string
}

func New() *BlackJack {
	bj := &BlackJack{}
	bj.hand = make(map[string]*Card, 0)
	bj.deck = make([]string, 5, 5)

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

func (b *BlackJack) Deal() {
	if len(b.deck) <= 0 {
		panic("somehow the deck of cards ran out playing Black Jack...")
	}

	newCard := b.deck[0]
	b.deck = b.deck[1:]
	b.hand[newCard] = cardDefinitions[newCard]

	aceCount := 0
	handPoints := 0
	for _, card := range b.hand {
		if card.IsAce {
			aceCount++
		}
		handPoints += card.Points
	}
	b.total = handPoints
	fmt.Println(b.total)
}

func (b *BlackJack) Hold() {
}

func (b *BlackJack) Bust() {
}
