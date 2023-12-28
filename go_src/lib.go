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
		fmt.Sprintf(`Vel_X: %f | Vel_Y: %f`, body.Vel.X, body.Vel.Y), 10, 40, 25, raylib.DarkGray,
	)
	raylib.DrawText(
		fmt.Sprintf(`body.Angle: %f`, body.Angle), 10, 70, 25, raylib.DarkGray,
	)
}

// Drawing Rotatable Line
func DrawRotationalLine(body *Body) {
	var Points [2]Vector
	length := body.Mass * 4
	Points[0] = *NewVector(body.Location.X+length*math.Cos(body.Angle), body.Location.Y-length*math.Sin(body.Angle))
	Points[1] = *NewVector(body.Location.X-length*math.Cos(body.Angle), body.Location.Y+length*math.Sin(body.Angle))
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

// Drawing Rotatable Square
func DrawRotationalSquare(body *Body) {
	radius := math.Sqrt(32*body.Mass*body.Mass) / 2

	var Points [4]Vector
	Points[0] = *NewVector( // Top Left
		body.Location.X-(radius*math.Sin(body.Angle+math.Pi/4)),
		body.Location.Y-(radius*math.Cos(body.Angle+math.Pi/4)),
	)
	Points[1] = *NewVector( // Top Right
		body.Location.X+(radius*math.Cos(body.Angle+math.Pi/4)),
		body.Location.Y-(radius*math.Sin(body.Angle+math.Pi/4)),
	)
	Points[2] = *NewVector( // Bottom Right
		body.Location.X+(radius*math.Sin(body.Angle+math.Pi/4)),
		body.Location.Y+(radius*math.Cos(body.Angle+math.Pi/4)),
	)
	Points[3] = *NewVector( // Bottom Left
		body.Location.X-(radius*math.Cos(body.Angle+math.Pi/4)),
		body.Location.Y+(radius*math.Sin(body.Angle+math.Pi/4)),
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
