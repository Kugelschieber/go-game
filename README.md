# go-game (package "goga")

Game engine written in Go using OpenGL and GLFW. Mostly for 2D rendering, but also capable of rendering 3D, providing everything to get you started.
**Under heavy development, do not use yet!**

## Install

go-game requires OpenGL and GLFW. The following three steps install everything you need:

```
go get github.com/go-gl/gl/v4.5-core/gl
go get github.com/go-gl/glfw/v3.1/glfw
go get github.com/DeKugelschieber/go-game
```

You also need a cgo compiler (typically gcc) and GL/GLFW development libraries and headers. You can find further instructions on the GitHub pages below (see dependencies).

## Usage

*WIP*

Examples can be found within the demo folder. For full reference visit: https://godoc.org/github.com/DeKugelschieber/go-game

## Dependencies

* https://github.com/go-gl/gl
    - 4.5-core
* https://github.com/go-gl/glfw
    - 3.1

To use an older GL version, you need to replace the GL imports in package goga. It should mostly be compatible down to 3.x.

## Contribute

To contribute, please create pull requests. The code must be formatted by gofmt and fit into the architecture.

## License

MIT
