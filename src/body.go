package src

// Physical Body

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Mass          float64
	Width         float64
	Height        float64
	Location      Vector
	Vel           Vector
	Acc           Vector
	RotationAngle float64
}

func NewBody(x, y, mass float64) *Body {
	return &Body{Mass: mass, Width: mass * 4, Height: mass * 4, Location: Vector{X: x, Y: y}}
}

// Drawing Function
func (body *Body) Draw() {
	DrawRotationalRect(&body.Location, body.Width, body.Height, body.RotationAngle)
}

// Updating Body State
func (body *Body) Update() {
	body.HandleForceStacking()
	body.Keys()
}

// Handle Multiple Forces on Body
func (body *Body) HandleForceStacking() {
	body.Vel.Limit(10)
	body.Location.Add(body.Vel)
	body.Vel.Add(body.Acc)
	body.Acc.Multi(0)
}

// Apply Force To The Body
func (body *Body) ApplyForce(force Vector) {
	body.Acc.Add(*Vector_Div(force, body.Mass))
}

// Attract other Body
func (body *Body) Attract(other *Body) {
	dir := SubVectors(body.Location, other.Location)
	distSQ := math.Max(math.Min(dir.Mag()*dir.Mag(), 25), 2500)
	strength := body.Mass * other.Mass * 0.3 / distSQ
	dir.Multi(strength)
	other.ApplyForce(*dir)
}

// Display Velocity Direction
func (body *Body) VelDirection() {
	vel_dir := AddVectors(body.Location, *Vector_Multi(body.Vel, 12))
	raylib.DrawLine(
		int32(
			body.Location.X,
		),
		int32(body.Location.Y),
		int32(vel_dir.X),
		int32(vel_dir.Y),
		raylib.Black,
	)
}

// Handling Keyboard Keyboard Input
func (body *Body) Keys() {
	if raylib.IsKeyDown(raylib.KeySpace) {
		body.RotationAngle += (math.Pi / 180) * 2
	}
	if raylib.IsKeyDown(raylib.KeyBackspace) {
		body.RotationAngle -= (math.Pi / 180) * 2
	}
	if raylib.IsKeyDown(raylib.KeyD) {
		body.ApplyForce(*NewVector(5, 0))
	} else if raylib.IsKeyDown(raylib.KeyA) {
		body.ApplyForce(*NewVector(-5, 0))
	}
	if raylib.IsKeyDown(raylib.KeyS) {
		body.ApplyForce(*NewVector(0, 5))
	} else if raylib.IsKeyDown(raylib.KeyW) {
		body.ApplyForce(*NewVector(0, -5))
	}
}
