package scenes

import (
	"fmt"
	"invadersGo/ecs"
	"invadersGo/globals"
	"invadersGo/systems"
)

func NewGameScene() ecs.Scene {
	return &GameScene{}
}

type GameScene struct {
	world ecs.World
}

func (s *GameScene) Init() {
	fmt.Println("GameScene Init")
	sprite := globals.GetImage("imgs/sprites.png")
	s.world = ecs.NewWorld()
	s.world.AddSystem(systems.NewDrawSystem())
	s.world.AddSystem(systems.NewInvadersSystem(sprite))
	s.world.AddSystem(systems.NewPlayerSystem(sprite))
}

func (t *GameScene) OnDestroy() {

}

func (t *GameScene) Tick(tickCnt uint64) {
	t.world.Tick(tickCnt)
}
