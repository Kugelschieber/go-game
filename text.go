package goga

import (
	"encoding/json"
	"github.com/go-gl/gl/v3.2-core/gl"
	"io/ioutil"
)

const (
	char_padding       = 2
	text_renderer_name = "textRenderer"
)

type character struct {
	char           byte
	min, max, size Vec2
	offset         float64
}

type jsonChar struct {
	Char         string
	X, Y, Offset float64
}

// Font represents a texture mapped font.
// It can be loaded from JSON together with a texture.
type Font struct {
	Tex              *Tex
	tileSize         float64
	CharPadding      Vec2
	Space, Tab, Line float64
	chars            []character
}

// Creates a new font for given texture.
// The tile size specifies the size of one character tile on texture.
// Characters must be added afterwards.
func NewFont(tex *Tex, tileSize float64) *Font {
	font := Font{}
	font.Tex = tex
	font.tileSize = tileSize
	font.CharPadding = Vec2{0.05, 0.05}
	font.Space = 0.3
	font.Tab = 1.2
	font.Line = 1
	font.chars = make([]character, 0)

	return &font
}

// Loads characters from JSON file.
// Format:
//
// [
//     {
//         "char": "a",
//         "x": 0,
//         "y": 0,
//		   "offset": 0
//     },
//     ...
// ]
//
// Where x and y start in the upper left corner of the texture, both of type int.
// Offset is optional and can be used to move a character up or down (relative to others).
// If cut is set to true, the characters will be true typed.
func (f *Font) FromJson(path string, cut bool) error {
	// load file content
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	// read json
	chars := make([]jsonChar, 0)

	if err = json.Unmarshal(content, &chars); err != nil {
		return err
	}

	f.extractChars(chars, cut)

	return nil
}

func (f *Font) extractChars(chars []jsonChar, cut bool) {
	for _, char := range chars {
		if len(char.Char) != 1 {
			continue
		}

		var min, max, size Vec2

		if !cut {
			min = Vec2{char.X * f.tileSize, char.Y * f.tileSize}
			max = Vec2{min.X + f.tileSize, min.Y + f.tileSize}
			size = Vec2{1, 1}
		} else {
			min, max, size = f.cutChar(int(char.X), int(char.Y))
		}

		f.chars = append(f.chars, character{char.Char[0], min, max, size, char.Offset})
	}
}

func (f *Font) cutChar(x, y int) (Vec2, Vec2, Vec2) {
	// find min/max corners of character on texture
	minX := int(f.Tex.GetSize().X)
	minY := int(f.Tex.GetSize().Y)
	maxX := 0
	maxY := 0
	rgba := f.Tex.GetRGBA()

	for ry := y * int(f.tileSize); ry < (y+1)*int(f.tileSize); ry++ {
		for rx := x * int(f.tileSize); rx < (x+1)*int(f.tileSize); rx++ {
			_, _, _, a := rgba.At(rx, ry).RGBA()

			if a == 0 {
				continue
			}

			if rx < minX {
				minX = rx
			} else if rx > maxX {
				maxX = rx
			}

			if ry < minY {
				minY = ry
			} else if ry > maxY {
				maxY = ry
			}
		}
	}

	// add padding
	minX -= char_padding
	maxX += char_padding
	minY -= char_padding
	maxY += char_padding

	texSize := f.Tex.GetSize()
	min := Vec2{float64(minX) / texSize.X, float64(maxY) / texSize.Y}
	max := Vec2{float64(maxX) / texSize.X, float64(minY) / texSize.Y}
	size := Vec2{float64(maxX-minX) / f.tileSize, float64(maxY-minY) / f.tileSize}

	return min, max, size
}

func (f *Font) getChar(char byte) *character {
	for _, character := range f.chars {
		if character.char == char {
			return &character
		}
	}

	return nil
}

// Renderable text component.
// Use together with Text and create using NewText().
type TextComponent struct {
	Color Vec4

	text                    string
	bounds                  Vec2
	index, vertex, texCoord *VBO
	vao                     *VAO
}

