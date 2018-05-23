package main

import (
	"invadersGo/ecs"

	"github.com/google/gxui"
)

func NewTitleScene() ecs.Scene {
	return &TitleScene{}
}

type TitleScene struct {
}

func (t *TitleScene) Init() {
	startScreen := GetImage("imgs/start.png")
	PrintImage(startScreen)
}

func (t *TitleScene) OnDestroy() {

}

func (t *TitleScene) Tick(tickCnt uint64) {
	if inputManager.IsKeyUp(gxui.KeyQ, gxui.ModNone) {
		driver.Call(func() {
			window.Close()
		})
	} else if inputManager.IsKeyUp(gxui.KeyS, gxui.ModNone) {
		dispatcher.Dispatch(func() {
			sceneManager.ChangeScene(NewGameScene())
		})
	}
}
