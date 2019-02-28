package main

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyboard(keycode sdl.Keycode, isRunning *bool) {
	switch keycode {
	case sdl.K_ESCAPE:
		*isRunning = false
	}
}

func startGame(resp chan<- string, wg *sync.WaitGroup, isRunning *bool) {
	defer wg.Done()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, sdl.MapRGB(surface.Format, 0, 10, 20))
	window.UpdateSurface()

	//*isRunning = true
	for *isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				*isRunning = false
				break
			case *sdl.KeyboardEvent:
				handleKeyboard(sdl.GetKeyFromScancode(t.Keysym.Scancode), isRunning)
			}
		}
		surface.FillRect(nil, sdl.MapRGB(surface.Format, 0, 10, 20))
	}

	resp <- "Game"
}
