#include "../lib/raylib/raylib.h"
#include "src.h"
#include <cmath>
#include <cstring>
#include <string>

// Debug physical body
void debug_body(Body &body) {
  // Velocity
  std::string vel_text_str = "vel_x: " + std::to_string(body.vel.x) +
                             " | vel_y: " + std::to_string(body.vel.y);
  char vel_text[vel_text_str.length() + 1];

  strcpy(vel_text, vel_text_str.c_str());
  DrawText(vel_text, 10, 40, 20, RAYWHITE);
  // Angle
  std::string angle_text_str = "angle: " + std::to_string(body.angle);
  char angle_text[angle_text_str.length() + 1];
  strcpy(angle_text, angle_text_str.c_str());
  DrawText(angle_text, 10, 80, 20, RAYWHITE);
}

// Draw rotational line
void DrawRotationalLine(Body &body) {
  Vector points[2];
  points[0] = Vector(body.location.x + body.mass * std::cos(body.angle),
                     body.location.y - body.mass * std::sin(body.angle));
  points[1] = Vector(body.location.x - body.mass * std::cos(body.angle),
                     body.location.y + body.mass * std::sin(body.angle));
  DrawLine(points[1].x, points[1].y, points[0].x, points[0].y, RAYWHITE);
}
