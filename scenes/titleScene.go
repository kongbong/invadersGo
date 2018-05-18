package scenes

import (
	"invadersGo/ecs"
	"invadersGo/globals"

	"github.com/google/gxui"
)

func NewTitleScene() ecs.Scene {
	return &TitleScene{}
}

type TitleScene struct {
}

func (t *TitleScene) Init() {
	startScreen := globals.GetImage("imgs/start.png")
	globals.PrintImage(startScreen)
}

func (t *TitleScene) OnDestroy() {

}

func (t *TitleScene) Tick(tickCnt uint64) {
	if globals.GInputManager.IsKeyUp(gxui.KeyQ, gxui.ModNone) {
		globals.GDriver.Call(func() {
			globals.GWindow.Close()
		})
	} else if globals.GInputManager.IsKeyUp(gxui.KeyS, gxui.ModNone) {
		globals.GDispatcher.Dispatch(func() {
			globals.GSceneManager.ChangeScene(NewGameScene())
		})
	}
}
