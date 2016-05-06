#version 130
precision highp float;

uniform sampler2D tex;

in vec2 tc;

out vec4 color;

void main(){
	color = texture(tex, tc);
}
