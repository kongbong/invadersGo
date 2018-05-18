package components

import "invaders2/ecs"

const StatusActive = 1
const StatusExplode = 2
const StatusDie = 3

type Status interface {
	ecs.Component
	GetStatus() int
}

func NewStatus(status int) Status {
	return &implStatus{status}
}

type implStatus struct {
	status int
}

func (s *implStatus) GetType() int {
	return CompTypeStatus
}

func (s *implStatus) GetStatus() int {
	return s.status
}
