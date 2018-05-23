package ecs

import (
	"github.com/google/gxui"
)

type InputManager interface {
	Tick()
	IsKeyDown(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool
	IsKeyUp(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool
	PressedAnyKey() bool
}

func NewInputManager(window gxui.Window) InputManager {
	i := &implInputManager{}
	i.downedKey = make(map[gxui.KeyboardKey]int)
	i.downedMod = make(map[gxui.KeyboardModifier]int)
	i.upEvents = make(map[gxui.KeyboardKey]int)
	i.upModEvents = make(map[gxui.KeyboardModifier]int)

	window.OnKeyDown(i.OnKeyDown)
	window.OnKeyUp(i.OnKeyUp)
	return i
}

type implInputManager struct {
	downedKey   map[gxui.KeyboardKey]int
	downedMod   map[gxui.KeyboardModifier]int
	upEvents    map[gxui.KeyboardKey]int
	upModEvents map[gxui.KeyboardModifier]int
}

func (i *implInputManager) Tick() {
	// clear upevents
	for key, value := range i.upEvents {
		if value > 0 {
			i.downedKey[key] -= value
		}
	}
	for key, value := range i.upModEvents {
		if value > 0 {
			i.downedMod[key] -= value
		}
	}

	i.upEvents = make(map[gxui.KeyboardKey]int)
	i.upModEvents = make(map[gxui.KeyboardModifier]int)
}

func (i *implInputManager) IsKeyDown(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool {
	return i.downedKey[key] > 0 && i.downedMod[mod] > 0
}

func (i *implInputManager) IsKeyUp(key gxui.KeyboardKey, mod gxui.KeyboardModifier) bool {
	return i.upEvents[key] > 0 && i.upModEvents[mod] > 0
}

func (i *implInputManager) PressedAnyKey() bool {
	for _, v := range i.downedKey {
		if v > 0 {
			return true
		}
	}
	return false
}

func (i *implInputManager) OnKeyDown(e gxui.KeyboardEvent) {
	i.downedKey[e.Key]++
	i.downedMod[e.Modifier]++
}

func (i *implInputManager) OnKeyUp(e gxui.KeyboardEvent) {
	i.upEvents[e.Key]++
	i.upModEvents[e.Modifier]++
}
