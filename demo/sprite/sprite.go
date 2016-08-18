package main

import (
	"github.com/DeKugelschieber/go-game"
)

const (
	gopher_path = "src/github.com/DeKugelschieber/go-game/demo/sprite/assets/gopher.png"
)

type Game struct{}

func (g *Game) Setup() {
	// load texture
	_, err := goga.LoadRes(gopher_path)

	if err != nil {
		panic(err)
	}

	// create sprite
	tex, err := goga.GetTex("gopher.png")

	if err != nil {
		panic(err)
	}

	sprite := goga.NewSprite(tex)
	renderer := goga.GetSpriteRenderer()
	renderer.Add(sprite.Actor, sprite.Pos2D, sprite.Tex)

	culling := goga.GetCulling2DSystem()
	culling.Add(sprite.Actor, sprite.Pos2D)
}

func (g *Game) Update(delta float64) {}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{1, 1, 1, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true}
	goga.Run(&game, &options)
}
