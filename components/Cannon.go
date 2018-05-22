package components

import "invadersGo/ecs"

type Cannon interface {
	ecs.Component
}

func NewCannon() Cannon {
	return &implCannon{}
}

type implCannon struct {
}

func (c *implCannon) GetType() int {
	return CompTypeCannon
}
