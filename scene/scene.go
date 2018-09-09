// This package is a factory/helper for scene management.
// Allows for more readable code and to easily track game state.

package scene

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	"github.com/fanus/prog_50016_blackjack/blackjack"
)

type Scene struct {
	game      *engine.Engine
	blackJack *blackjack.BlackJack

	cardBack  *render.Texture
	cardSheet *render.Image

	gameObjects map[string]*GameObject

	newGameNotice *GameObject
	turnNotice    *GameObject
	wonNotice     *GameObject
	lostNotice    *GameObject
	bustNotice    *GameObject
}

func New(g *engine.Engine, bj *blackjack.BlackJack) *Scene {
	s := &Scene{game: g, blackJack: bj}
	s.gameObjects = make(map[string]*GameObject, 0)

	// Back of a card texture
	cbImg := render.NewImage()
	cbImg.LoadImage("assets/card_back.png")
	s.cardBack = render.NewTexture(cbImg)

	// Sprite sheet of all cards
	ssImg := render.NewImage()
	ssImg.LoadImage("assets/card_sheet.png")
	s.cardSheet = ssImg

	// Background texture
	bgImg := render.NewImage()
	bgImg.LoadImage("assets/bg.png")
	bg := render.NewTexture(bgImg)

	// Background
	s.NewGameObject(g.NewEntity()).
		AddTransformComponent(&std.Vector3{0, 0, 1}, &std.Vector3{0, 0, 0}, &std.Vector3{10.24, 7.68, 1}).
		AddMaterialComponent(bg)

	// Static deck of cards
	s.NewGameObject(g.NewEntity()).
		AddTransformComponent(&std.Vector3{-5.12, 4.8, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.8, 2.7, 1}).
		AddMaterialComponent(s.cardBack)

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
				AddTransformComponent(&std.Vector3{0 + float32(idx), -4.8, float32(idx) * -0.05}, &std.Vector3{0, 0, 0}, &std.Vector3{1.8, 2.7, 1}).
				AddCardComponent(id, true, s.cardSheet)
		}
	}
	for idx, id := range dealerHand {
		if s.gameObjects[id] == nil {
			visible := false
			if idx == 0 {
				visible = true
			}
			s.gameObjects[id] = s.NewGameObject(s.game.NewEntity()).
				AddTransformComponent(&std.Vector3{1.3 + float32(idx), 4.8, float32(idx) * -0.05}, &std.Vector3{0, 0, 0}, &std.Vector3{1.8, 2.7, 1}).
				AddCardComponent(id, visible, s.cardSheet)
		}
	}
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
			s.flipDealerCards()
			s.turnNotice.Hide()
			s.wonNotice.Show()
		}
		break
	case blackjack.Lost:
		{
			s.flipDealerCards()
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

func (s *Scene) flipDealerCards() {
	_, dealerHand := s.blackJack.Hands()
	for idx, id := range dealerHand {
		if s.gameObjects[id] != nil && idx > 0 {
			blackjack.GetCard(s.gameObjects[id].entity).Flip()
		}
	}
}

func (s *Scene) createNotices() {
	s.newGameNotice = s.createNotice("assets/newgame.png", 1.11)
	s.turnNotice = s.createNotice("assets/turn.png", 1.48)
	s.bustNotice = s.createNotice("assets/bust.png", 1.48)
	s.wonNotice = s.createNotice("assets/won.png", 1.48)
	s.lostNotice = s.createNotice("assets/lost.png", 1.48)
}

func (s *Scene) createNotice(texturePath string, height float32) *GameObject {
	i := render.NewImage()
	i.LoadImage(texturePath)
	t := render.NewTexture(i)

	return s.NewGameObject(s.game.NewEntity()).
		Hide().
		AddTransformComponent(&std.Vector3{-6, -3, -0.1}, &std.Vector3{0, 0, 0}, &std.Vector3{3.08, height, 1}).
		AddMaterialComponent(t)
}
