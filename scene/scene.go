// This package is a factory/helper for scene management.
// Allows for more readable code and to easily track game state.

package scene

import (
	"fmt"

	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	"../blackjack"
)

type Scene struct {
	game      *engine.Engine
	blackJack *blackjack.BlackJack
	cardBack  *render.Texture

	gameObjects map[string]*GameObject

	newGameNotice *GameObject
	turnNotice    *GameObject
	wonNotice     *GameObject
	lostNotice    *GameObject
	bustNotice    *GameObject
}

func New(g *engine.Engine, bj *blackjack.BlackJack, cb *render.Texture) *Scene {
	s := &Scene{game: g, blackJack: bj, cardBack: cb}
	s.gameObjects = make(map[string]*GameObject, 0)

	s.createNotices()
	s.UpdateGameState()
	return s
}

func (s *Scene) NewGameObject(e *engine.Entity) *GameObject {
	return &GameObject{e}
}

// Updates the whole scene based on the state of the Black Jack game
func (s *Scene) UpdateCards() {
	playerHand, dealerHand := s.blackJack.Hands()
	for idx, id := range playerHand {
		if s.gameObjects[id] == nil {
			s.gameObjects[id] = s.NewGameObject(s.game.NewEntity()).
				AddTransformComponent(&std.Vector3{-1 + float32(idx), -5, float32(idx) * -0.05}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
				AddCardComponent(id, true, s.cardBack)
		}
	}
	for idx, id := range dealerHand {
		if s.gameObjects[id] == nil {
			visible := false
			if idx == 0 {
				visible = true
			}
			s.gameObjects[id] = s.NewGameObject(s.game.NewEntity()).
				AddTransformComponent(&std.Vector3{1 + float32(idx), 5, float32(idx) * -0.05}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
				AddCardComponent(id, visible, s.cardBack)
		}
	}
	fmt.Println(playerHand, dealerHand, s.blackJack.Deck())
}

func (s *Scene) UpdateGameState() {
	gameState := s.blackJack.GameState()
	switch gameState {
	case blackjack.NewGame:
		{
			s.newGameNotice.Show()
		}
		break
	case blackjack.Bust:
		{
			s.turnNotice.Hide()
			s.bustNotice.Show()
		}
		break
	case blackjack.Turn:
		{
			s.newGameNotice.Hide()
			s.bustNotice.Hide()
			s.wonNotice.Hide()
			s.lostNotice.Hide()
			s.turnNotice.Show()
		}
		break
	case blackjack.Won:
		{
			s.turnNotice.Hide()
			s.wonNotice.Show()
		}
		break
	case blackjack.Lost:
		{
			s.turnNotice.Hide()
			s.lostNotice.Show()
		}
		break
	}
}

func (s *Scene) ClearCards() {
	for _, gameObject := range s.gameObjects {
		s.game.DeleteEntity(gameObject.entity)
	}
	s.gameObjects = make(map[string]*GameObject, 0)
}

func (s *Scene) createNotices() {
	s.newGameNotice = s.createNotice("assets/newgame.png")
	s.turnNotice = s.createNotice("assets/turn.png")
	s.bustNotice = s.createNotice("assets/bust.png")
	s.wonNotice = s.createNotice("assets/won.png")
	s.lostNotice = s.createNotice("assets/lost.png")
}

func (s *Scene) createNotice(texturePath string) *GameObject {
	i := render.NewImage()
	i.LoadImage(texturePath)
	t := render.NewTexture(i)

	return s.NewGameObject(s.game.NewEntity()).
		Hide().
		AddTransformComponent(&std.Vector3{0, 0, 1}, &std.Vector3{0, 0, 0}, &std.Vector3{6, 1.5, 1}).
		AddMaterialComponent(t)
}
