package goga

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"log"
)

// Checks for GL errors and prints to log if one occured.
func CheckGLError() {
	error := gl.GetError()

	if error != 0 {
		log.Print(error)
	}
}
