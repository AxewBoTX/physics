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

// Drawing rotational line
void DrawRotationalLine(Body &body) {
  Vector points[2];
  points[0] = Vector(body.location.x + body.mass * std::cos(body.angle),
                     body.location.y - body.mass * std::sin(body.angle));
  points[1] = Vector(body.location.x - body.mass * std::cos(body.angle),
                     body.location.y + body.mass * std::sin(body.angle));
  DrawLine(points[1].x, points[1].y, points[0].x, points[0].y, RAYWHITE);
}

// Draw Rotational Rectangle
void DrawRotationalRectangle(Body &body) {
  Vector points[4];
  // Top Left
  points[0] =
      Vector(body.location.x - (body.size.x / 2) * std::cos(body.angle) -
                 (body.size.y / 2) * std::sin(body.angle),
             body.location.y - (body.size.x / 2) * std::sin(body.angle) +
                 (body.size.y / 2) * std::cos(body.angle));
  // Top Right
  points[1] =
      Vector(body.location.x + (body.size.x / 2) * std::cos(body.angle) -
                 (body.size.y / 2) * std::sin(body.angle),
             body.location.y + (body.size.x / 2) * std::sin(body.angle) +
                 (body.size.y / 2) * std::cos(body.angle));
  // Bottom Right
  points[2] =
      Vector(body.location.x + (body.size.x / 2) * std::cos(body.angle) +
                 (body.size.y / 2) * std::sin(body.angle),
             body.location.y + (body.size.x / 2) * std::sin(body.angle) -
                 (body.size.y / 2) * std::cos(body.angle));
  // Bottom Left
  points[3] =
      Vector(body.location.x - (body.size.x / 2) * std::cos(body.angle) +
                 (body.size.y / 2) * std::sin(body.angle),
             body.location.y - (body.size.x / 2) * std::sin(body.angle) -
                 (body.size.y / 2) * std::cos(body.angle));

  // Top
  DrawLine(points[0].x, points[0].y, points[1].x, points[1].y, RAYWHITE);
  // Right
  DrawLine(points[1].x, points[1].y, points[2].x, points[2].y, RED);
  // Bottom
  DrawLine(points[2].x, points[2].y, points[3].x, points[3].y, RAYWHITE);
  // Left
  DrawLine(points[3].x, points[3].y, points[0].x, points[0].y, RAYWHITE);
}
