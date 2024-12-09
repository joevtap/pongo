package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	Object
	Width  float32
	Height float32
}

func NewPaddle(x, y, width, height float32) *Paddle {
	return &Paddle{
		Object: Object{x, y},
		Width:  width,
		Height: height,
	}
}

func (p *Paddle) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && !(p.Y < 0) {
		p.MoveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) && !(p.Y+p.Height > screenHeight) {
		p.MoveDown()
	}
}

func (p *Paddle) MoveUp() {
	p.Y -= 4
}

func (p *Paddle) MoveDown() {
	p.Y += 4
}

func (p Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.X, p.Y, p.Width, p.Height, color.RGBA{255, 255, 255, 255}, false)
}
