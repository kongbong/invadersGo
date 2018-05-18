package ecs

import (
	"sync/atomic"
)

type Entity interface {
	Id() uint64
}

type Component interface {
	GetType() int
}

type System interface {
	Init(w World)
	Tick(tickCnt uint64)
	Register(id uint64, c Component)
	Unregister(id uint64, componentType int)
}

type implEntity struct {
	id uint64
}

var nextId uint64

func NewEntity() Entity {
	id := atomic.AddUint64(&nextId, 1)
	return &implEntity{id: id}
}

func (e *implEntity) Id() uint64 {
	return e.id
}
