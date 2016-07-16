package main

import (
	"github.com/DeKugelschieber/go-game"
)

const (
	font_path = "src/github.com/DeKugelschieber/go-game/demo/text/assets/font.png"
)

type Game struct{}

func (g *Game) Setup() {
	// load texture
	/*_, err := goga.LoadRes(gopher_path)

	if err != nil {
		panic(err)
	}

	// create sprite
	tex, err := goga.GetTex("gopher.png")

	if err != nil {
		panic(err)
	}

	sprite := goga.NewSprite(tex)
	renderer, ok := goga.GetSystemByName("spriteRenderer").(*goga.SpriteRenderer)

	if !ok {
		panic("Could not find renderer")
	}

	renderer.Add(sprite.Actor, sprite.Pos2D, sprite.Tex)

	culling, ok := goga.GetSystemByName("culling2d").(*goga.Culling2D)

	if !ok {
		panic("Could not find culling")
	}

	culling.Add(sprite.Actor, sprite.Pos2D)*/
}

func (g *Game) Update(delta float64) {}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{0, 0, 0, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true}
	goga.Run(&game, &options)
}
