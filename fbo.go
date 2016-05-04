package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

// Frame Buffer Object.
type FBO struct {
	id          uint32
	target      uint32
	attachments []uint32
}

// Creates a new FBO with given target.
func NewFBO(target uint32) *FBO {
	fbo := &FBO{}
	gl.GenFramebuffers(1, &fbo.id)
	fbo.target = target
	fbo.attachments = make([]uint32, 0)

	return fbo
}

// Creates a new FBO with given target and 2D texture.
// Used to render to texture.
func NewFBOWithTex2D(width, height, filter int32) (*FBO, *Tex) {
	tex := NewTex(gl.TEXTURE_2D)
	tex.Bind()
	tex.SetDefaultParams(filter)
	tex.Texture2D(0,
		gl.RGBA,
		width,
		height,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		nil)
	tex.Unbind()

	fbo := NewFBO(gl.FRAMEBUFFER)
	fbo.Bind()
	fbo.Texture2D(gl.COLOR_ATTACHMENT0, tex.GetId(), 0)
	fbo.Unbind()

	return fbo, tex
}

// Drops the FBO.
func (f *FBO) Drop() {
	gl.DeleteFramebuffers(1, &f.id)
}

// Binds the FBO for usage (rendered on).
func (f *FBO) Bind() {
	gl.BindFramebuffer(f.target, f.id)
}

// Unbinds.
func (f *FBO) Unbind() {
	gl.BindFramebuffer(f.target, 0)
}

// Sets draw buffer.
func (f *FBO) DrawBuffer(mode uint32) {
	gl.DrawBuffer(mode)
}

// Sets read buffer.
func (f *FBO) ReadBuffer(mode uint32) {
	gl.ReadBuffer(mode)
}

func (f *FBO) DrawBuffers(mode uint32) {
	gl.DrawBuffers(int32(len(f.attachments)), &f.attachments[0])
}

// Removes all attached textures from FBO.
func (f *FBO) ClearAttachments() {
	f.attachments = make([]uint32, 0)
}

// Attaches a texture.
func (f *FBO) Texture(attachment, texId uint32, level int32) {
	f.attachments = append(f.attachments, attachment)
	gl.FramebufferTexture(f.target, attachment, texId, level)
}

// Attaches a 1D texture.
func (f *FBO) Texture1D(attachment, texId uint32, level int32) {
	f.attachments = append(f.attachments, attachment)
	gl.FramebufferTexture1D(f.target, attachment, gl.TEXTURE_1D, texId, level)
}

// Attaches a 2D texture.
func (f *FBO) Texture2D(attachment, texId uint32, level int32) {
	f.attachments = append(f.attachments, attachment)
	gl.FramebufferTexture2D(f.target, attachment, gl.TEXTURE_2D, texId, level)
}

// Attaches a 3D texture.
func (f *FBO) Texture3D(attachment, texId uint32, level, layer int32) {
	f.attachments = append(f.attachments, attachment)
	gl.FramebufferTexture3D(f.target, attachment, gl.TEXTURE_3D, texId, level, layer)
}

// Returns the status of the FBO.
func (f *FBO) GetStatus() uint32 {
	return gl.CheckFramebufferStatus(f.target)
}

// Returns true if FBO is complete, else false.
func (f *FBO) Complete() bool {
	return f.GetStatus() == gl.FRAMEBUFFER_COMPLETE
}

// Returns the GL ID.
func (f *FBO) GetId() uint32 {
	return f.id
}

// Returns the target.
func (f *FBO) GetTarget() uint32 {
	return f.target
}
