package blackjack

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

var cardSheet map[string]*std.Vector2

func init() {
	cardSheet = map[string]*std.Vector2{
		"_":   &std.Vector2{X: 0, Y: 0},
		"C2":  &std.Vector2{X: 1, Y: 0},
		"C3":  &std.Vector2{X: 2, Y: 0},
		"C4":  &std.Vector2{X: 3, Y: 0},
		"C5":  &std.Vector2{X: 4, Y: 0},
		"C6":  &std.Vector2{X: 5, Y: 0},
		"C7":  &std.Vector2{X: 6, Y: 0},
		"C8":  &std.Vector2{X: 7, Y: 0},
		"C9":  &std.Vector2{X: 8, Y: 0},
		"C10": &std.Vector2{X: 9, Y: 0},
		"CA":  &std.Vector2{X: 10, Y: 0},
		"CJ":  &std.Vector2{X: 0, Y: 1},
		"CK":  &std.Vector2{X: 1, Y: 1},
		"CQ":  &std.Vector2{X: 2, Y: 1},
		"D2":  &std.Vector2{X: 3, Y: 1},
		"D3":  &std.Vector2{X: 4, Y: 1},
		"D4":  &std.Vector2{X: 5, Y: 1},
		"D5":  &std.Vector2{X: 6, Y: 1},
		"D6":  &std.Vector2{X: 7, Y: 1},
		"D7":  &std.Vector2{X: 8, Y: 1},
		"D8":  &std.Vector2{X: 9, Y: 1},
		"D9":  &std.Vector2{X: 0, Y: 2},
		"D10": &std.Vector2{X: 1, Y: 2},
		"DA":  &std.Vector2{X: 2, Y: 2},
		"DJ":  &std.Vector2{X: 3, Y: 2},
		"DK":  &std.Vector2{X: 4, Y: 2},
		"DQ":  &std.Vector2{X: 5, Y: 2},
		"H2":  &std.Vector2{X: 6, Y: 2},
		"H3":  &std.Vector2{X: 7, Y: 2},
		"H4":  &std.Vector2{X: 8, Y: 2},
		"H5":  &std.Vector2{X: 9, Y: 2},
		"H6":  &std.Vector2{X: 10, Y: 2},
		"H7":  &std.Vector2{X: 0, Y: 3},
		"H8":  &std.Vector2{X: 1, Y: 3},
		"H9":  &std.Vector2{X: 2, Y: 3},
		"H10": &std.Vector2{X: 3, Y: 3},
		"HA":  &std.Vector2{X: 4, Y: 3},
		"HJ":  &std.Vector2{X: 5, Y: 3},
		"HK":  &std.Vector2{X: 6, Y: 3},
		"HQ":  &std.Vector2{X: 7, Y: 3},
		"S2":  &std.Vector2{X: 8, Y: 3},
		"S3":  &std.Vector2{X: 9, Y: 3},
		"S4":  &std.Vector2{X: 10, Y: 3},
		"S5":  &std.Vector2{X: 0, Y: 4},
		"S6":  &std.Vector2{X: 1, Y: 4},
		"S7":  &std.Vector2{X: 2, Y: 4},
		"S8":  &std.Vector2{X: 3, Y: 4},
		"S9":  &std.Vector2{X: 4, Y: 4},
		"S10": &std.Vector2{X: 5, Y: 4},
		"SA":  &std.Vector2{X: 6, Y: 4},
		"SJ":  &std.Vector2{X: 7, Y: 4},
		"SK":  &std.Vector2{X: 8, Y: 4},
		"SQ":  &std.Vector2{X: 9, Y: 4},
	}
}

type Card struct {
	engine.Component
	id string

	Visible bool

	texture *render.Texture
}

func NewCard(id string, visible bool, t *render.Texture) *Card {
	return &Card{id: id, Visible: visible, texture: t}
}

func (c *Card) Flip() {
	c.Visible = !c.Visible
	c.UpdateMaterial()
}

func GetCard(e *engine.Entity) *Card {
	return e.Component(&Card{}).(*Card)
}

func (c *Card) UpdateMaterial() {
	if c.Visible {
		c.texture.SetOffset(cardSheet[c.id])
	} else {
		c.texture.SetOffset(cardSheet["_"])
	}
}
