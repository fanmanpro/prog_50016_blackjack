package blackjack

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
)

type Card struct {
	engine.Component
	id string

	Visible bool

	front *render.Texture
	back  *render.Texture
}

func NewCard(id string, visible bool, cardFront *render.Texture, cardBack *render.Texture) *Card {
	return &Card{id: id, Visible: visible, front: cardFront, back: cardBack}
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
		render.GetMaterial(c.Entity()).SetTexture(c.front)
	} else {
		render.GetMaterial(c.Entity()).SetTexture(c.back)
	}
}
