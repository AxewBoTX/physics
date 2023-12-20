package src

import (
	"fmt"
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func DrawRotationalRect(location *Vector, width, height, angle float64) {
	raylib.DrawText(fmt.Sprint(location), 10.0, 30.0, 20.0, raylib.Black)
	raylib.DrawText(fmt.Sprint(width), 10.0, 50.0, 20.0, raylib.Black)
	raylib.DrawText(fmt.Sprint(height), 40.0, 50.0, 20.0, raylib.Black)
	raylib.DrawText(fmt.Sprint(angle), 10.0, 70.0, 20.0, raylib.Black)
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
	raylib.DrawCircleLines(
		int32(location.X),
		int32(location.Y),
		float32(radius),
		raylib.Black,
	)
}
