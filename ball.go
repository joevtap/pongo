package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Vector2D struct {
	X float32
	Y float32
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

type Ball struct {
	Radius   float32
	Position Vector2D
	Speed    Vector2D
}

func NewBall(speed, position Vector2D, radius float32) *Ball {
	return &Ball{radius, position, speed}
}

func (b *Ball) Move(game *Game) {
	b.Position = b.Position.Add(b.Speed)

	if b.Position.X+b.Radius > 320 {
		game.cpuScore++
		game.Reset()
	}

	if b.Position.X-b.Radius < 0 {
		game.playerScore++
		game.Reset()
	}

	if b.Position.Y-b.Radius < 0 || b.Position.Y+b.Radius > 240 {
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
	return b.Position.X+b.Radius > paddle.X && b.Position.X-b.Radius < paddle.X+paddle.Width && b.Position.Y+b.Radius > paddle.Y && b.Position.Y-b.Radius < paddle.Y+paddle.Height
}

func (b Ball) Collides(ballCollider BallCollider) bool {
	return ballCollider.Collides(b)
}

func (b *Ball) ReverseX() {
	b.Speed.X = -b.Speed.X
}

func (b *Ball) ReverseY() {
	b.Speed.Y = -b.Speed.Y
}

func (b Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.Position.X, b.Position.Y, b.Radius, color.RGBA{255, 255, 255, 255}, false)
}
