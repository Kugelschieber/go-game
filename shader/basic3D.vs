#version 130

uniform mat4 pv, m;

in vec3 vertex;
in vec2 texCoord;

out vec2 tc;

void main(){
	tc = texCoord;
	gl_Position = pv*m*vec4(vertex, 1.0);
}