// Deletes GL buffers bound to this text component.
func (t *TextComponent) Drop() {
	t.index.Drop()
	t.vertex.Drop()
	t.texCoord.Drop()
	t.vao.Drop()
}

// Text is an actor representing text rendered as texture mapped font.
// Each Text has a position and its own buffers.
type Text struct {
	*Actor
	*Pos2D
	*TextComponent
}

// Returns a new renderable text object.
func NewText(font *Font, textStr string) *Text {
	text := Text{}
	text.Actor = NewActor()
	text.Pos2D = NewPos2D()
	text.TextComponent = &TextComponent{}
	text.index = NewVBO(gl.ELEMENT_ARRAY_BUFFER)
	text.vertex = NewVBO(gl.ARRAY_BUFFER)
	text.texCoord = NewVBO(gl.ARRAY_BUFFER)
	text.vao = NewVAO()
	text.SetText(font, textStr)
	text.Color = Vec4{1, 1, 1, 1}
	text.Size = Vec2{1, 1}
	text.Scale = Vec2{1, 1}
	text.Visible = true

	return &text
}

// Sets the given string as text and (re)creates buffers.
func (t *Text) SetText(font *Font, text string) {
	t.text = text
	indices := make([]uint32, len(text)*6)
	vertices := make([]float32, len(text)*8)
	texCoords := make([]float32, len(text)*8)
	chars := 0

	// create indices
	var index uint32 = 0

	for i := 0; i < len(text)*6; i += 6 {
		indices[i] = index
		indices[i+1] = index + 1
		indices[i+2] = index + 2
		indices[i+3] = index + 1
		indices[i+4] = index + 2
		indices[i+5] = index + 3

		index += 4
	}

	// create vertices/texCoords
	index = 0
	offset := Vec2{}
	var width, height float64

	for i := 0; i < len(text)*8 && int(index) < len(text); i += 8 {
		c := font.getChar(text[index])
		index++

		// whitespace and new line
		if text[index-1] == ' ' {
			offset.X += font.Space
			i -= 8
			continue
		}

		if text[index-1] == '\n' {
			offset.X = 0
			offset.Y -= font.Line
			i -= 8
			continue
		}

		if text[index-1] == '\t' {
			offset.X += font.Tab
			i -= 8
			continue
		}

		// character not found
		if c == nil {
			i -= 8
			continue
		}

		// usual character
		vertices[i] = float32(offset.X)
		vertices[i+1] = float32(offset.Y + c.offset)
		vertices[i+2] = float32(offset.X + c.size.X)
		vertices[i+3] = float32(offset.Y + c.offset)
		vertices[i+4] = float32(offset.X)
		vertices[i+5] = float32(offset.Y + c.size.Y + c.offset)
		vertices[i+6] = float32(offset.X + c.size.X)
		vertices[i+7] = float32(offset.Y + c.size.Y + c.offset)

		texCoords[i] = float32(c.min.X)
		texCoords[i+1] = float32(c.min.Y)
		texCoords[i+2] = float32(c.max.X)
		texCoords[i+3] = float32(c.min.Y)
		texCoords[i+4] = float32(c.min.X)
		texCoords[i+5] = float32(c.max.Y)
		texCoords[i+6] = float32(c.max.X)
		texCoords[i+7] = float32(c.max.Y)

		offset.X += c.size.X + font.CharPadding.X
		chars++

		if offset.X > width {
			width = offset.X
		}

		if offset.Y*-1+font.Line > height {
			height = offset.Y*-1 + font.Line
		}
	}

	t.bounds = Vec2{width, height}

	// fill GL buffer
	t.index.Fill(gl.Ptr(indices[:chars*6]), 4, chars*6, gl.STATIC_DRAW)
	t.vertex.Fill(gl.Ptr(vertices[:chars*8]), 4, chars*8, gl.STATIC_DRAW)
	t.texCoord.Fill(gl.Ptr(texCoords[:chars*8]), 4, chars*8, gl.STATIC_DRAW)

	CheckGLError()
}

