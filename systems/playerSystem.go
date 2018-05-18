package systems

import (
	"image"
	"invadersGo/components"
	"invadersGo/ecs"
	"invadersGo/globals"

	"github.com/google/gxui"
)

type playerSystem struct {
	src          image.Image
	world        ecs.World
	playerPos    components.Position
	playerStatus components.Status
	beamPos      components.Position
	beamId       uint64
}

func NewPlayerSystem(sprite image.Image) ecs.System {
	return &playerSystem{src: sprite}
}

var cannonSprite = image.Rect(20, 47, 38, 59)
var cannonExplode = image.Rect(0, 47, 16, 57)
var beamSprite = image.Rect(20, 60, 22, 65)

func (s *playerSystem) Init(w ecs.World) {
	s.world = w
	s.createPlayer(50, 250)
	s.world.Subscribe(components.CompTypePosition, s)
}

func (s *playerSystem) Tick(tickCnt uint64) {

	if s.beamPos != nil {
		s.beamPos.SetY(s.beamPos.Y() - 10)
		if s.beamPos.Y() < 0 {
			s.world.RemoveEntity(s.beamId)
		}
	}

	if s.playerStatus.GetStatus() != components.StatusActive {
		return
	}

	if globals.GInputManager.IsKeyDown(gxui.KeyRight, gxui.ModNone) && s.playerPos.X() < globals.Width-(2*alienSize) {
		s.playerPos.SetX(s.playerPos.X() + 10)
	} else if globals.GInputManager.IsKeyDown(gxui.KeyLeft, gxui.ModNone) && s.playerPos.X() > alienSize {
		s.playerPos.SetX(s.playerPos.X() - 10)
	}

	if globals.GInputManager.IsKeyDown(gxui.KeySpace, gxui.ModNone) {
		s.shootBeam(s.playerPos)
	}
}

func (s *playerSystem) Register(id uint64, c ecs.Component) {

}

func (s *playerSystem) Unregister(id uint64, componentType int) {
	if s.beamId == id {
		s.beamPos = nil
	}
}

func (s *playerSystem) shootBeam(p components.Position) {
	if s.beamPos != nil {
		return
	}

	e := s.world.SpawnEntity()
	pos := components.NewPosition(p.X()+7, 250)
	s.world.AddComponent(e.Id(), pos)
	s.world.AddComponent(e.Id(), components.NewStatus(components.StatusActive))
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, beamSprite, beamSprite, beamSprite))

	s.beamPos = pos
	s.beamId = e.Id()
}

// used for creating alien sprites
func (s *playerSystem) createPlayer(x, y int) {
	e := s.world.SpawnEntity()
	pos := components.NewPosition(x, y)
	s.world.AddComponent(e.Id(), pos)
	status := components.NewStatus(components.StatusActive)
	s.world.AddComponent(e.Id(), status)
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, cannonSprite, cannonSprite, cannonExplode))

	s.playerPos = pos
	s.playerStatus = status
}
