package main

import (
	"github.com/DeKugelschieber/go-game"
	"github.com/go-gl/gl/v3.2-core/gl"
)

const (
	font_path = "src/github.com/DeKugelschieber/go-game/demo/text/assets/victor.png"
	font_json = "src/github.com/DeKugelschieber/go-game/demo/text/assets/victor.json"
)

type Game struct{}

func (g *Game) Setup() {
	// load texture
	pngLoader, ok := goga.GetLoaderByExt("png").(*goga.PngLoader)

	if !ok {
		panic("Could not get PNG loader")
	}

	pngLoader.KeepData = true
	pngLoader.Filter = gl.NEAREST
	_, err := goga.LoadRes(font_path)

	if err != nil {
		panic(err)
	}

	pngLoader.KeepData = false
	pngLoader.Filter = gl.LINEAR

	// create font
	tex, err := goga.GetTex("victor.png")

	if err != nil {
		panic(err)
	}

	font := goga.NewFont(tex, 16)

	if err := font.FromJson(font_json, true); err != nil {
		panic(err)
	}

	// setup renderer
	renderer, ok := goga.GetSystemByName("textRenderer").(*goga.TextRenderer)

	if !ok {
		panic("Could not find renderer")
	}

	renderer.Font = font

	// create and add text
	text := goga.NewText(font, "Hello, World!_")
	text.Size = goga.Vec2{16, 16}
	text.Pos = goga.Vec2{20, 20}
	renderer.Prepare(text)
	renderer.Add(text.Actor, text.Pos2D, text.TextComponent)
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
