package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

const (
	model_renderer_name = "modelRenderer"
)

// Component representing a 3D mesh.
type Mesh struct {
	Index, Vertex, TexCoord *VBO
	Vao                     *VAO
}

// Model is an actor having a 3D position, a texture and a 3D mesh.
type Model struct {
	*Actor
	*Pos3D
	*Tex
	*Mesh
}

// The model renderer is a system rendering models.
// It has a 3D position component, to move all models at once.
type ModelRenderer struct {
	Pos3D

	Shader *Shader
	Camera *Camera
	ortho  bool

	models []Model
}

// Creates a new mesh with given GL buffers.
// The VAO must be prepared by ModelRenderer.
func NewMesh(index, vertex, texcoord *VBO) *Mesh {
	mesh := &Mesh{}
	mesh.Index = index
	mesh.Vertex = vertex
	mesh.TexCoord = texcoord

	CheckGLError()

	return mesh
}

// Drops the VBOs and VAO contained in mesh.
// This must not be done, if mesh was filled from outer source (like a ply file).
func (m *Mesh) Drop() {
	m.Index.Drop()
	m.Vertex.Drop()
	m.TexCoord.Drop()
	m.Vao.Drop()
}

// Creates a new model with given mesh and texture.
func NewModel(mesh *Mesh, tex *Tex) *Model {
	model := &Model{}
	model.Actor = NewActor()
	model.Pos3D = NewPos3D()
	model.Tex = tex
	model.Mesh = mesh
	model.Size = Vec3{1, 1, 1}
	model.Scale = Vec3{1, 1, 1}
	model.Visible = true

	CheckGLError()

	return model
}

// Creates a new model renderer using given shader and camera.
// If shader and/or camera are nil, the default one will be used.
// Orth can be set to true, to use orthogonal projection.
func NewModelRenderer(shader *Shader, camera *Camera, ortho bool) *ModelRenderer {
	if shader == nil {
		shader = Default3DShader
	}

	if camera == nil {
		camera = DefaultCamera
	}

	renderer := &ModelRenderer{}
	renderer.Shader = shader
	renderer.Camera = camera
	renderer.ortho = ortho
	renderer.models = make([]Model, 0)
	renderer.Size = Vec3{1, 1, 1}
	renderer.Scale = Vec3{1, 1, 1}

	CheckGLError()

	return renderer
}

func (s *ModelRenderer) Cleanup() {}

// Prepares a model to be rendered by setting up its VAO.
func (s *ModelRenderer) Prepare(model *Model) {
	model.Vao = NewVAO()
	model.Vao.Bind()
	s.Shader.EnableVertexAttribArrays()
	model.Index.Bind()
	model.Vertex.Bind()
	model.Vertex.AttribPointer(s.Shader.GetAttribLocation(Default_shader_3D_vertex_attrib), 3, gl.FLOAT, false, 0)
	model.TexCoord.Bind()
	model.TexCoord.AttribPointer(s.Shader.GetAttribLocation(Default_shader_3D_texcoord_attrib), 2, gl.FLOAT, false, 0)
	model.Vao.Unbind()
}

// Adds model to the renderer.
// Perpare it first!
func (s *ModelRenderer) Add(actor *Actor, pos *Pos3D, tex *Tex, mesh *Mesh) bool {
	id := actor.GetId()

	for _, model := range s.models {
		if id == model.Actor.GetId() {
			return false
		}
	}

	s.models = append(s.models, Model{actor, pos, tex, mesh})

	return true
}

// Removes model from renderer.
func (s *ModelRenderer) Remove(actor *Actor) bool {
	return s.RemoveById(actor.GetId())
}

// Removes model from renderer by ID.
func (s *ModelRenderer) RemoveById(id ActorId) bool {
	for i, model := range s.models {
		if model.Actor.GetId() == id {
			s.models = append(s.models[:i], s.models[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all sprites.
func (s *ModelRenderer) RemoveAll() {
	s.models = make([]Model, 0)
}

// Returns number of sprites.
func (s *ModelRenderer) Len() int {
	return len(s.models)
}

func (s *ModelRenderer) GetName() string {
	return model_renderer_name
}

// Render models.
func (s *ModelRenderer) Update(delta float64) {
	s.Shader.Bind()
	s.Shader.SendUniform1i(Default_shader_3D_tex, 0)

	if s.ortho {
		s.Shader.SendMat4(Default_shader_3D_pv, *MultMat4(s.Camera.CalcOrtho3D(), s.CalcModel()))
	} else {
		pv := s.Camera.CalcProjection()
		pv.Mult(s.Camera.CalcView())
		s.Shader.SendMat4(Default_shader_3D_pv, *pv)
	}

	var tid uint32

	for i := range s.models {
		if !s.models[i].Visible {
			continue
		}

		s.Shader.SendMat4(Default_shader_3D_model, *s.models[i].CalcModel())
		s.models[i].Vao.Bind()

		// prevent texture switching when not neccessary
		if tid != s.models[i].Tex.GetId() {
			tid = s.models[i].Tex.GetId()
			s.models[i].Tex.Bind()
		}

		gl.DrawElements(gl.TRIANGLES, s.models[i].Index.Size(), gl.UNSIGNED_INT, nil)
	}
}
