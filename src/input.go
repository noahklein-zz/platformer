package main

import "github.com/veandco/go-sdl2/sdl"

// Input is an enum type representing user input.
type Input int

const (
	UP Input = iota
	DOWN
	LEFT
	RIGHT
	QUIT
	NONE
)

func inputs() map[Input]bool {
	m := make(map[Input]bool)
	for {
		key, ok := sdl.PollEvent().(*sdl.KeyDownEvent)
		if !ok {
			m[NONE] = true
			return m
		}
		switch key.Keysym.Scancode {
		case sdl.SCANCODE_ESCAPE:
			m[QUIT] = true
			return m
		case sdl.SCANCODE_W:
			m[UP] = true
		case sdl.SCANCODE_S:
			m[DOWN] = true
		case sdl.SCANCODE_A:
			m[LEFT] = true
		case sdl.SCANCODE_D:
			m[RIGHT] = true
		default:
			m[NONE] = true
		}
	}
}
