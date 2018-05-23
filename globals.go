package main

import (
	"github.com/kongbong/invadersGo/ecs"

	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

var sceneManager ecs.SceneManager
var inputManager ecs.InputManager
var dispatcher ecs.Dispatcher
var driver gxui.Driver
var window gxui.Window
var windowImg gxui.Image
var Score int
var TickCnt uint64

func InitWindow(d gxui.Driver, width, height int, title string) {
	driver = d
	theme := dark.CreateTheme(d)
	window = theme.CreateWindow(width, height, title)
	windowImg = theme.CreateImage()
	window.AddChild(windowImg)

	sceneManager = ecs.NewSceneManager()
	inputManager = ecs.NewInputManager(window)
	dispatcher = ecs.NewDispatcher()

	window.OnClose(driver.Terminate)
}

func Tick() {
	TickCnt++
	sceneManager.Tick(TickCnt)
	inputManager.Tick()
	dispatcher.Tick(TickCnt)
}
