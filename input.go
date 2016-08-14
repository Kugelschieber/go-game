package goga

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	keyboardListener []KeyboardListener
	mouseListener    []MouseListener
)

// Interface for keyboard input events.
// Implement and register to receive keyboard input.
type KeyboardListener interface {
	OnKeyEvent(glfw.Key, int, glfw.Action, glfw.ModifierKey)
}

// Interface for mouse input events.
// Implement and register to receive mouse input.
type MouseListener interface {
	OnMouseButton(glfw.MouseButton, glfw.Action, glfw.ModifierKey)
	OnMouseMove(float64, float64)
	OnMouseScroll(float64, float64)
}

func initInput(wnd *glfw.Window) {
	wnd.SetKeyCallback(keyboardCallback)
	wnd.SetMouseButtonCallback(mouseButtonCallback)
	wnd.SetCursorPosCallback(mouseMoveCallback)
	wnd.SetScrollCallback(mouseScrollCallback)

	keyboardListener = make([]KeyboardListener, 0)
	mouseListener = make([]MouseListener, 0)
}

func keyboardCallback(wnd *glfw.Window, key glfw.Key, code int, action glfw.Action, mod glfw.ModifierKey) {
	for _, listener := range keyboardListener {
		listener.OnKeyEvent(key, code, action, mod)
	}
}

func mouseButtonCallback(wnd *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	for _, listener := range mouseListener {
		listener.OnMouseButton(button, action, mod)
	}
}

func mouseMoveCallback(wnd *glfw.Window, x float64, y float64) {
	for _, listener := range mouseListener {
		listener.OnMouseMove(x, float64(viewportHeight)-y)
	}
}

func mouseScrollCallback(wnd *glfw.Window, x float64, y float64) {
	for _, listener := range mouseListener {
		listener.OnMouseScroll(x, y)
	}
}

// Adds a new keyboard listener.
func AddKeyboardListener(listener KeyboardListener) {
	keyboardListener = append(keyboardListener, listener)
}

// Removes given keyboard listener if found.
func RemoveKeyboardListener(listener KeyboardListener) {
	for i, l := range keyboardListener {
		if l == listener {
			keyboardListener = append(keyboardListener[:i], keyboardListener[i+1:]...)
			return
		}
	}
}

// Removes all registered keyboard listeners.
func RemoveAllKeyboardListener() {
	keyboardListener = make([]KeyboardListener, 0)
}

// Adds a new mouse listener.
func AddMouseListener(listener MouseListener) {
	mouseListener = append(mouseListener, listener)
}

// Removes given mouse listener if found.
func RemoveMouseListener(listener MouseListener) {
	for i, l := range mouseListener {
		if l == listener {
			mouseListener = append(mouseListener[:i], mouseListener[i+1:]...)
			return
		}
	}
}

// Removes all registered mouse listeners.
func RemoveAllMouseListener() {
	mouseListener = make([]MouseListener, 0)
}