// Returns the text as string.
func (t *Text) GetText() string {
	return t.text
}

// Returns bounds of text, which is the size of characters.
func (t *Text) GetBounds() Vec2 {
	return Vec2{t.bounds.X * t.Size.X * t.Scale.X, t.bounds.Y * t.Size.Y * t.Scale.Y}
}

// The text renderer is a system rendering 2D texture mapped font.
// It has a 2D position component, to move all texts at once.
type TextRenderer struct {
	Pos2D

	Shader *Shader
	Camera *Camera
	Font   *Font
	texts  []Text
}

// Creates a new text renderer using given shader, camera and font.
// If shader and/or camera are nil, the default one will be used.
func NewTextRenderer(shader *Shader, camera *Camera, font *Font) *TextRenderer {
	if shader == nil {
		shader = DefaultTextShader
	}

	if camera == nil {
		camera = DefaultCamera
	}

	renderer := &TextRenderer{}
	renderer.Shader = shader
	renderer.Camera = camera
	renderer.Font = font
	renderer.texts = make([]Text, 0)
	renderer.Size = Vec2{1, 1}
	renderer.Scale = Vec2{1, 1}

	return renderer
}

// Prepares given text for rendering.
func (r *TextRenderer) Prepare(text *Text) {
	text.vao = NewVAO()
	text.vao.Bind()
	r.Shader.EnableVertexAttribArrays()
	text.index.Bind()
	text.vertex.Bind()
	text.vertex.AttribPointer(r.Shader.GetAttribLocation(Default_shader_text_vertex_attrib), 2, gl.FLOAT, false, 0)
	text.texCoord.Bind()
	text.texCoord.AttribPointer(r.Shader.GetAttribLocation(Default_shader_text_texcoord_attrib), 2, gl.FLOAT, false, 0)
	text.vao.Unbind()
}

// Frees recources created by text component.
// This is called automatically when system gets removed.
func (r *TextRenderer) Cleanup() {
	for _, text := range r.texts {
		text.Drop()
	}
}

// Adds text to the renderer.
func (r *TextRenderer) Add(actor *Actor, pos *Pos2D, text *TextComponent) bool {
	id := actor.GetId()

	for _, text := range r.texts {
		if id == text.Actor.GetId() {
			return false
		}
	}

	r.texts = append(r.texts, Text{actor, pos, text})

	return true
}

// Removes text from renderer.
func (r *TextRenderer) Remove(actor *Actor) bool {
	return r.RemoveById(actor.GetId())
}

// Removes text from renderer by ID.
func (r *TextRenderer) RemoveById(id ActorId) bool {
	for i, text := range r.texts {
		if text.Actor.GetId() == id {
			r.texts = append(r.texts[:i], r.texts[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all texts.
func (r *TextRenderer) RemoveAll() {
	r.texts = make([]Text, 0)
}

// Returns number of texts.
func (r *TextRenderer) Len() int {
	return len(r.texts)
}

func (r *TextRenderer) GetName() string {
	return text_renderer_name
}

// Renders texts.
func (r *TextRenderer) Update(delta float64) {
	if r.Font == nil {
		return
	}

	r.Shader.Bind()
	r.Shader.SendMat3(Default_shader_text_ortho, *MultMat3(r.Camera.CalcOrtho(), r.CalcModel()))
	r.Shader.SendUniform1i(Default_shader_text_tex, 0)
	r.Font.Tex.Bind()

	for i := range r.texts {
		if !r.texts[i].Visible {
			continue
		}

		r.texts[i].vao.Bind()
		r.Shader.SendUniform4f(Default_shader_text_color, float32(r.texts[i].Color.X), float32(r.texts[i].Color.Y), float32(r.texts[i].Color.Z), float32(r.texts[i].Color.W))
		r.Shader.SendMat3(Default_shader_text_model, *r.texts[i].CalcModel())

		gl.DrawElements(gl.TRIANGLES, r.texts[i].index.Size(), gl.UNSIGNED_INT, nil)
	}
}
