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
	paddleSpeed  = 4
)

// Game implements ebiten.Game interface.
type Game struct {
	ball        Ball
	player      Paddle
	cpu         CpuPaddle
	playerScore int
	cpuScore    int
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
	g.player = *NewPaddle(screenWidth-10, screenHeight/2-20, 5, 40)
	g.cpu = *NewCpuPaddle(5, screenHeight/2-20, 5, 40)
	g.playerScore = 0
	g.cpuScore = 0

}

func (g *Game) Reset() {
	g.RandomBall()
}

func (g *Game) RandomBall() {
	r := rand.Intn(4)

	if r == 0 {
		g.ball = *NewBall(
			Vector2D{X: ballSpeed, Y: ballSpeed},
			Vector2D{X: screenWidth / 2, Y: screenHeight / 2},
			5,
		)
	} else if r == 1 {
		g.ball = *NewBall(
			Vector2D{X: -ballSpeed, Y: ballSpeed},
			Vector2D{X: screenWidth / 2, Y: screenHeight / 2},
			5,
		)
	} else if r == 2 {
		g.ball = *NewBall(
			Vector2D{X: ballSpeed, Y: -ballSpeed},
			Vector2D{X: screenWidth / 2, Y: screenHeight / 2},
			5,
		)
	} else {
		g.ball = *NewBall(
			Vector2D{X: -ballSpeed, Y: -ballSpeed},
			Vector2D{X: screenWidth / 2, Y: screenHeight / 2},
			5,
		)
	}
}

func (g *Game) Update() error {
	g.ball.Move(g)
	g.player.Move()
	g.cpu.Move(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.player.Draw(screen)
	g.cpu.Draw(screen)

	playerScoreStr := fmt.Sprintf("Player: %d", g.playerScore)
	cpuScoreStr := fmt.Sprintf("CPU: %d", g.cpuScore)

	text.Draw(screen, playerScoreStr, basicfont.Face7x13, 10, 20, color.White)
	text.Draw(screen, cpuScoreStr, basicfont.Face7x13, 10, 40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
