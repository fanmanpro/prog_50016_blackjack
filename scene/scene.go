// This file's code is a factory/helper for scene management.
// Allows for more readable code and to easily track game state.

package scene

import (
	"fmt"

	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	"../blackjack"
)

type GameObject struct {
	entity *engine.Entity
}

func NewGameObject(e *engine.Entity) *GameObject {
	return &GameObject{e}
}

func (g *GameObject) Entity() *engine.Entity {
	return g.entity
}

func (g *GameObject) AddTransformComponent(p *std.Vector3, r *std.Vector3, s *std.Vector3) *GameObject {
	t := std.NewTransform()
	t.Set(p, r, s)
	g.entity.AddComponent(t)
	return g
}

func (g *GameObject) AddMaterialComponent(t *render.Texture) *GameObject {
	m := render.NewMaterial()
	m.Set(&std.Color{1, 1, 1, 1})
	m.SetTexture(t)
	g.entity.AddComponent(m)

	q := render.NewMesh()
	q.Set(std.QuadMesh)
	g.entity.AddComponent(q)
	return g
}

func (g *GameObject) AddCardComponent(id string, visible bool, cardBack *render.Texture) *GameObject {
	i := render.NewImage()
	i.LoadImage(fmt.Sprintf("assets/card_%v.png", id))
	cardFront := render.NewTexture(i)

	c := blackjack.NewCard(id, visible, cardFront, cardBack)
	g.entity.AddComponent(c)

	// Because the matieral is defined by which card it is, we are creating a empty material component first
	m := render.NewMaterial()
	m.Set(&std.Color{1, 1, 1, 1})
	g.entity.AddComponent(m)

	q := render.NewMesh()
	q.Set(std.QuadMesh)
	g.entity.AddComponent(q)

	// Once all the components are ready, we update the material for rendering
	c.UpdateMaterial()

	return g
}

type Scene struct {
	game      *engine.Engine
	blackJack *blackjack.BlackJack
	cardBack  *render.Texture

	gameObjects map[string]*GameObject
}

func New(g *engine.Engine, bj *blackjack.BlackJack, cb *render.Texture) *Scene {
	s := &Scene{game: g, blackJack: bj, cardBack: cb}
	s.gameObjects = make(map[string]*GameObject, 0)
	return s
}

// Updates the whole scene based on the state of the Black Jack game
func (s *Scene) UpdateScene() {
	playerHand, dealerHand := s.blackJack.Hands()
	for idx, id := range playerHand {
		if s.gameObjects[id] == nil {
			s.gameObjects[id] = NewGameObject(s.game.NewEntity()).
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
			s.gameObjects[id] = NewGameObject(s.game.NewEntity()).
				AddTransformComponent(&std.Vector3{1 + float32(idx), 5, float32(idx) * -0.05}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
				AddCardComponent(id, visible, s.cardBack)
		}
	}
	fmt.Println(playerHand, dealerHand)
}
