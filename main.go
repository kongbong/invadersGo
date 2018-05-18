package main

import (
	"invaders2/globals"
	"invaders2/scenes"
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
	globals.Init(d, Width, Height, Title)

	globals.GSceneManager.ChangeScene(&scenes.TitleScene{})

	ticker := time.NewTicker(time.Millisecond * TickInterval)
	var tickCnt uint64
	go func() {

		for _ = range ticker.C {
			tickCnt++
			minLoop(tickCnt)
		}
	}()

	globals.GWindow.OnClose(ticker.Stop)
}

func minLoop(tickCnt uint64) {
	globals.Tick(tickCnt)
}

func main() {
	gl.StartDriver(appMain)
}
