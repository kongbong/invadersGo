package ecs

type World interface {
	SpawnEntity() Entity
	GetEntity(id uint64) Entity
	RemoveEntity(id uint64)
	AddComponent(id uint64, c Component)
	RemoveComponent(id uint64, compType int)
	GetComponent(id uint64, compType int) Component
	ForeachComponents(compType int, cb func(Component))
	AddSystem(s System)
	Subscribe(compType int, s System)
	Tick(tickCnt uint64)
}

func NewWorld(d Dispatcher) World {
	w := &implWorld{}
	w.entities = make(map[uint64]Entity)
	w.components = make(map[uint64]map[int]Component)
	w.componentsForType = make(map[int]map[uint64]Component)
	w.systems = make([]System, 0)
	w.subscribers = make(map[int][]System)
	w.disptcher = d
	return w
}

type implWorld struct {
	entities          map[uint64]Entity
	components        map[uint64]map[int]Component
	componentsForType map[int]map[uint64]Component
	systems           []System
	subscribers       map[int][]System
	disptcher         Dispatcher
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

		compType := t
		w.disptcher.Dispatch(func() {
			w.onRemoveComponent(id, compType)
		})
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

	w.disptcher.Dispatch(func() {
		w.onAddComponent(id, c)
	})
}

func (w *implWorld) onAddComponent(id uint64, c Component) {
	for _, s := range w.subscribers[c.GetType()] {
		s.Register(id, c)
	}
}

func (w *implWorld) RemoveComponent(id uint64, compType int) {
	if w.components[id] == nil {
		return
	}
	delete(w.components[id], compType)
	delete(w.componentsForType[compType], id)

	w.disptcher.Dispatch(func() {
		w.onRemoveComponent(id, compType)
	})
}

func (w *implWorld) onRemoveComponent(id uint64, compType int) {
	for _, s := range w.subscribers[compType] {
		s.Unregister(id, compType)
	}
}

func (w *implWorld) GetComponent(id uint64, compType int) Component {
	if w.components[id] == nil {
		return nil
	}

	return w.components[id][compType]
}

func (w *implWorld) ForeachComponents(compType int, cb func(Component)) {
	if w.componentsForType[compType] == nil {
		return
	}

	for _, c := range w.componentsForType[compType] {
		cb(c)
	}
}

func (w *implWorld) AddSystem(s System) {
	w.systems = append(w.systems, s)
	w.disptcher.Dispatch(func() {
		s.Init(w)
	})
}

func (w *implWorld) Subscribe(compType int, s System) {
	if w.subscribers[compType] == nil {
		w.subscribers[compType] = make([]System, 0, 1)
	}
	w.subscribers[compType] = append(w.subscribers[compType], s)
}

func (w *implWorld) Tick(tickCnt uint64) {
	for _, s := range w.systems {
		s.Tick(tickCnt)
	}
}
