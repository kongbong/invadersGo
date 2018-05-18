package scenes

import (
	"fmt"
	"invaders2/ecs"
	"invaders2/systems"
)

//var background = globals.GetImage("imgs/bg.png")

type GameScene struct {
	world ecs.World
}

func (s *GameScene) Init() {
	fmt.Println("GameScene Init")
	s.world = ecs.NewWorld()
	s.world.AddSystem(systems.NewDrawSystem())
	s.world.AddSystem(systems.NewInvadersSystem())

	// dst := image.NewRGBA(image.Rect(0, 0, globals.Width, globals.Height))
	// gift.New().Draw(dst, background)

	// globals.PrintImage(dst)
}

func (t *GameScene) OnDestroy() {

}

func (t *GameScene) Tick(tickCnt uint64) {
	t.world.Tick(tickCnt)
}
