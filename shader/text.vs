#version 130

uniform mat3 o, m;

in vec2 vertex;
in vec2 texCoord;

out vec2 tc;

void main(){
	tc = texCoord;
	gl_Position = vec4(o*m*vec3(vertex, 1.0), 1.0);
}
