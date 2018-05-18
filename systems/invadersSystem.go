package systems

import (
	"fmt"
	"image"
	"invadersGo/components"
	"invadersGo/ecs"
	"invadersGo/globals"
	"math/rand"
)

var aliensPerRow = 8
var aliensStartCol = 100
var alienSize = 30

// sprites
var cannonSprite = image.Rect(20, 47, 38, 59)
var cannonExplode = image.Rect(0, 47, 16, 57)
var alien1Sprite = image.Rect(0, 0, 20, 14)
var alien1aSprite = image.Rect(20, 0, 40, 14)
var alien2Sprite = image.Rect(0, 14, 20, 26)
var alien2aSprite = image.Rect(20, 14, 40, 26)
var alien3Sprite = image.Rect(0, 27, 20, 40)
var alien3aSprite = image.Rect(20, 27, 40, 40)
var alienExplode = image.Rect(0, 60, 16, 68)
var beamSprite = image.Rect(20, 60, 22, 65)
var bombSprite = image.Rect(0, 70, 10, 79)

var alienDirection int = 1
var collidesWall bool = false
var bombProbability = 0.005
var bombSpeed = 10

type invadersSystem struct {
	invaders map[uint64]components.Position
	bombs    map[uint64]components.Position
	src      image.Image
	world    ecs.World
}

func NewInvadersSystem() ecs.System {
	return &invadersSystem{}
}

func (s *invadersSystem) Init(w ecs.World) {
	s.src = globals.GetImage("imgs/sprites.png")
	s.invaders = make(map[uint64]components.Position)
	s.bombs = make(map[uint64]components.Position)
	s.world = w
	w.Subscribe(components.CompTypePosition, s)

	// populate the aliens
	for i := aliensStartCol; i < aliensStartCol+(alienSize*aliensPerRow); i += alienSize {
		s.createAlien(i, 30, alien1Sprite, alien1aSprite, 30)
	}
	for i := aliensStartCol; i < aliensStartCol+(30*aliensPerRow); i += alienSize {
		s.createAlien(i, 55, alien2Sprite, alien2aSprite, 20)
	}
	for i := aliensStartCol; i < aliensStartCol+(30*aliensPerRow); i += alienSize {
		s.createAlien(i, 80, alien3Sprite, alien3aSprite, 10)
	}
}

func (s *invadersSystem) Register(id uint64, c ecs.Component) {

}

func (s *invadersSystem) Unregister(id uint64, componentType int) {
	if _, ok := s.invaders[id]; ok {
		delete(s.invaders, id)
	} else if _, ok := s.bombs[id]; ok {
		fmt.Println("Unregister bomb", id)
		delete(s.bombs, id)
	}
}

func (s *invadersSystem) Tick(tickCnt uint64) {

	for id, p := range s.bombs {
		p.SetY(p.Y() + bombSpeed)
		if p.Y() >= globals.Height {
			fmt.Println("remove bomb", id, p.Y())
			globals.GDispatcher.Dispatch(func() {
				s.world.RemoveEntity(id)
			})
		}
	}

	if collidesWall {
		alienDirection *= -1
		collidesWall = false
		for _, p := range s.invaders {
			p.SetY(p.Y() + 10)
		}
		return
	}

	for _, p := range s.invaders {
		p.SetX(p.X() + 5*alienDirection)
		if p.X() < alienSize || p.X() > globals.Width-(2*alienSize) {
			collidesWall = true
		}

		// drop torpedoes
		if rand.Float64() < bombProbability {
			s.dropBomb(p)
		}
	}
}

func (s *invadersSystem) dropBomb(p components.Position) {
	e := s.world.SpawnEntity()
	pos := components.NewPosition(p.X()+7, p.Y())
	s.world.AddComponent(e.Id(), pos)
	s.world.AddComponent(e.Id(), components.NewStatus(components.StatusActive))
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, bombSprite, bombSprite, bombSprite))

	s.bombs[e.Id()] = pos
}

// used for creating alien sprites
func (s *invadersSystem) createAlien(x, y int, sprite, alt image.Rectangle, points int) {
	e := s.world.SpawnEntity()
	pos := components.NewPosition(x, y)
	s.world.AddComponent(e.Id(), pos)
	s.world.AddComponent(e.Id(), components.NewStatus(components.StatusActive))
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, sprite, alt, alienExplode))

	s.invaders[e.Id()] = pos
}
