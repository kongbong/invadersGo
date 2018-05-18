package globals

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

var GSceneManager SceneManager
var GInputManager InputManager
var GDispatcher Dispatcher
var GDriver gxui.Driver
var GWindow gxui.Window
var GWindowImg gxui.Image
var Width int
var Height int

func Init(d gxui.Driver, width, height int, title string) {
	Width = width
	Height = height
	GDriver = d
	theme := dark.CreateTheme(d)
	GWindow = theme.CreateWindow(width, height, title)
	GWindowImg = theme.CreateImage()
	GWindow.AddChild(GWindowImg)

	GSceneManager = &implSceneManager{}

	i := &implInputManager{}
	i.downedKey = make(map[gxui.KeyboardKey]bool)
	i.downedMod = make(map[gxui.KeyboardModifier]bool)
	i.upEvents = make(map[gxui.KeyboardKey]bool)
	i.upModEvents = make(map[gxui.KeyboardModifier]bool)

	GWindow.OnKeyDown(i.OnKeyDown)
	GWindow.OnKeyUp(i.OnKeyUp)
	GInputManager = i

	GDispatcher = &implDispatcher{tasks: make([]func(), 0, 10)}

	GWindow.OnClose(GDriver.Terminate)
}

func Tick(tickCnt uint64) {
	GSceneManager.Tick(tickCnt)
	GInputManager.Tick()
	GDispatcher.Tick()
}
