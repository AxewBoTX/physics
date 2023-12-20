package src

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Attractor struct {
	Mass     float64
	Location Vector
	G        float64
}

func NewAttractor(x, y, mass, G float64) *Attractor {
	return &Attractor{Mass: mass, Location: *NewVector(x, y), G: G}
}

func (this *Attractor) Draw() {
	raylib.DrawCircle(
		int32(this.Location.X),
		int32(this.Location.Y),
		float32(this.Mass),
		raylib.Red,
	)
	this.Keys()
}

func (this *Attractor) Attract(body *Body) {
	dir := SubVectors(this.Location, body.Location)
	distSQ := math.Max(math.Min(dir.Mag()*dir.Mag(), 25), 2500)
	strength := this.Mass * body.Mass * this.G / distSQ
	dir.Multi(strength)
	body.ApplyForce(*dir)
}

func (this *Attractor) Keys() {
	if raylib.IsKeyDown(raylib.KeyD) {
		this.Location.X += 5
	} else if raylib.IsKeyDown(raylib.KeyA) {
		this.Location.X -= 5
	}
	if raylib.IsKeyDown(raylib.KeyS) {
		this.Location.Y += 5
	} else if raylib.IsKeyDown(raylib.KeyW) {
		this.Location.Y -= 5
	}
}
