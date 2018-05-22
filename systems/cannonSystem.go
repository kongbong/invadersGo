package systems

import (
	"image"
	"invadersGo/components"
	"invadersGo/ecs"
	"invadersGo/globals"

	"github.com/google/gxui"
)

type cannonSystem struct {
	src          image.Image
	world        ecs.World
	cannonPos    components.Position
	cannonStatus components.Status
	beamPos      components.Position
	beamId       uint64
}

func NewCannonSystem(sprite image.Image) ecs.System {
	return &cannonSystem{src: sprite}
}

var cannonSprite = image.Rect(20, 47, 38, 59)
var cannonExplode = image.Rect(0, 47, 16, 57)
var beamSprite = image.Rect(20, 60, 22, 65)

func (s *cannonSystem) Init(w ecs.World) {
	s.world = w
	s.createCannon(50, 250)
	s.world.Subscribe(components.CompTypePosition, s)
}

func (s *cannonSystem) Tick(tickCnt uint64) {

	if s.beamPos != nil {
		s.beamPos.SetY(s.beamPos.Y() - 10)
		if s.beamPos.Y() < 0 {
			s.world.RemoveEntity(s.beamId)
		}
	}

	if s.cannonStatus.GetStatus() != components.StatusActive {
		return
	}

	if globals.GInputManager.IsKeyDown(gxui.KeyRight, gxui.ModNone) && s.cannonPos.X() < globals.Width-(2*alienSize) {
		s.cannonPos.SetX(s.cannonPos.X() + 10)
	} else if globals.GInputManager.IsKeyDown(gxui.KeyLeft, gxui.ModNone) && s.cannonPos.X() > alienSize {
		s.cannonPos.SetX(s.cannonPos.X() - 10)
	}

	if globals.GInputManager.IsKeyDown(gxui.KeySpace, gxui.ModNone) {
		s.shootBeam(s.cannonPos)
	}
}

func (s *cannonSystem) Register(id uint64, c ecs.Component) {

}

func (s *cannonSystem) Unregister(id uint64, componentType int) {
	if s.beamId == id {
		s.beamPos = nil
	}
}

func (s *cannonSystem) shootBeam(p components.Position) {
	if s.beamPos != nil {
		return
	}

	e := s.world.SpawnEntity()
	pos := components.NewPosition(p.X()+7, 250)
	s.world.AddComponent(e.Id(), pos)
	s.world.AddComponent(e.Id(), components.NewStatus(components.StatusActive))
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, beamSprite, beamSprite, beamSprite))
	s.world.AddComponent(e.Id(), components.NewCollision(beamSprite))
	s.world.AddComponent(e.Id(), components.NewBeam())

	s.beamPos = pos
	s.beamId = e.Id()
}

// used for creating alien sprites
func (s *cannonSystem) createCannon(x, y int) {
	e := s.world.SpawnEntity()
	pos := components.NewPosition(x, y)
	s.world.AddComponent(e.Id(), pos)
	status := components.NewStatus(components.StatusActive)
	s.world.AddComponent(e.Id(), status)
	s.world.AddComponent(e.Id(), components.NewSprite(s.src, cannonSprite, cannonSprite, cannonExplode))
	s.world.AddComponent(e.Id(), components.NewCollision(cannonSprite))
	s.world.AddComponent(e.Id(), components.NewCannon())

	s.cannonPos = pos
	s.cannonStatus = status
}
