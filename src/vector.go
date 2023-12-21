package src

// Vector Class and functions

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) *Vector {
	return &Vector{X: x, Y: y}
}

func (v *Vector) Get() *Vector {
	vec := Vector{X: v.X, Y: v.Y}
	return &vec
}

func (v *Vector) Add(v2 Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector) Sub(v2 Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) Multi(n float64) *Vector {
	v.X *= n
	v.Y *= n
	return v
}

func (v *Vector) Div(n float64) *Vector {
	v.X /= n
	v.Y /= n
	return v
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vector) Normalize() *Vector {
	m := v.Mag()
	if m != 0 {
		v.Div(m)
	}
	return v
}

func (v *Vector) Limit(max float64) *Vector {
	if v.Mag() > max {
		v.Normalize()
		v.Multi(max)
	}
	return v
}

func AddVectors(v1, v2 Vector) *Vector {
	return &Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func SubVectors(v1, v2 Vector) *Vector {
	return &Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func Vector_Multi(v Vector, n float64) *Vector {
	return &Vector{X: v.X * n, Y: v.Y * n}
}

func Vector_Div(v Vector, n float64) *Vector {
	return &Vector{X: v.X / n, Y: v.Y / n}
}
