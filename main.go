package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"

	"nature/src"
)

func main() {
	const WIDTH = 1200
	const HEIGHT = 700
	raylib.InitWindow(WIDTH, HEIGHT, "nature of code")
	defer raylib.CloseWindow()

	const FPS = 60
	raylib.SetTargetFPS(FPS)

	world := src.NewWorld()
	world.AddBody(src.NewBody(WIDTH/2, HEIGHT/2, 15))

	for !raylib.WindowShouldClose() && !raylib.IsKeyPressed(raylib.KeyQ) {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)
		raylib.DrawFPS(10, 10)

		for i := range world.Bodies {
			world.Bodies[i].Draw()
			src.DebugBody(&world.Bodies[i])
			world.Bodies[i].Update()
			world.Bodies[i].VelDirection()
			world.GenerateFirction()
		}

		raylib.EndDrawing()
	}
}
