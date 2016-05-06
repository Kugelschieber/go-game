package goga

import ()

// A system provides logic for actors satisfying required components.
// They are automatically updated on each frame.
// When a system is removed from systems, the Cleanup() method will be called.
// This will also happen on program stop. It can be used to cleanup open resources
// (like GL objects).
type System interface {
	Update(float64)
	Cleanup()
	Add(interface{}) bool
	Remove(interface{}) bool
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

// Takes an actor and tries to add it to all systems that accept it.
// This maybe not as performant as directly adding it to the right system.
// Returns true if it could be added to at least one system, else false.
func AddActor(actor interface{}) bool {
	accepted := false

	for _, system := range systems {
		if system.Add(actor) {
			accepted = true
		}
	}

	return accepted
}

// Removes an actor from all systems.
// This maybe not as performant as directly removing it from the right system.
// Returns true if it could be removed from at least one system, else false.
func RemoveActor(actor interface{}) bool {
	removed := false

	for _, system := range systems {
		if system.Remove(actor) {
			removed = true
		}
	}

	return removed
}

// Removes an actor from all systems by ID.
// This maybe not as performant as directly removing it from the right system.
// Returns true if it could be removed from at least one system, else false.
func RemoveActorById(id ActorId) bool {
	removed := false

	for _, system := range systems {
		if system.RemoveById(id) {
			removed = true
		}
	}

	return removed
}
