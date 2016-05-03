package goga

import ()

// A system provides logic for actors satisfying required components.
// They are automatically updated on each frame.
type System interface {
	Update(float64)
	Add(actor *Actor) bool
	Remove(actor *Actor) bool
	RemoveById(id ActorId) bool
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
			systems = append(systems[:i], systems[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all systems.
func RemoveAllSystems() {
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
