#include "lib/raylib/raylib.h"
#include "src/src.h"

int main(void) {
  // Basic raylib setup
  const double WIDTH = 1200;
  const double HEIGHT = 700;
  const int FPS = 60;
  InitWindow(WIDTH, HEIGHT, "nature of code");
  SetTargetFPS(FPS);

  // Engine stuff
  World world;
  for (int i = 0; i < 1; i++) {
    Body body = Body(WIDTH / 2, HEIGHT / 2, 20, 0);
    world.add_body(body);
  }

  while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
    BeginDrawing();
    ClearBackground(BLACK);
    DrawFPS(10, 10);

    // Actual Code
    for (int i = 0; i < world.bodies.size(); i++) {
      DrawRotationalRectangle(world.bodies[i]);
      world.bodies[i].update();
      // world.bodies[i].vel_direction();
      debug_body(world.bodies[i]);
    }
    world.generate_friction();

    EndDrawing();
  }

  CloseWindow();
  return 0;
}
