#include "lib/raylib/raylib.h"

int main(void) {
  // Basic raylib setup
  const double WIDTH = 1200;
  const double HEIGHT = 700;
  const int FPS = 60;
  InitWindow(WIDTH, HEIGHT, "nature of code");
  SetTargetFPS(FPS);

  while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
    BeginDrawing();
    ClearBackground(BLACK);
    DrawFPS(10, 10);

    EndDrawing();
  }

  CloseWindow();
  return 0;
}
