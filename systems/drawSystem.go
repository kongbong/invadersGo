package systems

import (
	"image"
	"invaders2/components"
	"invaders2/ecs"
	"invaders2/globals"

	"github.com/disintegration/gift"
)

type drawer struct {
	sprite   components.Sprite
	position components.Position
	status   components.Status
}

type drawSystem struct {
	background image.Image
	drawObjs   map[uint64]*drawer
	world      ecs.World
}

func NewDrawSystem() ecs.System {
	return &drawSystem{drawObjs: make(map[uint64]*drawer)}
}

func (s *drawSystem) Init(w ecs.World) {
	s.world = w
	s.background = globals.GetImage("imgs/bg.png")
	w.Subscribe(components.CompTypeSprite, s)
}

func (s *drawSystem) Tick(tickCnt uint64) {
	dst := image.NewRGBA(image.Rect(0, 0, globals.Width, globals.Height))
	gift.New().Draw(dst, s.background)

	for _, d := range s.drawObjs {
		if d.status.GetStatus() == components.StatusActive {
			if tickCnt%2 == 0 {
				d.sprite.GetFilter().DrawAt(dst, d.sprite.GetSrc(), image.Pt(d.position.X(), d.position.Y()), gift.OverOperator)
			} else {
				d.sprite.GetFilterA().DrawAt(dst, d.sprite.GetSrc(), image.Pt(d.position.X(), d.position.Y()), gift.OverOperator)
			}
		} else if d.status.GetStatus() == components.StatusExplode {
			d.sprite.GetFilterE().DrawAt(dst, d.sprite.GetSrc(), image.Pt(d.position.X(), d.position.Y()), gift.OverOperator)
		}
	}

	globals.PrintImage(dst)
}

func (s *drawSystem) Register(id uint64, c ecs.Component) {
	if c.GetType() == components.CompTypeSprite {
		drawer := &drawer{}
		drawer.sprite = c.(components.Sprite)
		drawer.position = s.world.GetComponent(id, components.CompTypePosition).(components.Position)
		drawer.status = s.world.GetComponent(id, components.CompTypeStatus).(components.Status)
		s.drawObjs[id] = drawer
	}
}

func (s *drawSystem) Unregister(id uint64, componentType int) {
	if componentType == components.CompTypeSprite {
		delete(s.drawObjs, id)
	}
}
