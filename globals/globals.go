package globals

import (
	"invadersGo/ecs"

	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

var GSceneManager ecs.SceneManager
var GInputManager ecs.InputManager
var GDispatcher ecs.Dispatcher
var GDriver gxui.Driver
var GWindow gxui.Window
var GWindowImg gxui.Image
var Width int
var Height int
var Score int
var TickCnt uint64

func Init(d gxui.Driver, width, height int, title string) {
	Width = width
	Height = height
	GDriver = d
	theme := dark.CreateTheme(d)
	GWindow = theme.CreateWindow(width, height, title)
	GWindowImg = theme.CreateImage()
	GWindow.AddChild(GWindowImg)

	GSceneManager = ecs.NewSceneManager()
	GInputManager = ecs.NewInputManager(GWindow)
	GDispatcher = ecs.NewDispatcher()

	GWindow.OnClose(GDriver.Terminate)
}

func Tick() {
	TickCnt++
	GSceneManager.Tick(TickCnt)
	GInputManager.Tick()
	GDispatcher.Tick(TickCnt)
}
