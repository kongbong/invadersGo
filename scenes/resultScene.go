package scenes

import (
	"invadersGo/ecs"
	"invadersGo/globals"
)

func NewResultScene() ecs.Scene {
	return &ResultScene{}
}

type ResultScene struct {
}

func (t *ResultScene) Init() {
	startScreen := globals.GetImage("imgs/start.png")
	globals.PrintImage(startScreen)
}

func (t *ResultScene) OnDestroy() {

}

func (t *ResultScene) Tick(tickCnt uint64) {

}
