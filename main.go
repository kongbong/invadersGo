package main

import (
	"math/rand"
	"time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
)

const Width = 400
const Height = 300
const Title = "Space Invaders"
const TickInterval = 500

func appMain(d gxui.Driver) {
	rand.Seed(time.Now().UTC().UnixNano())
	InitWindow(d, Width, Height, Title)

	sceneManager.ChangeScene(NewTitleScene())

	ticker := time.NewTicker(time.Millisecond * TickInterval)
	go func() {

		for _ = range ticker.C {
			mainLoop()
		}
	}()

	window.OnClose(ticker.Stop)
}

func mainLoop() {
	Tick()
}

func main() {
	gl.StartDriver(appMain)
}
