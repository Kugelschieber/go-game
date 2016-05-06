package main

import (
	"github.com/DeKugelschieber/go-game"
)

const (
	gopher_path = "src/github.com/DeKugelschieber/go-game/demo/sprite/assets/gopher.png"
)

type Game struct{}

func (g *Game) Setup() {
	res, err := goga.LoadRes(gopher_path)

	if err != nil {
		panic(err)
	}

	tex, ok := res.(*goga.Tex)

	if !ok {
		panic("Resource is not a texture")
	}

	sprite := goga.NewSprite(tex)
	goga.AddActor(sprite)
}

func (g *Game) Update(delta float64) {
}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{1, 1, 1, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true}
	goga.Run(&game, &options)
}
