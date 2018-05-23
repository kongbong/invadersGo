package components

import "github.com/kongbong/invadersGo/ecs"

type Beam interface {
	ecs.Component
}

func NewBeam() Beam {
	return &implBeam{}
}

type implBeam struct {
}

func (b *implBeam) GetType() int {
	return CompTypeBeam
}
