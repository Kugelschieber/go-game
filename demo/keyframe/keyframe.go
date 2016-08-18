package main

import (
	"github.com/DeKugelschieber/go-game"
)

const (
	assets_dir = "src/github.com/DeKugelschieber/go-game/demo/keyframe/assets"
)

type Game struct{}

func (g *Game) Setup() {
	err := goga.LoadResFromFolder(assets_dir)

	if err != nil {
		panic(err)
	}

	tex, err := goga.GetTex("runningcat.png")

	if err != nil {
		panic(err)
	}

	// create a keyframe set
	set := goga.NewKeyframeSet()
	set.Add(goga.NewKeyframe(goga.Vec2{0, 0}, goga.Vec2{0.5, 0.25}))
	set.Add(goga.NewKeyframe(goga.Vec2{0.5, 0}, goga.Vec2{1, 0.25}))
	set.Add(goga.NewKeyframe(goga.Vec2{0, 0.25}, goga.Vec2{0.5, 0.5}))
	set.Add(goga.NewKeyframe(goga.Vec2{0.5, 0.25}, goga.Vec2{1, 0.5}))
	set.Add(goga.NewKeyframe(goga.Vec2{0, 0.5}, goga.Vec2{0.5, 0.75}))
	set.Add(goga.NewKeyframe(goga.Vec2{0.5, 0.5}, goga.Vec2{1, 0.75}))
	set.Add(goga.NewKeyframe(goga.Vec2{0, 0.75}, goga.Vec2{0.5, 1}))
	set.Add(goga.NewKeyframe(goga.Vec2{0.5, 0.75}, goga.Vec2{1, 1}))

	// create a new animated sprite
	sprite := goga.NewAnimatedSprite(tex, set, 512, 256)
	sprite.KeyframeAnimation = goga.NewKeyframeAnimation(0, 7, true, 20)

	// add to renderer
	renderer := goga.GetKeyframeRenderer()
	renderer.Add(sprite.Actor, sprite.Pos2D, sprite.Tex, sprite.KeyframeSet, sprite.KeyframeAnimation)
}

func (g *Game) Update(delta float64) {}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{0, 0, 0, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true,
		Fullscreen:          true}
	goga.Run(&game, &options)
}
