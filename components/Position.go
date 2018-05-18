package components

import (
	"invadersGo/ecs"
)

type Position interface {
	ecs.Component
	X() int
	Y() int
	SetX(x int)
	SetY(y int)
}

func NewPosition(x, y int) Position {
	return &implPosition{x, y}
}

type implPosition struct {
	x, y int
}

func (p *implPosition) GetType() int {
	return CompTypePosition
}

func (p *implPosition) X() int {
	return p.x
}

func (p *implPosition) Y() int {
	return p.y
}

func (p *implPosition) SetX(x int) {
	p.x = x
}

func (p *implPosition) SetY(y int) {
	p.y = y
}
