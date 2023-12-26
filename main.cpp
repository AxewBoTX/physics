#include "lib/raylib/raylib.h"

int main(void) {
  const int WIDTH = 1200;
  const int HEIGHT = 700;
  InitWindow(WIDTH, HEIGHT, "nature of code");

  const int FPS = 60;
  SetTargetFPS(FPS);

  while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
    BeginDrawing();
    ClearBackground(BLACK);
    DrawFPS(10, 10);

    // Actual Code

    EndDrawing();
  }

  CloseWindow();

  return 0;
}
