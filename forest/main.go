package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	N        = 800
	M        = 800
	EMPTY    = 0
	TREE     = 1
	FIRE     = 2
	CellSize = 5
)

type States int

func probRand() float64 {
	return rand.Float64()
}

func withinGrid(row, col int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= N || col >= M {
		return false
	}

	return true
}

func getNeighbors(x, y int, forest [][]int) bool {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) {
				if withinGrid(i, j) && forest[i][j] == FIRE {
					return true
				}
			}
		}
	}

	return false
}

func print(forest [][]int, win *pixelgl.Window) {
	imd := imdraw.New(nil)
	win.Clear(colornames.Black)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			switch forest[i][j] {
			case EMPTY:
				imd.Color = colornames.Black
			case TREE:
				imd.Color = colornames.Green
			case FIRE:
				imd.Color = colornames.Orange
				break
			}

			imd.Push(pixel.V(float64(i*CellSize), float64(j*CellSize)))
			imd.Push(pixel.V(float64(i*CellSize+CellSize), float64(j*CellSize+CellSize)))
			imd.Rectangle(0)
		}
	}

	imd.Draw(win)
	win.Update()
	time.Sleep(time.Millisecond * 200)
}

func game(forest [][]int, p, f float64) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			switch forest[i][j] {
			case EMPTY:
				if probRand() < p {
					forest[i][j] = TREE
				}
			case TREE:
				if getNeighbors(i, j, forest) || probRand() < f {
					forest[i][j] = FIRE
				}
			case FIRE:
				forest[i][j] = EMPTY
			}
		}
	}
}

func run() {
	forest := make([][]int, N)
	for i := 0; i < N; i++ {
		forest[i] = make([]int, M)
	}

	p := 0.02   // Probability of new growth
	f := 0.0002 // Probability of lightning

	// Matrix init
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			forest[i][j] = EMPTY
		}
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Forest Fire Simulation",
		Bounds: pixel.R(0, 0, float64(M*CellSize), float64(N*CellSize)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		print(forest, win)
		game(forest, p, f)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
