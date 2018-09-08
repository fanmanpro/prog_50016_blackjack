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

func (g *GameObject) Entity() *engine.Entity {
	return g.entity
}

func (g *GameObject) Hide() *GameObject {
	g.entity.SetActive(false)
	return g
}

func (g *GameObject) Show() *GameObject {
	g.entity.SetActive(true)
	return g
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
