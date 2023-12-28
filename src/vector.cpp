#include "src.h"
#include <cmath>

Vector::Vector(double _x, double _y) {
  this->x = _x;
  this->y = _y;
}

Vector Vector::get() { return Vector(this->x, this->y); }

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

Vector &Vector::div(double n) {
  this->x /= n;
  this->y /= n;
  return *this;
}

double Vector::mag() { return sqrt(pow(this->x, 2) + pow(this->y, 2)); }

Vector &Vector::normalize() {
  double m = this->mag();
  if (m != 0) {
    this->div(m);
  }
  return *this;
}

Vector &Vector::limit(double max) {
  if (this->mag() > max) {
    this->normalize();
    this->multi(max);
  }
  return *this;
}

// Static Methods
Vector Vector::AddVectors(Vector &v1, Vector &v2) {
  return Vector(v1.x + v2.x, v1.y + v2.y);
}
Vector Vector::SubVectors(Vector &v1, Vector &v2) {
  return Vector(v1.x - v2.x, v1.y - v2.y);
}
Vector Vector::Vector_multi(Vector &v1, double n) {
  return Vector(v1.x * n, v1.y * n);
}
Vector Vector::Vector_div(Vector &v1, double n) {
  return Vector(v1.x / n, v1.y / n);
}
