package main

import "github.com/veandco/go-sdl2/sdl"

type Player struct {
	pos Pos
}

func (p Player) draw(s *sdl.Surface, w *World) {
	rect := sdl.Rect{X: int32(p.pos.x), Y: int32(p.pos.y), H: 20, W: 20}
	s.FillRect(&rect, 0xffff0000)
}
