package main

import (
	"fmt"

	"github.com/kongbong/invadersGo/ecs"
)

func NewGameScene() ecs.Scene {
	return &GameScene{}
}

type GameScene struct {
	world ecs.World
}

func (s *GameScene) Init() {
	fmt.Println("GameScene Init")
	sprite := GetImage("imgs/sprites.png")
	s.world = ecs.NewWorld(dispatcher)
	s.world.AddSystem(NewDrawSystem())
	s.world.AddSystem(NewGameManager())
	s.world.AddSystem(NewInvadersSystem(sprite))
	s.world.AddSystem(NewCannonSystem(sprite))
}

func (t *GameScene) OnDestroy() {

}

func (t *GameScene) Tick(tickCnt uint64) {
	t.world.Tick(tickCnt)
}
