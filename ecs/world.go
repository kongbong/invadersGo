package ecs

type World interface {
	SpawnEntity() Entity
	GetEntity(id uint64) Entity
	RemoveEntity(id uint64)
	AddComponent(id uint64, c Component)
	RemoveComponent(id uint64, componentType int)
	GetComponent(id uint64, componentType int) Component
	ForeachComponents(componentType int, cb func(Component))
	AddSystem(s System)
	Subscribe(componentType int, s System)
	Tick(tickCnt uint64)
}

func NewWorld() World {
	w := &implWorld{}
	w.entities = make(map[uint64]Entity)
	w.components = make(map[uint64]map[int]Component)
	w.componentsForType = make(map[int]map[uint64]Component)
	w.systems = make([]System, 0)
	w.subscribers = make(map[int][]System)
	return w
}

type implWorld struct {
	entities          map[uint64]Entity
	components        map[uint64]map[int]Component
	componentsForType map[int]map[uint64]Component
	systems           []System
	subscribers       map[int][]System
}

func (w *implWorld) SpawnEntity() Entity {
	e := NewEntity()
	w.entities[e.Id()] = e
	return e
}

func (w *implWorld) GetEntity(id uint64) Entity {
	return w.entities[id]
}

func (w *implWorld) RemoveEntity(id uint64) {
	delete(w.entities, id)
	for t := range w.components[id] {
		delete(w.componentsForType[t], id)
		for _, s := range w.subscribers[t] {
			s.Unregister(id, t)
		}
	}
	delete(w.components, id)
}

func (w *implWorld) AddComponent(id uint64, c Component) {
	if w.components[id] == nil {
		w.components[id] = make(map[int]Component)
	}
	t := c.GetType()
	w.components[id][t] = c
	if w.componentsForType[t] == nil {
		w.componentsForType[t] = make(map[uint64]Component)
	}
	w.componentsForType[t][id] = c

	for _, s := range w.subscribers[t] {
		s.Register(id, c)
	}
}

func (w *implWorld) RemoveComponent(id uint64, componentType int) {
	if w.components[id] == nil {
		return
	}
	delete(w.components[id], componentType)
	delete(w.componentsForType[componentType], id)
	for _, s := range w.subscribers[componentType] {
		s.Unregister(id, componentType)
	}
}

func (w *implWorld) GetComponent(id uint64, componentType int) Component {
	if w.components[id] == nil {
		return nil
	}

	return w.components[id][componentType]
}

func (w *implWorld) ForeachComponents(componentType int, cb func(Component)) {
	if w.componentsForType[componentType] == nil {
		return
	}

	for _, c := range w.componentsForType[componentType] {
		cb(c)
	}
}

func (w *implWorld) AddSystem(s System) {
	w.systems = append(w.systems, s)
	s.Init(w)
}

func (w *implWorld) Subscribe(componentType int, s System) {
	if w.subscribers[componentType] == nil {
		w.subscribers[componentType] = make([]System, 0, 1)
	}
	w.subscribers[componentType] = append(w.subscribers[componentType], s)
}

func (w *implWorld) Tick(tickCnt uint64) {
	for _, s := range w.systems {
		s.Tick(tickCnt)
	}
}
