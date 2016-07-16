package goga

// Position component for 2D objects.
type Pos2D struct {
	Pos, Size, Scale, RotPoint Vec2
	Rot                        float64
	Visible                    bool
	M                          Mat3
}

// Position component for 3D objects
type Pos3D struct {
	Pos, Size, Scale, RotPoint, Rot Vec3
	Visible                         bool
	M                               Mat4
}

// Creates a default initialized Pos2D.
func NewPos2D() *Pos2D {
	m := Mat3{}
	m.Identity()
	return &Pos2D{Size: Vec2{1, 1}, Scale: Vec2{1, 1}, Visible: true, M: m}
}

// Calculates model matrix for 2D positioning.
func (p *Pos2D) CalcModel() *Mat3 {
	p.M.Identity()
	p.M.Translate(Vec2{p.Pos.X + p.RotPoint.X, p.Pos.Y + p.RotPoint.Y})
	p.M.Rotate(p.Rot)
	p.M.Translate(Vec2{-p.RotPoint.X, -p.RotPoint.Y})
	p.M.Scale(p.Size)
	p.M.Scale(p.Scale)

	return &p.M
}

// Returns the center of object.
// Assumes y = 0 is bottom left corner, if not you have to subtract height of object.
func (p *Pos2D) GetCenter() Vec2 {
	return Vec2{p.Pos.X + (p.Size.X*p.Scale.X)/2, p.Pos.Y + (p.Size.Y*p.Scale.Y)/2}
}

// Returns true when given point is within rectangle of this object.
func (p *Pos2D) PointInRect(point Vec2) bool {
	return point.X > p.Pos.X && point.X < p.Pos.X+p.Size.X*p.Scale.X && point.Y > p.Pos.Y && point.Y < p.Pos.Y+p.Size.Y*p.Scale.Y
}

// Creates a default initialized Pos3D.
func NewPos3D() *Pos3D {
	m := Mat4{}
	m.Identity()
	return &Pos3D{Size: Vec3{1, 1, 1}, Scale: Vec3{1, 1, 1}, Visible: true, M: m}
}

// Calculates model matrix for 3D positioning.
func (p *Pos3D) CalcModel() *Mat4 {
	p.M.Identity()
	p.M.Translate(Vec3{p.Pos.X + p.RotPoint.X, p.Pos.Y + p.RotPoint.Y, p.Pos.Z + p.RotPoint.Z})
	p.M.Rotate(p.Rot.X, Vec3{1, 0, 0})
	p.M.Rotate(p.Rot.Y, Vec3{0, 1, 0})
	p.M.Rotate(p.Rot.Z, Vec3{0, 0, 1})
	p.M.Translate(Vec3{-p.RotPoint.X, -p.RotPoint.Y, -p.RotPoint.Z})
	p.M.Scale(p.Size)
	p.M.Scale(p.Scale)

	return &p.M
}

// Returns the center of object.
// Assumes y = 0 is bottom left corner, if not you have to subtract height of object.
func (p *Pos3D) GetCenter() Vec3 {
	return Vec3{p.Pos.X + (p.Size.X*p.Scale.X)/2, p.Pos.Y + (p.Size.Y*p.Scale.Y)/2, p.Pos.Z + (p.Size.Z*p.Scale.Z)/2}
}

// Centers given sprite within rectangle.
// Does nothing if sprite is nil.
// TODO
/*func CenterSprite(sprite *Sprite, x, y, width, height int) {
	if sprite == nil {
		return
	}

	sprite.Pos.X = (float64(width-x) - sprite.Size.X) / 2
	sprite.Pos.Y = (float64(height-y) - sprite.Size.Y) / 2
}*/
