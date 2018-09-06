package main

import (
	// game engine
	"github.com/autovelop/playthos"

	// abstract keyboard system
	"github.com/autovelop/playthos/keyboard"

	// standards library (vectors, colors, quad meshes, math, etc.)
	"github.com/autovelop/playthos/std"

	// glfw window context
	_ "github.com/autovelop/playthos/glfw"

	// glfw input binding that integrates into the keyboard system
	_ "github.com/autovelop/playthos/glfw/keyboard"

	// opengl 4.5/4.3/3.5
	_ "github.com/autovelop/playthos/opengl"

	// windows deployment routines and asset management
	_ "github.com/autovelop/playthos/platforms/windows"

	// meshes, materials, textures, etc.
	render "github.com/autovelop/playthos/render"

	// the Black Jack game
	"./blackjack"
)

func main() {
	game := engine.New("BlackJack", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	kb := game.Listener(&keyboard.Keyboard{})

	// Static deck of cards
	NewGameObject(game.NewEntity()).
		NewTransform(&std.Vector3{0, 5, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
		NewMaterial("background.png", &std.Color{1, 0.3, 0.3, 1})

	bj := blackjack.New()
	bj.Shuffle()

	//NewGameObject(game.NewEntity()).
	//	NewTransform(&std.Vector3{2, -4, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
	//	NewMaterial("background.png", &std.Color{1, 1, 1, 1})
	//NewGameObject(game.NewEntity()).
	//	NewTransform(&std.Vector3{-2, -4, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
	//	NewMaterial("background.png", &std.Color{1, 1, 1, 1})

	kb.On(keyboard.KeyLeft, func(action ...int) {
		switch action[0] {
		case keyboard.ActionPress:
			bj.Hold()
		}
	})

	kb.On(keyboard.KeyRight, func(action ...int) {
		switch action[0] {
		case keyboard.ActionPress:
			bj.Deal()
		}
	})

	kb.On(keyboard.KeyEscape, func(action ...int) {
		switch action[0] {
		case keyboard.ActionRelease:
			game.Stop()
		}
	})

	game.Start()
}

type GameObject struct {
	entity *engine.Entity
}

func NewGameObject(e *engine.Entity) *GameObject {
	return &GameObject{e}
}

func (g *GameObject) NewTransform(p *std.Vector3, r *std.Vector3, s *std.Vector3) *GameObject {
	t := std.NewTransform()
	t.Set(p, r, s)
	g.entity.AddComponent(t)
	return g
}

func (g *GameObject) NewMaterial(f string, c *std.Color) *GameObject {
	m := render.NewMaterial()
	m.Set(c)
	// i := render.NewImage()
	// i.LoadImage(f)
	// t := render.NewTexture(i)
	// m.SetTexture(t)
	g.entity.AddComponent(m)

	q := render.NewMesh()
	q.Set(std.QuadMesh)
	g.entity.AddComponent(q)
	return g
}
