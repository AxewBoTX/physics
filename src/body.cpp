#include "../lib/raylib/raylib.h"
#include "src.h"
#include <cmath>

// Body Constructor
Body::Body(double _x, double _y, double _mass, double _angle) {
  this->mass = _mass;
  this->location.x = _x;
  this->location.y = _y;
  this->angle = _angle;
  this->size.x = this->mass * 2;
  this->size.y = this->mass;
}

// Updating body state
void Body::update() {
  this->handle_force_stacking();
  this->handle_input();
  // this->angle = std::atan2(this->vel.y, this->vel.x);
}

// Halding multiple forces on body
void Body::handle_force_stacking() {
  this->vel.limit(10);
  this->location.add(this->vel);
  this->vel.add(this->acc);
  this->acc.multi(0);
}

// Apply force to the body
void Body::apply_force(Vector &force) {
  Vector force_acc = Vector::Vector_div(force, this->mass);
  this->acc.add(force_acc);
}

// Show velocity directions
void Body::vel_direction() {
  Vector vel_vec = Vector::Vector_multi(this->vel, 8);
  Vector vel_dir = Vector::AddVectors(this->location, vel_vec);
  DrawLine(this->location.x, this->location.y, vel_dir.x, vel_dir.y, RED);
}

// Handle keyboard input
void Body::handle_input() {
  if (IsKeyDown(KEY_H)) {
    this->angle += (M_PI / 180) * 2;
  } else if (IsKeyDown(KEY_L)) {
    this->angle -= (M_PI / 180) * 2;
  }

  if (IsKeyDown(KEY_SPACE)) {
    Vector force =
        Vector(10 * std::cos(this->angle), 10 * std::sin(this->angle));
    this->apply_force(force);
  }

  if (IsKeyDown(KEY_D)) {
    Vector force = Vector(10, 0);
    this->apply_force(force);
  } else if (IsKeyDown(KEY_A)) {
    Vector force = Vector(-10, 0);
    this->apply_force(force);
  }
  if (IsKeyDown(KEY_S)) {
    Vector force = Vector(0, 10);
    this->apply_force(force);
  } else if (IsKeyDown(KEY_W)) {
    Vector force = Vector(0, -10);
    this->apply_force(force);
  }
}
