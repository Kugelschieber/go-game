package main

import (
	"github.com/DeKugelschieber/go-game"
)

const (
	assets_dir = "src/github.com/DeKugelschieber/go-game/demo/model/assets"
)

type Game struct {
	model *goga.Model
}

func (g *Game) Setup() {
	// load texture and ply mesh
	err := goga.LoadResFromFolder(assets_dir)

	if err != nil {
		panic(err)
	}

	// create model
	tex, err := goga.GetTex("cube.png")

	if err != nil {
		panic(err)
	}

	ply, err := goga.GetPly("cube.ply")

	if err != nil {
		panic(err)
	}

	mesh := goga.NewMesh(ply.IndexBuffer, ply.VertexBuffer, ply.TexCoordBuffer)

	model := goga.NewModel(mesh, tex)
	renderer, ok := goga.GetSystemByName("modelRenderer").(*goga.ModelRenderer)

	if !ok {
		panic("Could not find renderer")
	}

	renderer.Prepare(model)
	renderer.Add(model.Actor, model.Pos3D, model.Tex, model.Mesh)
	g.model = model

	// enable depth test and buffer
	goga.EnableDepthTest(true)
	goga.ClearDepthBuffer(true)
}

func (g *Game) Update(delta float64) {
	// update rotation on each frame around upper axis
	g.model.Rot.Z += delta * 45
}

func main() {
	game := Game{}
	options := goga.RunOptions{ClearColor: goga.Vec4{1, 1, 1, 0},
		Resizable:           true,
		SetViewportOnResize: true,
		ExitOnClose:         true}
	goga.Run(&game, &options)
}
