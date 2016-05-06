package goga

const (
	default_camera_pos_x = 5
	default_camera_pos_y = 5
	default_camera_pos_z = 5
	default_camera_up_x  = 0
	default_camera_up_y  = 0
	default_camera_up_z  = 1
	default_camera_fov   = 60
	default_camera_znear = 1
	default_camera_zfar  = 20
)

type Camera struct {
	Viewport                  Vec4
	Position, LookAt, Up      Vec3
	Fov, Ratio, Znear, Zfar   float64
	Projection, Ortho3D, View Mat4
	Ortho                     Mat3
}

// Creates a new 2D/3D camera.
// Pass the viewport as arguments.
func NewCamera(x, y, width, height int) *Camera {
	camera := &Camera{}
	camera.Position = Vec3{default_camera_pos_x, default_camera_pos_y, default_camera_pos_z}
	camera.Up = Vec3{default_camera_up_x, default_camera_up_y, default_camera_up_z}
	camera.Fov = default_camera_fov
	camera.Znear = default_camera_znear
	camera.Zfar = default_camera_zfar
	camera.SetViewport(x, y, width, height)
	camera.CalcRatio()

	return camera
}

// Updates viewport.
func (c *Camera) SetViewport(x, y, width, height int) {
	c.Viewport = Vec4{float64(x), float64(y), float64(width), float64(height)}
}

// Calculates viewport ratio (width/height).
func (c *Camera) CalcRatio() {
	c.Ratio = (c.Viewport.Z - c.Viewport.X) / (c.Viewport.W - c.Viewport.Y)
}

// Calculates projection matrix and returns it.
func (c *Camera) CalcProjection() *Mat4 {
	c.Projection.Identity()
	c.Projection.Perspective(c.Fov, c.Ratio, c.Znear, c.Zfar)

	return &c.Projection
}

// Calculates orthogonal projection matrix and returns it.
func (c *Camera) CalcOrtho() *Mat3 {
	c.Ortho.Identity()
	c.Ortho.Ortho(c.Viewport)

	return &c.Ortho
}

// Calculates 3D orthogonal projection matrix and returns it.
func (c *Camera) CalcOrtho3D() *Mat4 {
	c.Ortho3D.Identity()
	c.Ortho3D.Ortho(c.Viewport, c.Znear, c.Zfar)

	return &c.Ortho3D
}

// Calculates view matrix and returns it.
func (c *Camera) CalcView() *Mat4 {
	c.View.Identity()
	c.View.LookAt(c.Position, c.LookAt, c.Up)

	return &c.View
}
