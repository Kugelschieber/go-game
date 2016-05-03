package goga

import ()

// A system provides logic for actors satisfying required components.
// They are automatically updated on each frame.
// When a system is removed from systems, the Cleanup() method will be called.
// This will also happen on program stop. This can be used to cleanup open resources
// (like GL objects).
type System interface {
	Update(float64)
	Cleanup()
	Add(*Actor) bool
	Remove(*Actor) bool
	RemoveById(ActorId) bool
	RemoveAll()
	Len() int
	GetName() string
}

var (
	systems []System
)

// Adds a system to the game.
// Returns false if the system exists already.
func AddSystem(system System) bool {
	for _, sys := range systems {
		if sys == system {
			return false
		}
	}

	systems = append(systems, system)

	return true
}

// Removes the given system.
// Returns false if it could not be found.
func RemoveSystem(system System) bool {
	for i, sys := range systems {
		if sys == system {
			sys.Cleanup()
			systems = append(systems[:i], systems[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all systems.
func RemoveAllSystems() {
	for _, system := range systems {
		system.Cleanup()
	}

	systems = make([]System, 0)
}

// Finds and returns a system by name, or nil if not found.
func GetSystemByName(name string) System {
	for _, system := range systems {
		if system.GetName() == name {
			return system
		}
	}

	return nil
}

func updateSystems(delta float64) {
	for _, system := range systems {
		system.Update(delta)
	}
}
