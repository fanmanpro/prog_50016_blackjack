package main

import (
	// engine & systems
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/keyboard"
	"github.com/autovelop/playthos/render"

	// engine systems & platforms that produce init side effects
	_ "github.com/autovelop/playthos/glfw"
	_ "github.com/autovelop/playthos/glfw/keyboard"
	_ "github.com/autovelop/playthos/opengl"
	_ "github.com/autovelop/playthos/platforms/windows"

	"github.com/fanus/prog_50016_blackjack/blackjack"
	"github.com/fanus/prog_50016_blackjack/scene"
)

func main() {
	game := engine.New("BlackJack", &engine.Settings{
		false,
		1024,
		768,
		false,
	})

	// music := game.NewEntity()
	// sound := audio.NewSound()
	// clip := audio.NewClip()
	// clip.LoadClip("assets/music.wav")
	// sound.Set(clip)
	// music.AddComponent(sound)

	// src := audio.NewSource()
	// src.Set(&std.Vector3{0, 0, 0}, false, true)
	// src.PlaySound(sound)
	// music.AddComponent(src)

	kb := game.Listener(&keyboard.Keyboard{})

	// Back of a card texture
	cbImg := render.NewImage()
	cbImg.LoadImage("assets/card_back.png")
	cb := render.NewTexture(cbImg)

	// Empty Black Jack game
	bj := blackjack.New()

	// Empty scene with reference to Black Jack
	gameScene := scene.New(game, bj, cb)

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
