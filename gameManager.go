package main

import (
	"fmt"
	"image"

	"github.com/kongbong/invadersGo/components"
	"github.com/kongbong/invadersGo/ecs"
)

type actor struct {
	position  components.Position
	status    components.Status
	collision components.Collision
}

type invader struct {
	actor
	points int
}

type gameManager struct {
	invaders map[uint64]*invader
	bombs    map[uint64]*actor
	cannon   *actor
	cannonId uint64
	beams    map[uint64]*actor
	world    ecs.World
}

func NewGameManager() ecs.System {
	gm := &gameManager{}
	gm.invaders = make(map[uint64]*invader)
	gm.bombs = make(map[uint64]*actor)
	gm.beams = make(map[uint64]*actor)
	return gm
}

func NewInvaderActor(a *actor, points int) *invader {
	return &invader{actor{a.position, a.status, a.collision}, points}
}

func (gm *gameManager) Init(w ecs.World) {
	fmt.Println("gameManager Init")
	gm.world = w
	w.Subscribe(components.CompTypeInvader, gm)
	w.Subscribe(components.CompTypeBeam, gm)
	w.Subscribe(components.CompTypeBomb, gm)
	w.Subscribe(components.CompTypeCannon, gm)
}

func (gm *gameManager) Tick(tickCnt uint64) {

	if gm.cannon == nil || gm.cannon.status.GetStatus() != components.StatusActive {
		return
	}

	for id, bomb := range gm.bombs {
		if checkCollision(bomb.position, bomb.collision, gm.cannon.position, gm.cannon.collision) {
			// explode
			fmt.Println("explode cannon")
			gm.cannon.status.SetStatus(components.StatusExplode)
			gm.world.RemoveEntity(id)

			dispatcher.Dispatchn(TickCnt+2, func() {
				gm.world.RemoveEntity(gm.cannonId)
			})
			break
		}
	}

Loop:
	for id, beam := range gm.beams {
		for id2, invader := range gm.invaders {
			if invader.status.GetStatus() == components.StatusActive && checkCollision(beam.position, beam.collision, invader.position, invader.collision) {
				// explode
				fmt.Println("explode invader")
				invader.status.SetStatus(components.StatusExplode)
				gm.world.RemoveEntity(id)
				Score += invader.points

				invaderId := id2
				dispatcher.Dispatchn(TickCnt+2, func() {
					gm.world.RemoveEntity(invaderId)
				})
				break Loop
			}
		}
	}
}

func checkCollision(p1 components.Position, c1 components.Collision, p2 components.Position, c2 components.Collision) bool {
	spriteA := image.Rect(p1.X(), p1.Y(), p1.X()+c1.GetWidth(), p1.Y()+c1.GetHeight())
	spriteB := image.Rect(p2.X(), p2.Y(), p2.X()+c2.GetWidth(), p2.Y()+c2.GetHeight())
	if spriteA.Min.X < spriteB.Max.X && spriteA.Max.X > spriteB.Min.X &&
		spriteA.Min.Y < spriteB.Max.Y && spriteA.Max.Y > spriteB.Min.Y {
		return true
	}
	return false
}

func (gm *gameManager) Register(id uint64, c ecs.Component) {
	a := &actor{}
	a.collision = gm.world.GetComponent(id, components.CompTypeCollision).(components.Collision)
	a.status = gm.world.GetComponent(id, components.CompTypeStatus).(components.Status)
	a.position = gm.world.GetComponent(id, components.CompTypePosition).(components.Position)

	switch c.GetType() {
	case components.CompTypeInvader:
		i := c.(components.Invader)
		gm.invaders[id] = NewInvaderActor(a, i.GetScore())
	case components.CompTypeBeam:
		gm.beams[id] = a
	case components.CompTypeBomb:
		gm.bombs[id] = a
	case components.CompTypeCannon:
		fmt.Printf("register cannon %P\n", c)
		gm.cannon = a
		gm.cannonId = id
	}
}

func (gm *gameManager) Unregister(id uint64, componentType int) {
	switch componentType {
	case components.CompTypeInvader:
		delete(gm.invaders, id)
		if len(gm.invaders) == 0 {
			dispatcher.Dispatchn(TickCnt+2, func() {
				// change to result
				sceneManager.ChangeScene(NewResultScene())
			})
		}
	case components.CompTypeBeam:
		delete(gm.beams, id)
	case components.CompTypeBomb:
		delete(gm.bombs, id)
	case components.CompTypeCannon:
		gm.cannon = nil
		gm.cannonId = 0
		dispatcher.Dispatchn(TickCnt+2, func() {
			// change to result
			sceneManager.ChangeScene(NewResultScene())
		})
	}
}
