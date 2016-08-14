package goga

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"unsafe"
)

// Vertex Buffer Object.
type VBO struct {
	id     uint32
	target uint32
	size   int32
}

// Creates a new VBO with given target.
func NewVBO(target uint32) *VBO {
	vbo := &VBO{target: target}
	gl.GenBuffers(1, &vbo.id)

	return vbo
}

// Drops this VBO.
func (v *VBO) Drop() {
	gl.DeleteBuffers(1, &v.id)
}

// Binds VBO for rendering.
func (v *VBO) Bind() {
	gl.BindBuffer(v.target, v.id)
}

// Unbinds.
func (v *VBO) Unbind() {
	gl.BindBuffer(v.target, 0)
}

// Fills VBO with data.
// An unsafe pointer must be used out of Gos unsafe package.
func (v *VBO) Fill(data unsafe.Pointer, elements, size int, use uint32) {
	v.size = int32(size)

	v.Bind()
	gl.BufferData(v.target, elements*size, data, use)
	v.Unbind()
}

// Updates data or part of data.
// An unsafe pointer must be used out of Gos unsafe package.
func (v *VBO) Update(data unsafe.Pointer, elements, offset, size int) {
	v.size = int32(size)

	v.Bind()
	gl.BufferSubData(v.target, offset, elements*size, data)
	v.Unbind()
}

// Sets the attribute pointer for rendering.
// Used together with shader.
func (v *VBO) AttribPointer(attribLocation int32, size int32, btype uint32, normalized bool, stride int32) {
	gl.VertexAttribPointer(uint32(attribLocation), size, btype, normalized, stride, nil)
}

// Returns the GL ID.
func (v *VBO) GetId() uint32 {
	return v.id
}

// Returns the target.
func (v *VBO) GetTarget() uint32 {
	return v.target
}

// Returns the number of elements within this VBO.
func (v *VBO) Size() int32 {
	return v.size
}
