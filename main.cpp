#include "lib/raylib/raylib.h"

int main(void) {
  const int WIDTH = 1200;
  const int HEIGHT = 700;
  InitWindow(WIDTH, HEIGHT, "nature of code");

  while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
    BeginDrawing();
    ClearBackground(RAYWHITE);

    // Actual Code

    EndDrawing();
  }

  CloseWindow();

  return 0;
}
