package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

// Vertex Array Object.
type VAO struct {
	id uint32
}

// Creates a new VAO.
// Bind and use VBOs for rendering later.
func NewVAO() *VAO {
	vao := &VAO{}
	gl.GenVertexArrays(1, &vao.id)

	return vao
}

// Drops this VAO.
func (v *VAO) Drop() {
	gl.DeleteVertexArrays(1, &v.id)
}

// Binds VAO for rendering.
func (v *VAO) Bind() {
	gl.BindVertexArray(v.id)
}

// Unbinds.
func (v *VAO) Unbind() {
	gl.BindVertexArray(0)
}

// Returns the GL ID.
func (v *VAO) GetId() uint32 {
	return v.id
}
