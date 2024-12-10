package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BallCollider interface {
	Collides(ball Ball) bool
}

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
	p.Y -= paddleSpeed
}

func (p *Paddle) MoveDown() {
	p.Y += paddleSpeed
}

func (p Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.X, p.Y, p.Width, p.Height, color.RGBA{255, 255, 255, 255}, false)
}

func (p Paddle) Collides(ball Ball) bool {
	return ball.Position.X+ball.Radius > p.X && ball.Position.X-ball.Radius < p.X+p.Width && ball.Position.Y+ball.Radius > p.Y && ball.Position.Y-ball.Radius < p.Y+p.Height
}

type CpuPaddle struct {
	Paddle
}

func NewCpuPaddle(x, y, width, height float32) *CpuPaddle {
	return &CpuPaddle{
		Paddle: Paddle{
			Object: Object{x, y},
			Width:  width,
			Height: height,
		},
	}
}

func (c *CpuPaddle) Move(g *Game) {
	if g.ball.Position.Y < c.Y+c.Height/2 && !(c.Y < 0) {
		c.MoveUp()
	}

	if g.ball.Position.Y > c.Y+c.Height/2 && !(c.Y+c.Height > screenHeight) {
		c.MoveDown()
	}
}

func (c *CpuPaddle) MoveUp() {
	c.Y -= paddleSpeed
}

func (c *CpuPaddle) MoveDown() {
	c.Y += paddleSpeed
}
