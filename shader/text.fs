#version 130
precision highp float;

uniform sampler2D tex;
uniform vec4 color;

in vec2 tc;

out vec4 c;

void main(){
	c = texture(tex, tc)*color;
}
