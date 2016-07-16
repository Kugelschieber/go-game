package goga

const (
	// constants for default 2D shader
	Default_shader_2D_vertex_attrib   = "vertex"
	Default_shader_2D_texcoord_attrib = "texCoord"
	Default_shader_2D_ortho           = "o"
	Default_shader_2D_model           = "m"
	Default_shader_2D_tex             = "tex"

	// source for 2D shader
	default_shader_2d_vertex_src = `#version 130
		uniform mat3 o, m;
		in vec2 vertex;
		in vec2 texCoord;
		out vec2 tc;
		void main(){
			tc = texCoord;
			gl_Position = vec4(o*m*vec3(vertex, 1.0), 1.0);
		}`
	default_shader_2d_fragment_src = `#version 130
		precision highp float;
		uniform sampler2D tex;
		in vec2 tc;
		out vec4 color;
		void main(){
			color = texture(tex, tc);
		}`

	// constants for default 3D shader
	Default_shader_3D_vertex_attrib   = "vertex"
	Default_shader_3D_texcoord_attrib = "texCoord"
	Default_shader_3D_pv              = "pv"
	Default_shader_3D_model           = "m"
	Default_shader_3D_tex             = "tex"

	// source for 3D shader
	default_shader_3d_vertex_src = `#version 130
		uniform mat4 pv, m;
		in vec3 vertex;
		in vec2 texCoord;
		out vec2 tc;
		void main(){
			tc = texCoord;
			gl_Position = pv*m*vec4(vertex, 1.0);
		}`
	default_shader_3d_fragment_src = `#version 130
		precision highp float;
		uniform sampler2D tex;
		in vec2 tc;
		out vec4 color;
		void main(){
			color = texture(tex, tc);
		}`

	// constants for default text shader
	Default_shader_text_vertex_attrib   = "vertex"
	Default_shader_text_texcoord_attrib = "texCoord"
	Default_shader_text_ortho           = "o"
	Default_shader_text_model           = "m"
	Default_shader_text_tex             = "tex"
	Default_shader_text_color           = "color"

	// source for text shader
	default_shader_text_vertex_src = `#version 130
		uniform mat3 o, m;
		in vec2 vertex;
		in vec2 texCoord;
		out vec2 tc;
		void main(){
			tc = texCoord;
			gl_Position = vec4(o*m*vec3(vertex, 1.0), 1.0);
		}`
	default_shader_text_fragment_src = `#version 130
		precision highp float;
		uniform sampler2D tex;
		uniform vec4 color;
		in vec2 tc;
		out vec4 c;
		void main(){
			c = texture(tex, tc)*color;
		}`
)
