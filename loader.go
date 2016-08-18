package goga

import (
	"bufio"
	"errors"
	"github.com/go-gl/gl/v3.2-core/gl"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"
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

// Standford ply file resource.
type Ply struct {
	name string
	path string
	ext  string

	firstLine, data, hasVertex, hasTexCoord, hasNormal bool
	elements, faces                                    int
	indices                                            []uint32
	vertices, texCoords, normals                       []float32

	IndexBuffer, VertexBuffer, TexCoordBuffer, NormalBuffer *VBO
}

// Loads ply files and creates VBOs within the Ply resource.
// The indices must be present as triangles.
// Expected type is float32. If it fails to parse, it will panic.
type PlyLoader struct {
	VboUsage uint32
}

// Drops contained GL buffers.
func (p *Ply) Drop() {
	if p.IndexBuffer != nil {
		p.IndexBuffer.Drop()
	}

	if p.VertexBuffer != nil {
		p.VertexBuffer.Drop()
	}

	if p.TexCoordBuffer != nil {
		p.TexCoordBuffer.Drop()
	}

	if p.NormalBuffer != nil {
		p.NormalBuffer.Drop()
	}
}

// Returns the name of this resource.
func (p *Ply) GetName() string {
	return p.name
}

// Sets the name of this resource.
func (p *Ply) SetName(name string) {
	p.name = name
}

// Returns the path of this resource.
func (p *Ply) GetPath() string {
	return p.path
}

// Sets the path of this resource.
func (p *Ply) SetPath(path string) {
	p.path = path
}

// Returns the file extension of this resource.
func (p *Ply) GetExt() string {
	return p.ext
}

// Sets the file extension of this resource.
func (p *Ply) SetExt(ext string) {
	p.ext = ext
}

func (p *PlyLoader) Load(file string) (Res, error) {
	handle, err := os.Open(file)
	defer handle.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(handle)
	ply := Ply{}
	ply.indices = make([]uint32, 0)
	ply.vertices = make([]float32, 0)
	ply.texCoords = make([]float32, 0)
	ply.normals = make([]float32, 0)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		if ply.data && ply.elements == 0 && ply.faces > 0 {
			ply.faces--

			if err := ply.parseIndices(line); err != nil {
				return nil, err
			}
		}

		if ply.data && ply.elements > 0 {
			ply.elements--

			if err := ply.parseData(line); err != nil {
				return nil, err
			}
		}

		if err := ply.parseHeader(line); err != nil {
			return nil, err
		}

		ply.firstLine = false
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	index, vertex, texCoord, normal := ply.createVBOs(p.VboUsage)
	ply.IndexBuffer = index
	ply.VertexBuffer = vertex
	ply.TexCoordBuffer = texCoord
	ply.NormalBuffer = normal

	return &ply, nil
}

func (p *Ply) createVBOs(vboUsage uint32) (*VBO, *VBO, *VBO, *VBO) {
	index := NewVBO(gl.ELEMENT_ARRAY_BUFFER)
	index.Fill(gl.Ptr(p.indices), 4, len(p.indices), vboUsage)

	vertex := NewVBO(gl.ARRAY_BUFFER)
	vertex.Fill(gl.Ptr(p.vertices), 4, len(p.vertices), vboUsage)

	var texCoord, normal *VBO

	if p.hasTexCoord {
		texCoord = NewVBO(gl.ARRAY_BUFFER)
		texCoord.Fill(gl.Ptr(p.texCoords), 4, len(p.texCoords), vboUsage)
	}

	if p.hasNormal {
		normal = NewVBO(gl.ARRAY_BUFFER)
		normal.Fill(gl.Ptr(p.normals), 4, len(p.normals), vboUsage)
	}

	return index, vertex, texCoord, normal
}

func (p *Ply) parseHeader(line string) error {
	if p.firstLine && line != "ply" { // make sure it's a ply file
		return errors.New("File is not of type ply")
	} else if strings.Contains(line, "element vertex") { // number of elements
		elements, err := strconv.Atoi(line[15:])

		if err != nil {
			return errors.New("Elements could not be parsed")
		}

		p.elements = elements
	} else if strings.Contains(line, "property float") {
		line = line[15:]

		if line == "x" || line == "y" || line == "z" {
			p.hasVertex = true
		} else if line == "nx" || line == "ny" || line == "nz" {
			p.hasNormal = true
		} else if line == "s" || line == "t" {
			p.hasTexCoord = true
		}
	} else if strings.Contains(line, "element face") { // number of faces
		faces, err := strconv.Atoi(line[13:])

		if err != nil {
			return errors.New("Faces could not be parsed")
		}

		p.faces = faces
	} else if strings.Contains(line, "end_header") {
		p.data = true
	}

	return nil
}

func (p *Ply) parseData(line string) error {
	if !p.hasVertex {
		return errors.New("ply must have vertex data")
	}

	parts := strings.Split(line, " ")
	i := 0

	p.vertices = append(p.vertices, parseFloat32(parts[0]))
	p.vertices = append(p.vertices, parseFloat32(parts[1]))
	p.vertices = append(p.vertices, parseFloat32(parts[2]))

	if p.hasNormal {
		i += 3

		p.normals = append(p.normals, parseFloat32(parts[3]))
		p.normals = append(p.normals, parseFloat32(parts[4]))
		p.normals = append(p.normals, parseFloat32(parts[5]))
	}

	if p.hasTexCoord {
		p.texCoords = append(p.texCoords, parseFloat32(parts[3+i]))
		p.texCoords = append(p.texCoords, parseFloat32(parts[4+i]))
	}

	return nil
}

func parseFloat32(str string) float32 {
	float, err := strconv.ParseFloat(str, 32)

	if err != nil {
		panic(err)
	}

	return float32(float)
}

func (p *Ply) parseIndices(line string) error {
	parts := strings.Split(line, " ")

	if len(parts) != 4 || parts[0] != "3" {
		return errors.New("Expected triangles for indices")
	}

	p.indices = append(p.indices, parseUint32(parts[1]))
	p.indices = append(p.indices, parseUint32(parts[2]))
	p.indices = append(p.indices, parseUint32(parts[3]))

	return nil
}

func parseUint32(str string) uint32 {
	i, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return uint32(i)
}

func (p *PlyLoader) Ext() string {
	return "ply"
}
