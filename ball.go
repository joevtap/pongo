package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	Object
	Radius float32
	SpeedX float32
	SpeedY float32
}

func NewBall(x, y, radius, speedX, speedY float32) *Ball {
	return &Ball{
		Object: Object{x, y},
		Radius: radius,
		SpeedX: speedX,
		SpeedY: speedY,
	}
}

func (b *Ball) Move(game *Game) {
	b.X += b.SpeedX
	b.Y += b.SpeedY

	if b.X+b.Radius > 320 {
		game.Reset()
	}

	if b.X-b.Radius < 0 {
		b.ReverseX()
	}

	if b.Y-b.Radius < 0 || b.Y+b.Radius > 240 {
		b.ReverseY()
	}

	if b.Collides(game.player) {
		game.score++
		b.ReverseX()
	}
}

func (b *Ball) Collides(paddle Paddle) bool {
	return b.X+b.Radius > paddle.X && b.X-b.Radius < paddle.X+paddle.Width && b.Y+b.Radius > paddle.Y && b.Y-b.Radius < paddle.Y+paddle.Height
}

func (b *Ball) ReverseX() {
	b.SpeedX = -b.SpeedX
}

func (b *Ball) ReverseY() {
	b.SpeedY = -b.SpeedY
}

func (b Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.X, b.Y, b.Radius, color.RGBA{255, 255, 255, 255}, true)
}
