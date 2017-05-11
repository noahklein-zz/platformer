package main

import "github.com/veandco/go-sdl2/sdl"

// FRICTION is the player's friction.
const FRICTION float64 = 0.9

// GRAVITY is the gravity in the world.
const GRAVITY float64 = 2.5

// Entity is a game object.
type Entity interface {
	draw(s *sdl.Surface, w *World)
}

// Pos represents position and velocity
type Pos struct {
	x  float64
	y  float64
	vx float64
	vy float64
}

func main() {
	config := getConfig()
	window, surface, world := initialize(config)
	defer window.Destroy()

	for world.isRunning {
		draw(window, surface, world)
		world.inputs = inputs(world.inputs)
		for input, isDown := range world.inputs {
			world = handleInput(world, input, isDown)
		}
		world = update(world)
		sdl.Delay(uint32(1000 / config.framerate))
	}
	sdl.Quit()
}

func draw(win *sdl.Window, s *sdl.Surface, w *World) {
	// clear screen
	s.FillRect(nil, 0x000000)
	for _, ent := range w.allEntities() {
		ent.draw(s, w)
	}
	win.UpdateSurface()
}

func initialize(config Config) (*sdl.Window, *sdl.Surface, *World) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		config.width, config.height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	initialWorld := &World{
		height:    config.height,
		width:     config.width,
		player:    Player{pos: Pos{3, 3, 0, 0}, canJump: true},
		isRunning: true,
		inputs: map[Input]bool{
			QUIT:  false,
			UP:    false,
			DOWN:  false,
			LEFT:  false,
			RIGHT: false,
			NONE:  false,
		},
	}

	return window, surface, initialWorld
}
