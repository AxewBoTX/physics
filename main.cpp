#include "lib/raylib/raylib.h"
#include "src/src.h"

int main(void) {
  // Basic raylib setup
  const double WIDTH = 1200;
  const double HEIGHT = 700;
  const int FPS = 60;
  InitWindow(WIDTH, HEIGHT, "nature of code");
  SetTargetFPS(FPS);

  // Engine
  World world;
  for (int a = 0; a < 10; a++) {
    Body body(a * 30, a * 28, 1, 10);
    world.bodies.push_back(body);
  }

  while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
    BeginDrawing();
    ClearBackground(BLACK);
    DrawFPS(10, 10);

    for (int a = 0; a < world.bodies.size(); a++) {
      DrawCircle(world.bodies[a].location.x, world.bodies[a].location.y, 10,
                 RAYWHITE);
      world.bodies[a].update();
    }

    EndDrawing();
  }

  CloseWindow();
  return 0;
}
