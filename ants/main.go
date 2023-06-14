package main

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Ant struct {
	Position pixel.Vec
	Heading  float64
}

func (a *Ant) Draw(imd *imdraw.IMDraw) {
	imd.Clear()
	imd.Color = colornames.Skyblue
	imd.Push(a.Position)
	imd.Circle(4, 0)
}

func (a *Ant) Move() {
	a.MoveForward()
	if a.Position.X > 1024 {
		a.Position.X = 0
	}
	if a.Position.X < 0 {
		a.Position.X = 1024
	}
	if a.Position.Y > 768 {
		a.Position.Y = 0
	}
	if a.Position.Y < 0 {
		a.Position.Y = 768
	}
}

func (a *Ant) MoveForward() {
	a.Heading += rand.Float64()*math.Pi/2 - math.Pi/4

	a.Position.X += math.Cos(a.Heading)
	a.Position.Y += math.Sin(a.Heading)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//center starting position
	cen := win.Bounds().Center()

	ant := Ant{
		Position: cen,
		Heading:  0,
	}

	//create imdraw

	imd := imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Draw(win)
		ant.Draw(imd)
		ant.Move()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
