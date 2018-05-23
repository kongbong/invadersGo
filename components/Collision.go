package components

import (
	"image"

	"github.com/kongbong/invadersGo/ecs"
)

type Collision interface {
	ecs.Component
	GetWidth() int
	GetHeight() int
}

func NewCollision(rect image.Rectangle) Collision {
	return &implCollision{rect.Max.X - rect.Min.X, rect.Max.Y - rect.Min.Y}
}

type implCollision struct {
	width, height int
}

func (c *implCollision) GetType() int {
	return CompTypeCollision
}

func (c *implCollision) GetWidth() int {
	return c.width
}

func (c *implCollision) GetHeight() int {
	return c.height
}
