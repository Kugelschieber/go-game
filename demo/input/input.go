package main

import (
	"github.com/DeKugelschieber/go-game"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	gopher_path = "src/github.com/DeKugelschieber/go-game/demo/input/assets/gopher.png"
)

type Game struct {
	mouseX, mouseY float64
	sprite         *goga.Sprite
}

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
	sprite.Size.X = sprite.Size.X / 4
	sprite.Size.Y = sprite.Size.Y / 4
	g.sprite = sprite
	renderer, ok := goga.GetSystemByName("spriteRenderer").(*goga.SpriteRenderer)

	if !ok {
		panic("Could not find renderer")
	}

	renderer.Add(sprite.Actor, sprite.Pos2D, sprite.Tex)

	culling, ok := goga.GetSystemByName("culling2d").(*goga.Culling2D)

	if !ok {
		panic("Could not find culling")
	}

	culling.Add(sprite.Actor, sprite.Pos2D)

	// register input listeners
	goga.AddKeyboardListener(g)
	goga.AddMouseListener(g)
}

func (g *Game) Update(delta float64) {}

func (g *Game) OnKeyEvent(key glfw.Key, code int, action glfw.Action, mod glfw.ModifierKey) {
	// ESC
	if key == 256 {
		goga.Stop()
	}
}

func (g *Game) OnMouseButton(button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if button == 0 {
		g.sprite.Pos.X = g.mouseX
		g.sprite.Pos.Y = g.mouseY
	}
}

func (g *Game) OnMouseMove(x float64, y float64) {
	g.mouseX = x
	g.mouseY = y
}

func (g *Game) OnMouseScroll(x float64, y float64) {}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{1, 1, 1, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true}
	goga.Run(&game, &options)
}
