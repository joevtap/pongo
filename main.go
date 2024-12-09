package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"

	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 320
	screenHeight = 240
	ballSpeed    = 3
)

// Game implements ebiten.Game interface.
type Game struct {
	ball   Ball
	player Paddle
	score  int
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PonGo - A pong game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game.Init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Init() {
	g.RandomBall()
	g.player = *NewPaddle(screenWidth-10, screenHeight/2, 5, 40)
	g.score = 0
}

func (g *Game) Reset() {
	g.RandomBall()
	g.score = 0
}

func (g *Game) RandomBall() {
	r := rand.Intn(4)

	if r == 0 {
		g.ball = *NewBall(screenWidth/2, screenHeight/2, 5, ballSpeed, ballSpeed)
	} else if r == 1 {
		g.ball = *NewBall(screenWidth/2, screenHeight/2, 5, -ballSpeed, -ballSpeed)
	} else if r == 2 {
		g.ball = *NewBall(screenWidth/2, screenHeight/2, 5, ballSpeed, -ballSpeed)
	} else {
		g.ball = *NewBall(screenWidth/2, screenHeight/2, 5, -ballSpeed, ballSpeed)
	}
}

func (g *Game) Update() error {
	g.ball.Move(g)
	g.player.Move()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.player.Draw(screen)

	scoreStr := fmt.Sprintf("Score: %d", g.score)

	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
