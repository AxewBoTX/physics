#include "src.h"
#include <cmath>

Vector::Vector(float _x, float _y) {
  this->x = _x;
  this->y = _y;
}

Vector Vector::get() { return Vector(this->x, this->y); }

float Vector::mag() { return sqrt(pow(this->x, 2) + pow(this->y, 2)); }

Vector &Vector::add(Vector &v) {
  this->x += v.x;
  this->y += v.y;
  return *this;
}

Vector &Vector::sub(Vector &v) {
  this->x -= v.x;
  this->y -= v.y;
  return *this;
}

Vector &Vector::multi(double n) {
  this->x *= n;
  this->y *= n;
  return *this;
}
