package goga

import ()

// An actor ID is a unique integer,
// which can be used to reference an actor.
type ActorId uint64

// A basic actor, having a unique ID.
// Use NewActor() to create new actors.
type Actor struct {
	id ActorId
}

var (
	actorIdGen = ActorId(0)
)

// Creates a new basic actor with unique ID.
func NewActor() *Actor {
	actorIdGen++
	return &Actor{actorIdGen}
}

// Returns the ID of actor.
func (a *Actor) GetId() ActorId {
	return a.id
}
