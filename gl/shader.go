package gl

import (
	"errors"
	"geo"
	"github.com/go-gl/gl/v4.5-core/gl"
	"strings"
)

// Combination of shaders and shader program.
type Shader struct {
	program, vertex, fragment uint32
	attrIndex                 uint32
	attributes                []uint32
	uniformMap                map[string]int32
	attrMap                   map[string]int32
}

// Creates a new shader program by given vertex and fragment shader source code.
// The shaders itself will be deleted when compiled, only the program ID will be kept.
func NewShader(vertexShader, fragmentShader string) (*Shader, error) {
	shader := &Shader{}
	shader.attributes = make([]uint32, 0)
	shader.uniformMap = make(map[string]int32)
	shader.attrMap = make(map[string]int32)

	shader.program = gl.CreateProgram()
	shader.vertex = gl.CreateShader(gl.VERTEX_SHADER)
	shader.fragment = gl.CreateShader(gl.FRAGMENT_SHADER)

	if err := compileShader(&shader.vertex, vertexShader+NullTerminator); err != nil {
		gl.DeleteShader(shader.vertex)
		gl.DeleteShader(shader.fragment)
		shader.Drop()

		return nil, err
	}

	if err := compileShader(&shader.fragment, fragmentShader+NullTerminator); err != nil {
		gl.DeleteShader(shader.vertex)
		gl.DeleteShader(shader.fragment)
		shader.Drop()

		return nil, err
	}

	gl.AttachShader(shader.program, shader.vertex)
	gl.AttachShader(shader.program, shader.fragment)

	if err := linkProgram(shader.program); err != nil {
		gl.DeleteShader(shader.vertex)
		gl.DeleteShader(shader.fragment)
		shader.Drop()

		return nil, err
	}

	// we don't need to keep them in memory
	gl.DetachShader(shader.program, shader.vertex)
	gl.DetachShader(shader.program, shader.fragment)
	gl.DeleteShader(shader.vertex)
	gl.DeleteShader(shader.fragment)

	return shader, nil
}

func compileShader(shader *uint32, source string) error {
	csrc := gl.Str(source)
	gl.ShaderSource(*shader, 1, &csrc, nil)
	gl.CompileShader(*shader)

	if err := shaderCheckError(*shader); err != nil {
		return err
	}

	return nil
}

func shaderCheckError(shader uint32) error {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return errors.New("compiler error:\r\n" + log)
	}

	return nil
}

func linkProgram(program uint32) error {
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return errors.New("linker error:\r\n" + log)
	}

	return nil
}

// Deletes the GL program object.
func (s *Shader) Drop() {
	gl.DeleteProgram(s.program)
}

// Binds the shader for usage.
func (s *Shader) Bind() {
	gl.UseProgram(s.program)
}

// Unbinds shader.
func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

// Binds an attribute to this shader, name must be present in shader source.
func (s *Shader) BindAttribIndex(name string, index uint32) {
	gl.BindAttribLocation(s.program, index, gl.Str(name+NullTerminator))
	s.attributes = append(s.attributes, index)
}

// Binds an attribute to this shader (with default index), name must be present in shader source.
func (s *Shader) BindAttrib(name string) {
	s.BindAttribIndex(name, s.attrIndex)
	s.attrIndex++
}

// Enables all bound attributes.
func (s *Shader) EnableVertexAttribArrays() {
	for i := range s.attributes {
		gl.EnableVertexAttribArray(s.attributes[i])
	}
}

// Disables all bound attributes.
func (s *Shader) DisableVertexAttribArrays() {
	for i := range s.attributes {
		gl.DisableVertexAttribArray(s.attributes[i])
	}
}

// Returns the uniform location of a variable.
// The name must be present in the shader source.
func (s *Shader) GetUniformLocation(name string) int32 {
	_, exists := s.uniformMap[name]

	if !exists {
		s.uniformMap[name] = gl.GetUniformLocation(s.program, gl.Str(name+NullTerminator))
	}

	return s.uniformMap[name]
}

// Returns the location of an attribute.
// The name must be present in the shader source.
func (s *Shader) GetAttribLocation(name string) int32 {
	_, exists := s.attrMap[name]

	if !exists {
		s.attrMap[name] = gl.GetAttribLocation(s.program, gl.Str(name+NullTerminator))
	}

	return s.attrMap[name]
}

func (s *Shader) SendUniform1i(name string, v int32) {
	gl.Uniform1i(s.GetUniformLocation(name), v)
}

func (s *Shader) SendUniform2i(name string, v0, v1 int32) {
	gl.Uniform2i(s.GetUniformLocation(name), v0, v1)
}

func (s *Shader) SendUniform3i(name string, v0, v1, v2 int32) {
	gl.Uniform3i(s.GetUniformLocation(name), v0, v1, v2)
}

func (s *Shader) SendUniform4i(name string, v0, v1, v2, v3 int32) {
	gl.Uniform4i(s.GetUniformLocation(name), v0, v1, v2, v3)
}

func (s *Shader) SendUniform1f(name string, v float32) {
	gl.Uniform1f(s.GetUniformLocation(name), v)
}

func (s *Shader) SendUniform2f(name string, v0, v1 float32) {
	gl.Uniform2f(s.GetUniformLocation(name), v0, v1)
}

func (s *Shader) SendUniform3f(name string, v0, v1, v2 float32) {
	gl.Uniform3f(s.GetUniformLocation(name), v0, v1, v2)
}

func (s *Shader) SendUniform4f(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(s.GetUniformLocation(name), v0, v1, v2, v3)
}

func (s *Shader) SendUniform1iv(name string, count int32, data *int32) {
	gl.Uniform1iv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform2iv(name string, count int32, data *int32) {
	gl.Uniform2iv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform3iv(name string, count int32, data *int32) {
	gl.Uniform3iv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform4iv(name string, count int32, data *int32) {
	gl.Uniform4iv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform1fv(name string, count int32, data *float32) {
	gl.Uniform1fv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform2fv(name string, count int32, data *float32) {
	gl.Uniform2fv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform3fv(name string, count int32, data *float32) {
	gl.Uniform3fv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform4fv(name string, count int32, data *float32) {
	gl.Uniform4fv(s.GetUniformLocation(name), count, data)
}

func (s *Shader) SendUniform2x2(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix2fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform3x3(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix3fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform4x4(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix4fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform2x3(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix2x3fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform3x2(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix3x2fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform2x4(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix2x4fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform4x2(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix4x2fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform3x4(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix3x4fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendUniform4x3(name string, data *float32, count int32, transpose bool) {
	gl.UniformMatrix4x3fv(s.GetUniformLocation(name), count, transpose, data)
}

func (s *Shader) SendMat3(name string, m geo.Mat3) {
	var data [9]float32

	for i := 0; i < 9; i++ {
		data[i] = float32(m.Values[i])
	}

	gl.UniformMatrix3fv(s.GetUniformLocation(name), 1, false, &data[0])
}

func (s *Shader) SendMat4(name string, m geo.Mat4) {
	var data [16]float32

	for i := 0; i < 16; i++ {
		data[i] = float32(m.Values[i])
	}

	gl.UniformMatrix4fv(s.GetUniformLocation(name), 1, false, &data[0])
}

// Retuns the program GL ID.
func (s *Shader) GetProgramId() uint32 {
	return s.program
}

// Returns the vertex shader GL ID.
func (s *Shader) GetVertexId() uint32 {
	return s.vertex
}

// Returns the fragment shader GL ID.
func (s *Shader) GetFragmentId() uint32 {
	return s.fragment
}
