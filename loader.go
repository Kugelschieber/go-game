package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// Loads textures from png files.
// If keepData is set to true,
// pixel data will be stored inside the texture
// (additionally to VRAM).
type PngLoader struct {
	Filter   int32
	KeepData bool
}

func (p *PngLoader) Load(file string) (Res, error) {
	// load texture
	imgFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	img, err := png.Decode(imgFile)

	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	// create GL texture
	tex := NewTex(gl.TEXTURE_2D)
	tex.Bind()
	tex.SetDefaultParams(p.Filter)
	tex.Texture2D(0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		rgba.Pix)

	if p.KeepData {
		tex.SetRGBA(rgba)
	}

	return tex, nil
}

func (p *PngLoader) Ext() string {
	return "png"
}
