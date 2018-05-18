package ecs

type Scene interface {
	Init()
	OnDestroy()
	Tick(tickCnt uint64)
}
