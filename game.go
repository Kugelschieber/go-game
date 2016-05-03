package goga

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"math"
	"time"
)

const (
	default_width         = uint32(600)
	default_height        = uint32(400)
	default_title         = "Game"
	default_exit_on_close = true
)

// Run options allow to set some parameters on startup.
type RunOptions struct {
	Title       string
	Width       uint32
	Height      uint32
	ExitOnClose bool
}

// Main game object.
type Game interface {
}

var (
	running = true
)

// Creates a new window with given options and starts the game.
// The game struct must implement the Game interface.
// If options is nil, the default options will be used.
// This function will panic on error.
func Run(game Game, options *RunOptions) {
	// init glfw
	if err := glfw.Init(); err != nil {
		panic("Error initializing GLFW: " + err.Error())
	}

	defer glfw.Terminate()

	// create window
	width := default_width
	height := default_height
	title := default_title
	exitOnClose := default_exit_on_close

	if options != nil {
		width = options.Width
		height = options.Height
		title = options.Title
		exitOnClose = options.ExitOnClose
	}

	wnd, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)

	if err != nil {
		panic("Error creating GLFW window: " + err.Error())
	}

	// start and loop
	wnd.MakeContextCurrent()
	delta := time.Duration(0)
	var deltaSec float64

	for running {
		if exitOnClose && wnd.ShouldClose() {
			return
		}

		start := time.Now()

		if !math.IsInf(deltaSec, 0) && !math.IsInf(deltaSec, -1) {
			updateSystems(deltaSec)
		}

		delta = time.Since(start)
		deltaSec = delta.Seconds()
		wnd.SwapBuffers()
		glfw.PollEvents()
	}
}

// Stops the game and closes the window.
func Stop() {
	running = false
}
