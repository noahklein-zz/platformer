package main

import "github.com/veandco/go-sdl2/sdl"
import "math"

type World struct {
	height    int
	width     int
	player    Player
	isRunning bool
}

const FRICTION float64 = 0.9
const GRAVITY float64 = 2.5

func (w World) allEntities() []Entity {
	return []Entity{
		w.player,
	}
}

type Entity interface {
	draw(s *sdl.Surface, w *World)
}

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
		draw(surface, world)
		window.UpdateSurface()
		for input := range inputs() {
			world = update(world, input)
		}
		sdl.Delay(uint32(1000 / config.framerate))
	}
	sdl.Quit()
}

func draw(s *sdl.Surface, w *World) {
	// clear screen
	s.FillRect(nil, 0x000000)
	for _, ent := range w.allEntities() {
		ent.draw(s, w)
	}
}

func update(w *World, input Input) *World {
	pos := &w.player.pos
	switch input {
	case UP:
		if int(pos.y) == (*w).height-100 {
			pos.vy = -30
		}
	case RIGHT:
		pos.vx = 10
	case LEFT:
		pos.vx = -10
	case QUIT:
		w.isRunning = false
	}
	pos.x += pos.vx
	pos.y += pos.vy
	pos.y = math.Min(float64((*w).height)-100, pos.y)
	pos.vx *= FRICTION
	pos.vy += GRAVITY

	return w
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
		player:    Player{pos: Pos{3, 3, 0, 0}},
		isRunning: true,
	}

	return window, surface, initialWorld
}
