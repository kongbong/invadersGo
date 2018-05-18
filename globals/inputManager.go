package globals

import (
	"github.com/google/gxui"
)

type InputManager interface {
	Tick()
	IsKeyDown(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool
	IsKeyUp(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool
}

type implInputManager struct {
	downedKey   map[gxui.KeyboardKey]bool
	downedMod   map[gxui.KeyboardModifier]bool
	upEvents    map[gxui.KeyboardKey]bool
	upModEvents map[gxui.KeyboardModifier]bool
}

func (i *implInputManager) Tick() {
	// clear upevents
	for key, value := range i.upEvents {
		if value == true {
			i.downedKey[key] = false
		}
	}
	for key, value := range i.upModEvents {
		if value == true {
			i.downedMod[key] = false
		}
	}

	i.upEvents = make(map[gxui.KeyboardKey]bool)
	i.upModEvents = make(map[gxui.KeyboardModifier]bool)
}

func (i *implInputManager) IsKeyDown(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool {
	return i.downedKey[key] && i.downedMod[mod]
}

func (i *implInputManager) IsKeyUp(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool {
	return i.upEvents[key] && i.upModEvents[mod]
}

func (i *implInputManager) OnKeyDown(e gxui.KeyboardEvent) {
	i.downedKey[e.Key] = true
	i.downedMod[e.Modifier] = true
}

func (i *implInputManager) OnKeyUp(e gxui.KeyboardEvent) {
	i.upEvents[e.Key] = true
	i.upModEvents[e.Modifier] = true
}
