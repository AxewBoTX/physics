package src

// Important functions and extensions of raylib-go

import (
	"fmt"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

// Debug Body
func DebugBody(body *Body) {
	raylib.DrawText(
		fmt.Sprintf(`Vel_X: %f | Vel_Y: %f`, body.Vel.X, body.Vel.Y), 10, 30, 25, raylib.DarkGreen,
	)
}

// Drawing Rotatable Line
func DrawRotationalLine(location *Vector, length, angle float64) {
	var Points [2]Vector
	Points[0] = *NewVector(location.X+length*math.Cos(angle), location.Y-length*math.Sin(angle))
	Points[1] = *NewVector(location.X-length*math.Cos(angle), location.Y+length*math.Sin(angle))
	raylib.DrawLine(
		int32(
			Points[0].X,
		),
		int32(Points[0].Y),
		int32(Points[1].X),
		int32(Points[1].Y),
		raylib.Black,
	)
}

// Drawing Rotatable Rectangle
func DrawRotationalRect(location *Vector, width, height, angle float64) {
	radius := math.Sqrt(width*width+height*height) / 2

	var Points [4]Vector
	Points[0] = *NewVector( // Top Left
		location.X-(radius*math.Sin(angle+math.Pi/4)),
		location.Y-(radius*math.Cos(angle+math.Pi/4)),
	)
	Points[1] = *NewVector( // Top Right
		location.X+(radius*math.Cos(angle+math.Pi/4)),
		location.Y-(radius*math.Sin(angle+math.Pi/4)),
	)
	Points[2] = *NewVector( // Bottom Right
		location.X+(radius*math.Sin(angle+math.Pi/4)),
		location.Y+(radius*math.Cos(angle+math.Pi/4)),
	)
	Points[3] = *NewVector( // Bottom Left
		location.X-(radius*math.Cos(angle+math.Pi/4)),
		location.Y+(radius*math.Sin(angle+math.Pi/4)),
	)

	raylib.DrawLine( //Top
		int32(Points[0].X),
		int32(Points[0].Y),
		int32(Points[1].X),
		int32(Points[1].Y),
		raylib.Black,
	)
	raylib.DrawLine( // Right
		int32(Points[1].X),
		int32(Points[1].Y),
		int32(Points[2].X),
		int32(Points[2].Y),
		raylib.Black,
	)
	raylib.DrawLine( //Bottom
		int32(Points[2].X),
		int32(Points[2].Y),
		int32(Points[3].X),
		int32(Points[3].Y),
		raylib.Black,
	)
	raylib.DrawLine( //Left
		int32(Points[3].X),
		int32(Points[3].Y),
		int32(Points[0].X),
		int32(Points[0].Y),
		raylib.Black,
	)
}
