package goga

const (
	culling_2d_name = "culling2d"
)

type Cullable struct {
	*Actor
	*Pos2D
}

type Culling2D struct {
	cullables []Cullable
	viewport  Vec4
}

// Creates a new sprite culling system.
// To update the viewport, call SetViewport().
func NewCulling2D(x, y, width, height int) *Culling2D {
	culling := &Culling2D{}
	culling.cullables = make([]Cullable, 0)
	culling.viewport = Vec4{float64(x), float64(y), float64(width), float64(height)}

	return culling
}

// Sets the culling outer bounds.
// Actors outside of this box won't be rendered.
func (c *Culling2D) SetViewport(x, y, width, height int) {
	c.viewport = Vec4{float64(x), float64(y), float64(width), float64(height)}
}

func (c *Culling2D) Cleanup() {}

// Adds actor with Pos2D to the system.
func (c *Culling2D) Add(actor *Actor, pos *Pos2D) bool {
	id := actor.GetId()

	for _, cull := range c.cullables {
		if id == cull.Actor.GetId() {
			return false
		}
	}

	c.cullables = append(c.cullables, Cullable{actor, pos})

	return true
}

// Removes actor with Pos2D from system.
func (c *Culling2D) Remove(actor *Actor) bool {
	return c.RemoveById(actor.GetId())
}

// Removes actor with Pos2D from system by ID.
func (c *Culling2D) RemoveById(id ActorId) bool {
	for i, cull := range c.cullables {
		if cull.GetId() == id {
			c.cullables = append(c.cullables[:i], c.cullables[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all cullable objects.
func (c *Culling2D) RemoveAll() {
	c.cullables = make([]Cullable, 0)
}

// Returns number of cullable objects.
func (c *Culling2D) Len() int {
	return len(c.cullables)
}

func (c *Culling2D) GetName() string {
	return culling_2d_name
}

// Updates visibility of all contained sprites.
func (c *Culling2D) Update(delta float64) {
	for i := range c.cullables {
		if c.cullables[i].Pos.X > c.viewport.Z ||
			c.cullables[i].Pos.X+c.cullables[i].Size.X < c.viewport.X ||
			c.cullables[i].Pos.Y > c.viewport.W ||
			c.cullables[i].Pos.Y+c.cullables[i].Size.Y < c.viewport.Y {
			c.cullables[i].Visible = false
		} else {
			c.cullables[i].Visible = true
		}
	}
}
