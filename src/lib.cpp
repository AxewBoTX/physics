#include "../lib/raylib/raylib.h"
#include "src.h"
#include <cstring>
#include <string>

// Debug physical body
void debug_body(Body &body) {
  std::string vel_text_str = "vel_x: " + std::to_string(body.vel.x) +
                             " | vel_y: " + std::to_string(body.vel.y);
  char vel_text[vel_text_str.length() + 1];

  strcpy(vel_text, vel_text_str.c_str());
  DrawText(vel_text, 10, 40, 25, RAYWHITE);
}
