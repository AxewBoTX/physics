// #include "lib/raylib/raylib.h"
#include "src/src.h"
#include <iostream>

int main(void) {
  // const int WIDTH = 1200;
  // const int HEIGHT = 700;
  // InitWindow(WIDTH, HEIGHT, "nature of code");
  //
  // const int FPS = 60;
  // SetTargetFPS(FPS);
  //
  // while (!WindowShouldClose() && !IsKeyPressed(KEY_Q)) {
  //   BeginDrawing();
  //   ClearBackground(BLACK);
  //   DrawFPS(10, 10);
  //
  //   // Actual Code
  //
  //   EndDrawing();
  // }
  //
  // CloseWindow();
  Vector location(10, 10);
  Vector new_loc = location.get();
  new_loc.add(location).multi(10);
  std::cout << "location_x:" << location.x << " location_y:" << location.y
            << std::endl;
  std::cout << "new_x:" << new_loc.x << " new_y:" << new_loc.y << std::endl;
  return 0;
}
