package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"log"
)

// Checks for GL errors and prints to log if one occured.
func CheckGLError() {
	error := gl.GetError()

	if error != 0 {
		log.Print(error)
	}
}

// Creates three VBOs for a 2D rectangle.
func CreateRectMesh(flip bool) (*VBO, *VBO, *VBO) {
	ib, vb, tb := createIndexVertexTexCoordBuffer()

	indexData := []uint32{0, 1, 2, 1, 2, 3}
	vertexData := []float32{0, 0, 1, 0, 0, 1, 1, 1}
	texData := make([]float32, 8)

	if flip {
		texData = []float32{0, 0, 1, 0, 0, 1, 1, 1}
	} else {
		texData = []float32{0, 1, 1, 1, 0, 0, 1, 0}
	}

	ib.Fill(gl.Ptr(indexData), 4, 6, gl.STATIC_DRAW)
	vb.Fill(gl.Ptr(vertexData), 4, 8, gl.STATIC_DRAW)
	tb.Fill(gl.Ptr(texData), 4, 8, gl.STATIC_DRAW)

	return ib, vb, tb
}

// Creates three VBOs for a 3D cube mesh.
// Texture coordinates won't map properly.
// This function is supposed to be used for 3D testing.
func CreateCubeMesh() (*VBO, *VBO, *VBO) {
	ib, vb, tb := createIndexVertexTexCoordBuffer()

	indexData := []uint32{3, 1, 0,
		0, 2, 3,
		7, 5, 1,
		1, 3, 7,
		6, 4, 5,
		5, 7, 6,
		2, 0, 4,
		4, 6, 2,
		7, 3, 2,
		2, 6, 7,
		1, 5, 4,
		4, 0, 1}
	vertexData := []float32{0, 0, 0,
		1, 0, 0,
		0, 1, 0,
		1, 1, 0,
		0, 0, 1,
		1, 0, 1,
		0, 1, 1,
		1, 1, 1}
	texData := []float32{0, 0,
		1, 0,
		0, 1,
		1, 1,
		1, 1,
		0, 1,
		1, 0,
		0, 0}

	ib.Fill(gl.Ptr(indexData), 4, 36, gl.STATIC_DRAW)
	vb.Fill(gl.Ptr(vertexData), 4, 24, gl.STATIC_DRAW)
	tb.Fill(gl.Ptr(texData), 4, 12, gl.STATIC_DRAW)

	return ib, vb, tb
}

func createIndexVertexTexCoordBuffer() (*VBO, *VBO, *VBO) {
	return NewVBO(gl.ELEMENT_ARRAY_BUFFER),
		NewVBO(gl.ARRAY_BUFFER),
		NewVBO(gl.ARRAY_BUFFER)
}
