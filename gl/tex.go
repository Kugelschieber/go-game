package gl

import (
	"geo"
	"github.com/go-gl/gl/v4.5-core/gl"
	"image"
)

// Texture object.
type Tex struct {
	id            uint32
	target        uint32
	activeTexture uint32
	size          geo.Vec3
	rgba          *image.RGBA // optional, most of the time nil
}

// Creates a new texture for given target (e.g. GL_TEXTURE_2D).
func NewTex(target uint32) *Tex {
	tex := &Tex{}
	tex.target = target
	tex.activeTexture = gl.TEXTURE0
	gl.GenTextures(1, &tex.id)

	return tex
}

// Drops the texture.
func (t *Tex) Drop() {
	gl.DeleteBuffers(1, &t.id)
}

// Binds the texture for rendering.
func (t *Tex) Bind() {
	gl.ActiveTexture(t.activeTexture)
	gl.BindTexture(t.target, t.id)
}

// Unbinds.
func (t *Tex) Unbind() {
	gl.BindTexture(t.target, 0)
}

// Sets the default parameters, which are passed filter and CLAMP_TO_EDGE.
func (t *Tex) SetDefaultParams(filter int32) {
	t.Parameteri(gl.TEXTURE_MIN_FILTER, filter)
	t.Parameteri(gl.TEXTURE_MAG_FILTER, filter)
	t.Parameteri(gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	t.Parameteri(gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
}

// Creates a new 1D texture.
func (t *Tex) Texture1D(level, internalFormat, width int32, format, ttype uint32, data []uint8) {
	t.size = geo.Vec3{float64(width), 0, 0}
	t.Bind()

	if data != nil {
		gl.TexImage1D(t.target, level, internalFormat, width, 0, format, ttype, gl.Ptr(data))
	} else {
		gl.TexImage1D(t.target, level, internalFormat, width, 0, format, ttype, nil)
	}
}

// Creates a new 2D texture.
func (t *Tex) Texture2D(level, internalFormat, width, height int32, format, ttype uint32, data []uint8) {
	t.size = geo.Vec3{float64(width), float64(height), 0}
	t.Bind()

	if data != nil {
		gl.TexImage2D(t.target, level, internalFormat, width, height, 0, format, ttype, gl.Ptr(data))
	} else {
		gl.TexImage2D(t.target, level, internalFormat, width, height, 0, format, ttype, nil)
	}
}

// Creates a new 3D texture.
func (t *Tex) Texture3D(level, internalFormat, width, height, depth int32, format, ttype uint32, data []uint8) {
	t.size = geo.Vec3{float64(width), float64(height), float64(depth)}
	t.Bind()

	if data != nil {
		gl.TexImage3D(t.target, level, internalFormat, width, height, depth, 0, format, ttype, gl.Ptr(data))
	} else {
		gl.TexImage3D(t.target, level, internalFormat, width, height, depth, 0, format, ttype, nil)
	}
}

// Sets integer parameter.
func (t *Tex) Parameteri(name uint32, param int32) {
	gl.TexParameteri(t.target, name, param)
}

// Sets float parameter.
func (t *Tex) Parameterf(name uint32, param float32) {
	gl.TexParameterf(t.target, name, param)
}

// Sets which texture boundary is used when bound for rendering.
// Can be GL_TEXTURE0, GL_TEXTURE1, ... GL_TEXTUREn.
func (t *Tex) SetActiveTexture(activeTexture uint32) {
	t.activeTexture = activeTexture
}

// Sets pixel data.
func (t *Tex) SetRGBA(rgba *image.RGBA) {
	t.rgba = rgba
}

// Returns the GL ID.
func (t *Tex) GetId() uint32 {
	return t.id
}

// Returns the texture target
func (t *Tex) GetTarget() uint32 {
	return t.target
}

// Returns the active texture used when bound for rendering.
func (t *Tex) getActiveTexture() uint32 {
	return t.activeTexture
}

// Returns the size of this texture.
func (t *Tex) GetSize() geo.Vec3 {
	return t.size
}

// Returns the pixel data.
func (t *Tex) GetRGBA() *image.RGBA {
	return t.rgba
}
