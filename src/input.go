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

func inputs(m map[Input]bool) map[Input]bool {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch t := e.(type) {
		case *sdl.QuitEvent:
			m[QUIT] = true
		case *sdl.KeyDownEvent:
			input := scancodeToInput(t.Keysym.Scancode)
			m[input] = true
		case *sdl.KeyUpEvent:
			input := scancodeToInput(t.Keysym.Scancode)
			m[input] = false
		}
	}
	return m
}

func scancodeToInput(scancode sdl.Scancode) Input {
	switch scancode {
	case sdl.SCANCODE_ESCAPE:
		return QUIT
	case sdl.SCANCODE_W:
		return UP
	case sdl.SCANCODE_S:
		return DOWN
	case sdl.SCANCODE_A:
		return LEFT
	case sdl.SCANCODE_D:
		return RIGHT
	default:
		return NONE
	}
}
