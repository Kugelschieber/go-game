package goga

// A scene used to switch between game states.
// The Cleanup() method is called when a scene is removed
// or the program is stopped. It can be used to cleanup open resources
// (like GL objects).
type Scene interface {
	Pause()
	Resume()
	Cleanup()
	Resize(int, int)
	GetName() string
}

var (
	scenes      []Scene
	activeScene Scene
)

// Adds a scene to game.
// Returns false if the scene exists already.
func AddScene(scene Scene) bool {
	for _, s := range scenes {
		if s == scene {
			return false
		}
	}

	scenes = append(scenes, scene)

	return true
}

// Removes a given scene.
// Returns false if it could not be found.
func RemoveScene(scene Scene) bool {
	for i, s := range scenes {
		if s == scene {
			s.Cleanup()
			scenes = append(scenes[:i], scenes[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all scenes.
func RemoveAllScenes() {
	for _, s := range scenes {
		s.Cleanup()
	}

	scenes = make([]Scene, 0)
}

// Finds and returns a scene by name, or nil if not found.
func GetSceneByName(name string) Scene {
	for _, s := range scenes {
		if s.GetName() == name {
			return s
		}
	}

	return nil
}

// Switches to given scene.
// This will pause the currently active scene.
func SwitchScene(scene Scene) {
	if activeScene != nil {
		activeScene.Pause()
	}

	activeScene = scene

	for _, s := range scenes {
		if s == activeScene {
			activeScene.Resume()
			break
		}
	}
}

// Switches to given existing scene by name.
// Returns false if the scene does not exist.
func SwitchSceneByName(name string) bool {
	scene := GetSceneByName(name)

	if scene == nil {
		return false
	}

	SwitchScene(scene)

	return true
}

// Returns the currently active scene.
// Can be nil if no scene was set.
func GetActiveScene() Scene {
	return activeScene
}
