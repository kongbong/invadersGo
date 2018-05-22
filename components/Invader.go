package components

import "invadersGo/ecs"

type Invader interface {
	ecs.Component
	GetScore() int
}

func NewInvader(score int) Invader {
	return &implInvader{score}
}

type implInvader struct {
	score int
}

func (i *implInvader) GetType() int {
	return CompTypeInvader
}

func (i *implInvader) GetScore() int {
	return i.score
}
