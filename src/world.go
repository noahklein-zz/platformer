package main

import "math"

type World struct {
	height    int
	width     int
	player    Player
	isRunning bool
	inputs    map[Input]bool
}

func (w World) allEntities() []Entity {
	return []Entity{
		w.player,
	}
}

func handleInput(w *World, input Input, isDown bool) *World {
	pos := &w.player.pos
	switch input {
	case UP:
		if !isDown {
			w.player.canJump = true
		}
		if int(pos.y) == (*w).height-100 && w.player.canJump && isDown {
			w.player.canJump = false
			pos.vy = -30
		}
	case RIGHT:
		if isDown {
			pos.vx = 10
		}
	case LEFT:
		if isDown {
			pos.vx = -10
		}
	case QUIT:
		if isDown {
			w.isRunning = false
		}
	}
	return w
}

func update(w *World) *World {
	pos := &w.player.pos
	pos.x += pos.vx
	pos.y += pos.vy
	pos.y = math.Min(float64((*w).height)-100, pos.y)
	pos.vx *= FRICTION
	pos.vy += GRAVITY
	return w
}
