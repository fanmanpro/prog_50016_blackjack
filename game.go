package main

import (
	// engine & systems
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/keyboard"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"

	// engine systems & platforms that produce init side effects
	_ "github.com/autovelop/playthos/glfw"
	_ "github.com/autovelop/playthos/glfw/keyboard"
	_ "github.com/autovelop/playthos/opengl"
	_ "github.com/autovelop/playthos/platforms/windows"

	"./blackjack"
	"./scene"
)

func main() {
	game := engine.New("BlackJack", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	kb := game.Listener(&keyboard.Keyboard{})

	// Back of a card texture
	cbImg := render.NewImage()
	cbImg.LoadImage("assets/card_back.png")
	cb := render.NewTexture(cbImg)

	// Empty Black Jack game
	bj := blackjack.New()

	// Empty scene with reference to Black Jack
	gameScene := scene.New(game, bj, cb)

	// Background
	bgImg := render.NewImage()
	bgImg.LoadImage("assets/bg.png")
	bg := render.NewTexture(bgImg)
	gameScene.NewGameObject(game.NewEntity()).
		AddTransformComponent(&std.Vector3{0, 0, 1}, &std.Vector3{0, 0, 0}, &std.Vector3{10.24, 7.68, 1}).
		AddMaterialComponent(bg)

	// Static deck of cards
	gameScene.NewGameObject(game.NewEntity()).
		AddTransformComponent(&std.Vector3{-5, 5, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
		AddMaterialComponent(cb)

	//NewGameObject(game.NewEntity()).
	//	NewTransform(&std.Vector3{2, -4, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
	//	NewMaterial("background.png", &std.Color{1, 1, 1, 1})
	//NewGameObject(game.NewEntity()).
	//	NewTransform(&std.Vector3{-2, -4, 0}, &std.Vector3{0, 0, 0}, &std.Vector3{1.3, 2, 1}).
	//	NewMaterial("background.png", &std.Color{1, 1, 1, 1})

	kb.On(keyboard.KeySpace, func(action ...int) {
		switch action[0] {
		case keyboard.ActionPress:
			if bj.GameState() == blackjack.NewGame {
				bj.Reset()
				gameScene.ClearCards()

				bj.Shuffle()
				bj.Deal()
				gameScene.UpdateCards()
			} else if bj.GameState() == blackjack.Bust || bj.GameState() == blackjack.Won || bj.GameState() == blackjack.Lost {
				bj.Reset()
				gameScene.ClearCards()
				bj.Shuffle()
				bj.Deal()
				gameScene.UpdateCards()
			}
			gameScene.UpdateGameState()
		}
	})

	kb.On(keyboard.KeyLeft, func(action ...int) {
		switch action[0] {
		case keyboard.ActionPress:
			if bj.GameState() == blackjack.Turn {
				bj.Stand()
				gameScene.UpdateCards()
				gameScene.UpdateGameState()
			}
		}
	})

	kb.On(keyboard.KeyRight, func(action ...int) {
		switch action[0] {
		case keyboard.ActionPress:
			if bj.GameState() == blackjack.Turn {
				bj.HitPlayer()
				gameScene.UpdateCards()
				gameScene.UpdateGameState()
			}
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
