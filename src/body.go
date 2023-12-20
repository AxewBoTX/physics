package src

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Mass     float64
	Width    float64
	Height   float64
	Location Vector
	Vel      Vector
	Acc      Vector
	Points   [4]Vector
	Rotation float64
}

func NewBody(x, y, mass float64) *Body {
	return &Body{Mass: mass, Width: mass * 4, Height: mass * 4, Location: Vector{X: x, Y: y}}
}

// Drawing Function
func (body *Body) Draw() {
	DrawRotationalRect(&body.Location, body.Width, body.Height, body.Rotation)
}

// Updating Body State
func (body *Body) Update() {
	body.Vel.Limit(15)
	body.Location.Add(body.Vel)
	body.Vel.Add(body.Acc)
	body.Acc.Multi(0)
	body.VelDirection()
	body.Keys()

	// //Friction
	// body.ApplyForce(
	// 	*body.Vel.Get().Multi(-1).Multi(0.2),
	// )
}

func (this *Body) Attract(other *Body) {
	dir := SubVectors(this.Location, other.Location)
	distSQ := math.Max(math.Min(dir.Mag()*dir.Mag(), 25), 2500)
	strength := this.Mass * other.Mass * 0.3 / distSQ
	dir.Multi(strength)
	other.ApplyForce(*dir)
}

// Apply Force To The Body
func (body *Body) ApplyForce(force Vector) {
	body.Acc.Add(*Vector_Div(force, body.Mass))
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

// Keyboard Input
func (body *Body) Keys() {
	if raylib.IsKeyDown(raylib.KeySpace) {
		body.Rotation += (math.Pi / 180) * 2
	}
	if raylib.IsKeyDown(raylib.KeyBackspace) {
		body.Rotation -= (math.Pi / 180) * 2
	}
	if raylib.IsKeyDown(raylib.KeyD) {
		body.Location.X += 5
	} else if raylib.IsKeyDown(raylib.KeyA) {
		body.Location.X -= 5
	}
	if raylib.IsKeyDown(raylib.KeyS) {
		body.Location.Y += 5
	} else if raylib.IsKeyDown(raylib.KeyW) {
		body.Location.Y -= 5
	}
}
