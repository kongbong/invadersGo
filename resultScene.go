package main

import (
	"fmt"
	"image"
	"image/color"
	"invadersGo/ecs"

	"github.com/disintegration/gift"
)

func NewResultScene() ecs.Scene {
	return &ResultScene{}
}

type ResultScene struct {
}

func (t *ResultScene) Init() {
	background := GetImage("imgs/bg.png")
	dst := image.NewRGBA(image.Rect(0, 0, Width, Height))
	gift.New().Draw(dst, background)

	addLabel(dst, 50, 50, "GAME OVER", color.RGBA{255, 255, 255, 255})
	str := fmt.Sprintf("SCORE: %d", Score)
	addLabel(dst, 50, 70, str, color.RGBA{255, 255, 255, 255})
	PrintImage(dst)
}

func (t *ResultScene) OnDestroy() {

}

func (t *ResultScene) Tick(tickCnt uint64) {
	if inputManager.PressedAnyKey() {
		dispatcher.Dispatch(func() {
			sceneManager.ChangeScene(NewTitleScene())
		})
	}
}
