package components

import "github.com/kongbong/invadersGo/ecs"

type Bomb interface {
	ecs.Component
}

func NewBomb() Bomb {
	return &implBomb{}
}

type implBomb struct {
}

func (b *implBomb) GetType() int {
	return CompTypeBomb
}
