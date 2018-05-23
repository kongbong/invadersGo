package components

import "github.com/kongbong/invadersGo/ecs"

const (
	StatusActive = iota
	StatusExplode
	StatusDie
)

type Status interface {
	ecs.Component
	GetStatus() int
	SetStatus(status int)
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

func (s *implStatus) SetStatus(status int) {
	s.status = status
}
