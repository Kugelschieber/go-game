package goga

import (
	"github.com/go-gl/gl/v3.2-core/gl"
)

const (
	sprite_renderer_name = "spriteRenderer"
)

// Sprite is an actor having a 2D position and a texture.
type Sprite struct {
	*Actor
	*Pos2D
	*Tex
}

// Creates a new sprite with given texture.
func NewSprite(tex *Tex) *Sprite {
	sprite := &Sprite{}
	sprite.Actor = NewActor()
	sprite.Pos2D = NewPos2D()
	sprite.Tex = tex
	sprite.Size = Vec2{tex.GetSize().X, tex.GetSize().Y}
	sprite.Scale = Vec2{1, 1}
	sprite.Visible = true

	CheckGLError()

	return sprite
}

// The sprite renderer is a system rendering sprites.
// It has a 2D position component, to move all sprites at once.
type SpriteRenderer struct {
	Pos2D

	Shader *Shader
	Camera *Camera

	sprites                 []Sprite
	index, vertex, texCoord *VBO
	vao                     *VAO
}

// Creates a new sprite renderer using given shader and camera.
// If shader and/or camera are nil, the default one will be used.
func NewSpriteRenderer(shader *Shader, camera *Camera, flip bool) *SpriteRenderer {
	if shader == nil {
		shader = Default2DShader
	}

	if camera == nil {
		camera = DefaultCamera
	}

	renderer := &SpriteRenderer{}
	renderer.Shader = shader
	renderer.Camera = camera
	renderer.sprites = make([]Sprite, 0)
	renderer.index, renderer.vertex, renderer.texCoord = CreateRectMesh(flip)
	renderer.Size = Vec2{1, 1}
	renderer.Scale = Vec2{1, 1}

	renderer.vao = NewVAO()
	renderer.vao.Bind()
	renderer.Shader.EnableVertexAttribArrays()
	renderer.index.Bind()
	renderer.vertex.Bind()
	renderer.vertex.AttribPointer(shader.GetAttribLocation(Default_shader_2D_vertex_attrib), 2, gl.FLOAT, false, 0)
	renderer.texCoord.Bind()
	renderer.texCoord.AttribPointer(shader.GetAttribLocation(Default_shader_2D_texcoord_attrib), 2, gl.FLOAT, false, 0)
	renderer.vao.Unbind()

	CheckGLError()

	return renderer
}

// Frees recources created by sprite renderer.
// This is called automatically when system gets removed.
func (s *SpriteRenderer) Cleanup() {
	s.index.Drop()
	s.vertex.Drop()
	s.texCoord.Drop()
	s.vao.Drop()
}

// Adds sprite to the renderer.
func (s *SpriteRenderer) Add(actor *Actor, pos *Pos2D, tex *Tex) bool {
	id := actor.GetId()

	for _, sprite := range s.sprites {
		if id == sprite.Actor.GetId() {
			return false
		}
	}

	s.sprites = append(s.sprites, Sprite{actor, pos, tex})

	return true
}

// Removes sprite from renderer.
func (s *SpriteRenderer) Remove(actor *Actor) bool {
	return s.RemoveById(actor.GetId())
}

// Removes sprite from renderer by ID.
func (s *SpriteRenderer) RemoveById(id ActorId) bool {
	for i, sprite := range s.sprites {
		if sprite.Actor.GetId() == id {
			s.sprites = append(s.sprites[:i], s.sprites[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all sprites from renderer.
func (s *SpriteRenderer) RemoveAll() {
	s.sprites = make([]Sprite, 0)
}

// Returns number of sprites.
func (s *SpriteRenderer) Len() int {
	return len(s.sprites)
}

func (s *SpriteRenderer) GetName() string {
	return sprite_renderer_name
}

// Render sprites.
func (s *SpriteRenderer) Update(delta float64) {
	s.Shader.Bind()
	s.Shader.SendMat3(Default_shader_2D_ortho, *MultMat3(s.Camera.CalcOrtho(), s.CalcModel()))
	s.Shader.SendUniform1i(Default_shader_2D_tex, 0)
	s.vao.Bind()
	var tid uint32

	for _, sprite := range s.sprites {
		if !sprite.Visible {
			continue
		}

		s.Shader.SendMat3(Default_shader_2D_model, *sprite.CalcModel())

		// prevent texture switching when not neccessary
		if tid != sprite.Tex.GetId() {
			tid = sprite.Tex.GetId()
			sprite.Tex.Bind()
		}

		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	}
}
