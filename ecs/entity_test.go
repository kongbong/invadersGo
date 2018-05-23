package ecs_test

import (
	"testing"

	"github.com/kongbong/invadersGo/ecs"
)

const TestComponentType = 1

type TestComponent struct {
}

func (t *TestComponent) GetType() int {
	return TestComponentType
}

func TestWorld(t *testing.T) {
	world := ecs.NewWorld()
	entity := world.SpawnEntity()

	if entity == nil {
		t.Error("Entity should be not nil")
	}

	id := entity.Id()
	e := world.GetEntity(id)
	if e != entity {
		t.Error("GetEntity should return same entity")
	}

	component := &TestComponent{}
	world.AddComponent(id, component)
	c := world.GetComponent(id, TestComponentType)

	if c == nil {
		t.Error("GetComponent should return not nil")
	}

	called := false
	world.ForeachComponents(1, func(comp ecs.Component) {
		if c != comp {
			t.Error("c and comp should be equal")
		}
		called = true
	})

	if !called {
		t.Error("Foreach should be called")
	}

	world.RemoveComponent(id, TestComponentType)
	c = world.GetComponent(id, TestComponentType)

	if c != nil {
		t.Error("c should be removed")
	}

	world.ForeachComponents(TestComponentType, func(comp ecs.Component) {
		t.Error("this cb should be not called")
	})

	world.AddComponent(id, component)
	world.RemoveEntity(id)

	e = world.GetEntity(id)
	if e != nil {
		t.Error("e should be null")
	}

	c = world.GetComponent(id, TestComponentType)
	if c != nil {
		t.Error("c should be nil")
	}

	world.ForeachComponents(TestComponentType, func(comp ecs.Component) {
		t.Error("this cb should be not called")
	})
}

type TestSystem struct {
	initCalled       bool
	tickCalled       bool
	registerCalled   bool
	compType         int
	unregisterCalled bool
	unregiCompType   int
}

func (s *TestSystem) Init(w ecs.World) {
	s.initCalled = true
}

func (s *TestSystem) Tick(tickCnt uint64) {
	s.tickCalled = true
}

func (s *TestSystem) Register(id uint64, c ecs.Component) {
	s.registerCalled = true
	s.compType = c.GetType()
}

func (s *TestSystem) Unregister(id uint64, componentType int) {
	s.unregisterCalled = true
	s.unregiCompType = componentType
}

func TestSystems(t *testing.T) {
	world := ecs.NewWorld()
	s := &TestSystem{}
	world.AddSystem(s)

	if !s.initCalled {
		t.Error("system Init should be called")
	}

	world.Tick(1)

	if !s.tickCalled {
		t.Error("system Tick should be called")
	}

	world.Subscribe(TestComponentType, s)

	entity := world.SpawnEntity()
	id := entity.Id()
	component := &TestComponent{}
	world.AddComponent(id, component)

	if !s.registerCalled {
		t.Error("register should be called")
	}
	if s.compType != TestComponentType {
		t.Error("register component type should be same")
	}

	world.RemoveComponent(id, TestComponentType)

	if !s.unregisterCalled {
		t.Error("unregister should be called")
	}
	if s.unregiCompType != TestComponentType {
		t.Error("unregister component type should be same")
	}

	s.registerCalled = false
	s.compType = 0
	s.unregisterCalled = false
	s.unregiCompType = 0

	world.AddComponent(id, component)
	if !s.registerCalled {
		t.Error("register should be called")
	}
	if s.compType != TestComponentType {
		t.Error("register component type should be same")
	}

	world.RemoveEntity(id)
	if !s.unregisterCalled {
		t.Error("unregister should be called")
	}
	if s.unregiCompType != TestComponentType {
		t.Error("unregister component type should be same")
	}
}
