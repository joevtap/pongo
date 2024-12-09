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
		game.cpuScore++
		game.Reset()
	}

	if b.X-b.Radius < 0 {
		game.playerScore++
		game.Reset()
	}

	if b.Y-b.Radius < 0 || b.Y+b.Radius > 240 {
		b.ReverseY()
	}

	if b.Collides(game.player) {
		b.ReverseX()
	}

	if b.Collides(game.cpu) {
		b.ReverseX()
	}
}

func (b *Ball) CollidesWithCpu(paddle CpuPaddle) bool {
	return b.X+b.Radius > paddle.X && b.X-b.Radius < paddle.X+paddle.Width && b.Y+b.Radius > paddle.Y && b.Y-b.Radius < paddle.Y+paddle.Height
}

func (b Ball) Collides(ballCollider BallCollider) bool {
	return ballCollider.Collides(b)
}

func (b *Ball) ReverseX() {
	b.SpeedX = -b.SpeedX
}

func (b *Ball) ReverseY() {
	b.SpeedY = -b.SpeedY
}

func (b Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.X, b.Y, b.Radius, color.RGBA{255, 255, 255, 255}, false)
}
